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

	mux := http.DefaultServeMux
	mux.HandleFunc("/users", usersHandler(db, conf.Auth.Username, conf.Auth.Password))
	mux.HandleFunc("/", carbonHandler(db, p))

	srv := http.Server{Addr: conf.Address, Handler: mux}
	if conf.TLS.CertFile != "" && conf.TLS.KeyFile != "" {
		log.Printf("starting https on %s", conf.Address)
		log.Fatal(srv.ListenAndServeTLS(conf.TLS.CertFile, conf.TLS.CertFile))
	}
	log.Printf("starting http on %s", conf.Address)
	log.Fatal(srv.ListenAndServe())
}

// basicAuth wraps h with http basic auth
func basicAuth(h http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || (username != u && password != p) {
			httpErrorCode(w, http.StatusUnauthorized)
			return
		}
		h(w, r)
	}
}

// DELETE POST /users
func usersHandler(db *user.DB, username, password string) http.HandlerFunc {
	return basicAuth(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			username, ok := r.URL.Query()["username"]
			if ok && username[0] != "" {
				if err := db.Delete(username[0]); err != nil {
					httpError(w, err)
					return
				}
			}
			w.WriteHeader(http.StatusOK)
			return
		case http.MethodPost:
			u := &user.User{}
			if err := json.NewDecoder(r.Body).Decode(u); err != nil {
				httpErrorCode(w, http.StatusBadRequest)
				return
			}

			// validate fields
			if len(u.Username) < 4 || len(u.Password) < 4 || len(u.Globs) == 0 {
				httpErrorCode(w, http.StatusBadRequest)
				return
			}

			if err := db.Save(u); err != nil {
				httpError(w, err)
				return
			}

			w.WriteHeader(http.StatusOK)
		default:
			httpErrorCode(w, http.StatusMethodNotAllowed)
		}
	}, username, password)
}

type filterFunc func(*user.User, *http.Response) ([]byte, error)

var endpoints = map[string]filterFunc{
	"/metrics/find": filterFind,
	//"/render":       "target",
}

// copy from carbonapi
type treejson struct {
	AllowChildren int            `json:"allowChildren"`
	Expandable    int            `json:"expandable"`
	Leaf          int            `json:"leaf"`
	ID            string         `json:"id"`
	Text          string         `json:"text"`
	Context       map[string]int `json:"context"` // unused
}

func filterFind(u *user.User, r *http.Response) ([]byte, error) {
	var tree []*treejson
	if err := json.NewDecoder(r.Body).Decode(&tree); err != nil {
		return nil, err
	}

	res := make([]*treejson, 0, len(tree))
	for _, leaf := range tree {
		if u.Can(leaf.ID) {
			res = append(res, leaf)
		}
	}
	return json.Marshal(res)
}

// matchRoute matches the named request path to one of
// supported routes and returns query parameter name or
// returns an empty string
func matchRoute(path string) filterFunc {
	for p, fn := range endpoints {
		if strings.HasPrefix(path, p) {
			return fn
		}
	}
	return nil
}

// ANY /
func carbonHandler(db *user.DB, p *proxy.Proxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if fn := matchRoute(r.URL.Path); fn != nil {
			username, password, _ := r.BasicAuth()
			u, err := db.FindByUsernameAndPassword(username, password)
			if err == user.ErrInvalidCredentials {
				httpErrorCode(w, http.StatusUnauthorized)
				return
			}

			res, err := p.Proxy(r)
			if err != nil {
				httpError(w, err)
				return
			}
			defer res.Body.Close()

			// filter response
			b, err := fn(u, res)
			if err != nil {
				httpError(w, err)
				return
			}

			for k, hh := range res.Header {
				for _, h := range hh {
					w.Header().Add(k, h)
				}
			}
			w.WriteHeader(res.StatusCode)
			w.Write(b)
		}
	}
}

func httpError(w http.ResponseWriter, err error) {
	log.Printf("error: %v", err)
	httpErrorCode(w, http.StatusInternalServerError)
}

func httpErrorCode(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
