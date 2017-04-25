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

	b, err := ioutil.ReadFile(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	var conf config
	if err = yaml.Unmarshal(b, &conf); err != nil {
		log.Fatal(err)
	}

	db, err := user.Open(conf.DatabaseURL, conf.Salt)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()

	reverse, err := proxy.New(conf.Backends...)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.DefaultServeMux
	mux.HandleFunc("/users", basicAuth(saveUser(db), conf.Auth.Username, conf.Auth.Password))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/metrics/find") {
			username, password, ok := r.BasicAuth()
			if !ok {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			u, err := db.Find(username, password)
			if err == user.ErrInvalidCredentials {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			q, ok := r.URL.Query()["query"]
			if !ok {

			}

			ok = false
			for _, rx := range u.Regexps {
				if rx.MatchString(q[0]) {
					ok = true
				}
			}

			if !ok {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}

		if err := reverse.Proxy(w, r); err != nil {
			httpError(w, err)
			return
		}
	})

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
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		h(w, r)
	}
}

// ANY /users
func saveUser(db *user.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// DELETE
		if r.Method == http.MethodDelete {
			username, ok := r.URL.Query()["username"]
			if ok && username[0] != "" {
				if err := db.Delete(username[0]); err != nil {
					httpError(w, err)
					return
				}
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		u := &user.User{}
		if err := json.NewDecoder(r.Body).Decode(u); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// validate fields
		if len(u.Username) < 4 || len(u.Password) < 4 || len(u.Regexps) == 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if err := db.Save(u); err != nil {
			httpError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func httpError(w http.ResponseWriter, err error) {
	log.Printf("error: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
