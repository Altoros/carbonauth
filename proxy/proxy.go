package proxy

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
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

var hopHeaders = map[string]bool{
	"Connection":          true,
	"Proxy-Connection":    true,
	"Keep-Alive":          true,
	"Proxy-Authenticate":  true,
	"Proxy-Authorization": true,
	"Te":                true,
	"Trailer":           true,
	"Transfer-Encoding": true,
	"Upgrade":           true,
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
	return &Proxy{backends: backends, Client: http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 50,
		},
	}}, nil
}

// ErrNoBackends is returned when all backends are currently suspended
var ErrNoBackends = errors.New("no backend are availabe")

type rc struct {
	*strings.Reader
}

func (*rc) Close() error {
	return nil
}

// ServeHTTP proxies http requests to one of backends using round-robin algorithm
func (p *Proxy) Proxy(r *http.Request) (*http.Response, error) {
	p.mu.RLock()
	if len(p.backends) == 0 {
		p.mu.RUnlock()
		return nil, ErrNoBackends
	}

	backend := p.backends[rand.Intn(len(p.backends))]
	p.mu.RUnlock()

	req := &http.Request{}
	*req = *r // copy request
	req.URL.Scheme = backend.Scheme
	req.URL.Path = backend.Path + r.URL.Path
	req.URL.Host = backend.Host
	req.RequestURI = ""
	req.Close = false

	// ParseForm() drains Body
	if r.PostForm != nil {
		frm := r.PostForm.Encode()
		req.Body = &rc{strings.NewReader(frm)}
		req.Header.Set("Content-Length", strconv.Itoa(len(frm)))
		req.ContentLength = int64(len(frm))
	}

	// http transport requests gzip encoding by
	// default and decompresses the body by itself
	req.Header.Del("Accept-Encoding")

	if ip, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := req.Header["X-Forwarded-For"]; ok {
			ip = strings.Join(prior, ", ") + ", " + ip
		}
		req.Header.Set("X-Forwarded-For", ip)
	}

	res, err := p.Client.Do(req)
	if err != nil {
		if err == ErrNoBackends {
			return nil, err
		}

		p.suspend(backend, err)
		return p.Proxy(r)
	}

	for k, _ := range hopHeaders {
		res.Header.Del(k)
	}
	return res, nil
}

// suspend temporarily removes u from the backends list
func (p *Proxy) suspend(u *url.URL, err error) {
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

	log.Printf("proxy: %s suspended, err=%v", u, err)
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
