package main

import (
	"fmt"
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
		t.Fatal("TEST_DATABASE_URL is not set")
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
		w.Write([]byte(`[{"id": "foo.bar.baz"}, {"id": "bar.baz"}]`))
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
		"globs":    ["foo.*"]
	}`))
	r.SetBasicAuth("admin", "secret")
	uh.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("create user code = %d, want %d", w.Code, http.StatusOK)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/metrics/find?query=foo.*", nil)
	r.SetBasicAuth("user", "secret")
	ch.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("carbonapi request code = %d, want %d", w.Code, http.StatusOK)
	}

	fmt.Println(w.Body.String())
}
