package proxy

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func bodyString(r *http.Response) string {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

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
		r := httptest.NewRequest(http.MethodGet, "/asd", nil)
		b, err := p.Proxy(r)
		if err != nil {
			t.Fatal(err)
		}
		res[bodyString(b)] = res[bodyString(b)] + 1
	}

	if res["one"] == 0 || res["two"] == 0 {
		t.Errorf("proxied to only one backend: one = %d, two = %d", res["one"], res["two"])
	}

	// stop first backend
	b1.Close()
	for i := 0; i < 10; i++ {
		r := httptest.NewRequest(http.MethodGet, "/asd", nil)
		b, err := p.Proxy(r)
		if err != nil {
			t.Fatal(err)
		}

		if bodyString(b) == "one" {
			t.Error("first backend expected to be down but it's not")
		}
	}

	// stop second backend
	b2.Close()
	r := httptest.NewRequest(http.MethodGet, "/asd", nil)
	_, err = p.Proxy(r)
	if err != ErrNoBackends {
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
	go s.ListenAndServe()
	defer s.Close()

	time.Sleep(10 * time.Millisecond)

	r = httptest.NewRequest(http.MethodGet, "/asd", nil)
	b, err := p.Proxy(r)
	if err != nil {
		t.Fatal(err)
	}

	if bodyString(b) != "one" {
		t.Error("want fist backend to be up")
	}
}

func TestProxy_ProxyGzip(t *testing.T) {
	t.Parallel()

	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(http.StatusOK)

		gz := gzip.NewWriter(w)
		gz.Write([]byte("gzip"))
		gz.Close()
	})

	s := httptest.NewServer(m)
	defer s.Close()

	p, err := New(s.URL)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(http.MethodGet, "/asd", nil)
	r.Header.Set("Accept-Encoding", "gzip, deflate")
	b, err := p.Proxy(r)
	if err != nil {
		t.Fatal(err)
	}

	got := bodyString(b)
	if got != "gzip" {
		t.Errorf("gzip request body = %q, want = %q", got, "gzip")
	}
}

func TestProxy_Form(t *testing.T) {
	t.Parallel()

	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() = %v", err)
		}

		if !reflect.DeepEqual(r.Form["foo"], []string{"bar"}) {
			t.Errorf(`Form["foo"] = %v, want %v`, r.Form["foo"], []string{"bar"})
		}
	})

	s := httptest.NewServer(m)
	defer s.Close()

	p, err := New(s.URL)
	if err != nil {
		t.Fatal(err)
	}

	f := url.Values{}
	f.Set("foo", "bar")

	r := httptest.NewRequest(http.MethodPost, "/post2", strings.NewReader(f.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err = p.Proxy(r)
	if err != nil {
		t.Fatal(err)
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
