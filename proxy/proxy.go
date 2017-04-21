package proxy

import (
	"errors"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Proxy is a multi-host reverse proxy / load balancer
type Proxy struct {
	num      int
	backends []*httputil.ReverseProxy
}

// New creates new proxy instance
func New(urls ...string) (*Proxy, error) {
	if len(urls) == 0 {
		return nil, errors.New("no backends")
	}

	backends := make([]*httputil.ReverseProxy, 0, len(urls))
	for _, s := range urls {
		u, err := url.Parse(s)
		if err != nil {
			return nil, err
		}
		backends = append(backends, httputil.NewSingleHostReverseProxy(u))
	}
	return &Proxy{backends: backends, num: len(backends)}, nil
}

// TODO: update backends state
// ServeHTTP proxies http requests to one of backends using round-robin algorithm
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.backends[rand.Intn(p.num)].ServeHTTP(w, r)
}
