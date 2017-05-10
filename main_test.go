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
	"encoding/json"
	"reflect"
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
			{"target": "foo.*", "datapoints": [[1, 0]]}
		]`))
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
		fn     func(p *proxy.Proxy) echo.HandlerFunc
		want   string
		method string
	}{
		"/metrics/find?query=foo.*": {
			fn:     findHandler,
			want:   `[{"id": "foo.bar.baz","isLeaf": 1}]`,
			method: http.MethodGet,
		},
		"/render?target=foo.*&format=json": {
			fn:     renderHandler,
			want:   `[{"target": "foo.*","datapoints": [[1, 0]]}]`,
			method: http.MethodPost,
		},
	} {
		t.Run(url, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(d.method, url, nil)
			r.SetBasicAuth("user", "secret")

			h := authUser(db)(d.fn(p))
			if err := h(e.NewContext(r, w)); err != nil {
				t.Errorf("GET %s err = %v", url, err)
			}

			ok, err := jsonEqual(w.Body.Bytes(), []byte(d.want))
			if err != nil {
				t.Fatal(err)
			}

			if !ok {
				t.Errorf("GET %s body = %q, want %q", url, w.Body.String(), d.want)
			}
		})
	}
}

func jsonEqual(j1, j2 []byte) (bool, error) {
	var v1 interface{}
	var v2 interface{}

	if err := json.Unmarshal(j1, &v1); err != nil {
		return false, err
	}
	if err := json.Unmarshal(j2, &v2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(v1, v2), nil
}
