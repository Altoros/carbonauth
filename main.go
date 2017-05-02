package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Altoros/carbonauth/proxy"
	"github.com/Altoros/carbonauth/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	glog "github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
)

type config struct {
	Address string `yaml:"address"`

	Backends []string `yaml:"backends"`

	Auth struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"auth"`

	Salt        string `yaml:"salt"`
	DatabaseURL string `yaml:"database_url"`

	TLS struct {
		CertFile string `yaml:"cert_file"`
		KeyFile  string `yaml:"key_file"`
	} `yaml:"tls"`
}

func main() {
	confPath := flag.String("c", "config.yml", "path to config file")
	flag.Parse()

	// read configuration from the provided yml file
	b, err := ioutil.ReadFile(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	var conf config
	if err = yaml.Unmarshal(b, &conf); err != nil {
		log.Fatalf("config parsing error: %v", err)
	}

	db, err := user.Open(conf.DatabaseURL, conf.Salt)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	p, err := proxy.New(conf.Backends...)
	if err != nil {
		log.Fatalf("proxy error: %v", err)
	}

	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(glog.DEBUG)
	e.Pre(middleware.Logger())
	e.Pre(middleware.Gzip())
	e.Pre(middleware.AddTrailingSlash())

	// users management
	g := e.Group("/users", authAdmin(conf.Auth.Username, conf.Auth.Password))
	g.POST("/", saveUser(db))
	g.GET("/:username/", showUser(db))
	g.DELETE("/:username/", deleteUser(db))

	// proxy section
	a := authUser(db)
	e.GET("/metrics/find*", carbonHandler(p, filterFind), a)
	e.POST("/render*", carbonHandler(p, filterRender), a)

	if conf.TLS.CertFile != "" && conf.TLS.KeyFile != "" {
		panic(e.StartTLS(conf.Address, conf.TLS.CertFile, conf.TLS.KeyFile))
	}
	panic(e.Start(conf.Address))
}

func authAdmin(username, password string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(u string, p string, _ echo.Context) (error, bool) {
		return nil, username == u && password == p
	})
}

func authUser(db *user.DB) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, c echo.Context) (error, bool) {
		u, err := db.FindByUsernameAndPassword(username, password)
		if err != nil {
			if err == user.ErrNotFound {
				err = nil
			}
			return err, false
		}

		// save user in the context
		c.Set(userKey, u)
		return nil, true
	})
}

// GET /users/:username
func showUser(db *user.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, err := db.FindByUsername(c.Param("username"))
		if err != nil {
			if err == user.ErrNotFound {
				err = echo.ErrNotFound
			}
			return err
		}
		return c.JSON(http.StatusOK, u)
	}
}

// POST /users
func saveUser(db *user.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u user.User
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			return err
		}

		// TODO: c.Validate
		if len(u.Username) < 4 || len(u.Password) < 4 || len(u.Globs) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		}

		if err := db.Save(&u); err != nil {
			return err
		}
		return c.NoContent(http.StatusOK)
	}
}

// DELETE /users/:username
func deleteUser(db *user.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return db.Delete(c.Param("username"))
	}
}

type filterFunc func(*user.User, *http.Response) ([]json.RawMessage, error)

// response json format: ["id": "...", ...}, ...]
func filterFind(u *user.User, r *http.Response) ([]json.RawMessage, error) {
	metrics := []json.RawMessage{}
	if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
		return nil, err
	}

	var m struct {
		ID string `json:"id"`
	}

	result := make([]json.RawMessage, 0, len(metrics))
	for _, metric := range metrics {
		if err := json.Unmarshal(metric, &m); err != nil {
			return nil, err
		}

		if !u.Can(m.ID) {
			continue
		}
		result = append(result, metric)
	}
	return result, nil
}

// response json format: [{"target": "...", ...}, ...]
func filterRender(u *user.User, r *http.Response) ([]json.RawMessage, error) {
	metrics := []json.RawMessage{}
	if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
		return nil, err
	}

	var m struct {
		Target string `json:"target"`
	}

	result := make([]json.RawMessage, 0, len(metrics))
	for _, metric := range metrics {
		if err := json.Unmarshal(metric, &m); err != nil {
			return nil, err
		}

		if !u.Can(m.Target) {
			continue
		}
		result = append(result, metric)
	}
	return result, nil
}

const userKey = "user"

// GET /metrics/find*
// POST /render*
func carbonHandler(p *proxy.Proxy, fn filterFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get(userKey).(*user.User)
		res, err := p.Proxy(c.Request())
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// copy response headers
		for k, hh := range res.Header {
			for _, h := range hh {
				c.Response().Header().Add(k, h)
			}
		}

		if res.StatusCode != http.StatusOK {
			return c.Stream(res.StatusCode, res.Header.Get("Content-Type"), res.Body)
		}

		if !strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
			return echo.ErrMethodNotAllowed
		}

		// filter response when needed
		v, err := fn(u, res)
		if err != nil {
			return err
		}
		return c.JSON(res.StatusCode, v)
	}
}
