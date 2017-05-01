package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Altoros/carbonauth/proxy"
	"github.com/Altoros/carbonauth/user"
	"github.com/labstack/echo"
)

func TestFlow(t *testing.T) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "sqlite3://:memory:"
	}

	e := echo.New()

	db, err := user.Open(databaseURL, "seasalt")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Clean()

	// emulate carbonapi
	c := http.NewServeMux()
	c.HandleFunc("/metrics/find", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{"id": "foo.bar.baz", "isLeaf": 1},
			{"id": "bar.baz.bam", "isLeaf": 0}
		]`))
	})
	c.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{"target": "foo.*", "datapoints": [[1, 0]]},
			{"target": "bar.baz.bam", "datapoints": [[0, 1]]
		}]`))
	})

	b := httptest.NewServer(c)
	defer b.Close()

	p, err := proxy.New(b.URL)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users", strings.NewReader(`{
		"username": "user",
		"password": "secret",
		"globs":    ["foo.*", "bar.baz"]
	}`))
	r.SetBasicAuth("admin", "secret")

	h := authAdmin("admin", "secret")(saveUser(db))
	if err := h(e.NewContext(r, w)); err != nil {
		t.Errorf("POST /user err = %v", err)
	}

	for url, d := range map[string]struct {
		fn     filterFunc
		want   string
		method string
	}{
		"/metrics/find?query=foo.*": {
			fn:     filterFind,
			want:   `[{"id":"foo.bar.baz","isLeaf":1}]`,
			method: http.MethodGet,
		},
		"/render?target=foo.*&format=json": {
			fn:     filterRender,
			want:   `[{"target":"foo.*","datapoints":[[1,0]]}]`,
			method: http.MethodPost,
		},
	} {
		t.Run(url, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(d.method, url, nil)
			r.SetBasicAuth("user", "secret")

			h := authUser(db)(carbonHandler(p, d.fn))
			if err := h(e.NewContext(r, w)); err != nil {
				t.Errorf("GET %s err = %v", url, err)
			}

			if w.Body.String() != d.want {
				t.Errorf("GET %s body = %q, want %q", url, w.Body.String(), d.want)
			}
		})
	}
}
