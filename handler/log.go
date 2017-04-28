package handler

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type rw struct {
	http.ResponseWriter

	r *http.Request
	b []byte
}

func (rw *rw) WriteHeader(code int) {
	rw.ResponseWriter.WriteHeader(code)

	q := rw.r.URL.RawQuery
	if q != "" {
		q = "?" + q
	}

	log.Printf("%s %s%s %d", rw.r.Method, rw.r.URL.Path, q, code)
	if len(rw.b) > 0 {
		log.Printf("     %s", rw.b)
	}
}

type rb struct {
	r io.Reader
	c io.ReadCloser
}

func (b *rb) Read(p []byte) (int, error) {
	return b.r.Read(p)
}

func (b *rb) Close() error {
	return b.c.Close()
}

func Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		r.Body = &rb{r: bytes.NewReader(b), c: r.Body}
		h(&rw{ResponseWriter: w, r: r, b: b}, r)
	}
}
