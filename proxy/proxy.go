package proxy

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// defaultSuspendTime is used when SuspendTime is zero
var defaultSuspendTime = 5 * time.Second

// Proxy is a multi-host http reverse proxy / load balancer
type Proxy struct {
	mu       sync.RWMutex
	backends []*url.URL

	// Client is an HTTP client
	Client http.Client

	// SuspendTime determines for how long a backend server is removed
	// from the backends pool in case of some connection error
	// if it is zero then the defaultSuspendTime value is used
	SuspendTime time.Duration
}

// New creates new proxy instance
func New(urls ...string) (*Proxy, error) {
	if len(urls) == 0 {
		return nil, errors.New("no backends")
	}

	backends := make([]*url.URL, 0, len(urls))
	for _, s := range urls {
		u, err := url.Parse(s)
		if err != nil {
			return nil, err
		}
		backends = append(backends, u)
	}
	return &Proxy{backends: backends}, nil
}

// ErrNoBackends is returned when all backends are currently suspended
var ErrNoBackends = errors.New("no backend are availabe")

// ServeHTTP proxies http requests to one of backends using round-robin algorithm
func (p *Proxy) Proxy(w http.ResponseWriter, r *http.Request) error {
	p.mu.RLock()
	if len(p.backends) == 0 {
		p.mu.RUnlock()
		return ErrNoBackends
	}

	backend := p.backends[rand.Intn(len(p.backends))]
	p.mu.RUnlock()

	u := *r.URL // copy request url
	u.Scheme = backend.Scheme
	u.Host = backend.Host
	u.Path = backend.Path + r.URL.Path

	req, err := http.NewRequest(r.Method, u.String(), r.Body)
	if err != nil {
		return err
	}

	// copy headers
	for k, h := range r.Header {
		for _, hh := range h {
			req.Header.Add(k, hh)
		}
	}

	if ip, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := req.Header["X-Forwarded-For"]; ok {
			ip = strings.Join(prior, ", ") + ", " + ip
		}
		req.Header.Set("X-Forwarded-For", ip)
	}

	res, err := p.Client.Do(req)
	if err != nil {
		if err == ErrNoBackends {
			return err
		}

		p.suspend(backend)
		return p.Proxy(w, r)
	}

	w.WriteHeader(res.StatusCode)
	_, err = io.Copy(w, res.Body)
	res.Body.Close()
	return err
}

// suspend temporarily removes u from the backends list
func (p *Proxy) suspend(u *url.URL) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var found bool
	for i, uu := range p.backends {
		if u == uu {
			p.backends[i] = p.backends[len(p.backends)-1]
			p.backends = p.backends[:len(p.backends)-1]
			found = true
			break
		}
	}

	// someone else suspended the backend before
	if !found {
		return
	}

	log.Printf("proxy: %s suspended", u)
	go func() {
		d := p.SuspendTime
		if d == 0 {
			d = defaultSuspendTime
		}

		<-time.NewTimer(d).C

		p.mu.Lock()
		p.backends = append(p.backends, u)
		p.mu.Unlock()

		log.Printf("proxy: %s is back", u)
	}()
}
