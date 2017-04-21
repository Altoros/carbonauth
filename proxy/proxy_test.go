package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy_ServeHTTP(t *testing.T) {
	t.Parallel()

	b1 := newBackend("one")
	defer b1.Close()

	b2 := newBackend("two")
	defer b1.Close()

	p, err := New(b1.URL, b2.URL)
	if err != nil {
		t.Fatal(err)
	}

	res := map[string]int{}
	for i := 0; i < 10; i++ {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		p.ServeHTTP(w, r)
		res[w.Body.String()] = res[w.Body.String()] + 1
	}

	if res["one"] == 0 || res["two"] == 0 {
		t.Errorf("proxied to only one backend: one = %d, two = %d", res["one"], res["two"])
	}
}

func newBackend(msg string) *httptest.Server {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	})
	return httptest.NewServer(m)
}
