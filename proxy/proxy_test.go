package proxy

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestProxy_Proxy(t *testing.T) {
	t.Parallel()

	b1 := newBackend("one")
	defer b1.Close()

	b2 := newBackend("two")
	defer b1.Close()

	p, err := New(b1.URL, b2.URL)
	if err != nil {
		t.Fatal(err)
	}
	p.SuspendTime = time.Millisecond

	res := map[string]int{}
	for i := 0; i < 10; i++ {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/asd", nil)
		if err := p.Proxy(w, r); err != nil {
			t.Fatal(err)
		}
		res[w.Body.String()] = res[w.Body.String()] + 1
	}

	if res["one"] == 0 || res["two"] == 0 {
		t.Errorf("proxied to only one backend: one = %d, two = %d", res["one"], res["two"])
	}

	// stop first backend
	b1.Close()
	for i := 0; i < 10; i++ {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/asd", nil)
		if err := p.Proxy(w, r); err != nil {
			t.Fatal(err)
		}

		if w.Body.String() == "one" {
			t.Error("first backend expected to be down but it's not")
		}
	}

	// stop second backend
	b2.Close()
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/asd", nil)
	if err = p.Proxy(w, r); err != ErrNoBackends {
		t.Errorf("err = %v, want ErrNoBackends", err)
	}

	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("one"))
	})

	u, err := url.Parse(b1.URL)
	if err != nil {
		t.Fatal(err)
	}

	s := http.Server{Addr: ":" + u.Port(), Handler: m}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(10 * time.Millisecond)
	defer s.Close()

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/asd", nil)
	if err = p.Proxy(w, r); err != nil {
		t.Fatal(err)
	}

	if w.Body.String() != "one" {
		t.Error("want fist backend to be up")
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
