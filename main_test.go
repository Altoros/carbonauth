package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Altoros/carbonauth/proxy"
	"github.com/Altoros/carbonauth/user"
)

func TestFlow(t *testing.T) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "sqlite3://:memory:"
	}

	db, err := user.Open(databaseURL, "seasalt")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Clean()

	// emulate carbonapi
	h := http.NewServeMux()
	h.HandleFunc("/metrics/find", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{"id": "foo.bar.baz", "isLeaf": 1},
			{"id": "bar.baz.bam", "isLeaf": 0}
		]`))
	})
	h.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{"target": "foo.*", "datapoints": [[1, 0]]},
			{"target": "bar.baz.bam", "datapoints": [[0, 1]]
		}]`))
	})

	b := httptest.NewServer(h)
	defer b.Close()

	p, err := proxy.New(b.URL)
	if err != nil {
		t.Fatal(err)
	}

	uh := usersHandler(db, "admin", "secret")
	ch := carbonHandler(db, p)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users", strings.NewReader(`{
		"username": "user",
		"password": "secret",
		"globs":    ["foo.*", "bar.baz"]
	}`))
	r.SetBasicAuth("admin", "secret")
	uh.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("create user code = %d, want %d", w.Code, http.StatusOK)
	}

	for url, body := range map[string]string{
		"/metrics/find?query=foo.*":        `[{"id":"foo.bar.baz","isLeaf":1}]`,
		"/render?target=foo.*&format=json": `[{"target":"foo.*","datapoints":[[1,0]]}]`,
	} {
		t.Run(url, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", url, nil)
			r.SetBasicAuth("user", "secret")
			ch.ServeHTTP(w, r)

			if w.Code != http.StatusOK {
				t.Errorf("GET %s code = %d, want %d", url, w.Code, http.StatusOK)
			} else if w.Body.String() != body {
				t.Errorf("GET %s body = %q, want %q", url, w.Body.String(), body)
			}
		})
	}
}
