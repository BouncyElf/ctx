package ctx

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Router is the http router in ctx. It's the entry of your app.
type Router struct {
	r    *httprouter.Router
	prev Handlers
	next Handlers
}

// New returns a new router.
func New() *Router {
	return &Router{
		r:    httprouter.New(),
		prev: nil,
		next: nil,
	}
}

// Run runs the app.
func (r *Router) Run(addr ...string) {
	port := ":8080"
	if len(addr) != 0 {
		port = addr[0]
	}
	log.Printf("%s listen at%s.\n", "[ctx]", port)
	http.ListenAndServe(port, r.r)
}

// Use is a alias of Prev, it register `hs` as a banch of prev handler.
// `hs` will be execute before the handler.
func (r *Router) Use(hs ...Handler) {
	r.Prev(hs...)
}

// Prev register `hs` as a banch of prev handler. `hs` will be execute before
// the handler.
func (r *Router) Prev(hs ...Handler) {
	if len(hs) == 0 {
		return
	}
	if r.prev == nil {
		r.prev = hs
		return
	}
	r.prev = append(r.prev, hs...)
}

// Next register `hs` as a banch of next handler. `hs` will be execute after the
// handler.
func (r *Router) Next(hs ...Handler) {
	if len(hs) == 0 {
		return
	}
	if r.next == nil {
		r.next = hs
		return
	}
	r.next = append(r.next, hs...)
}

// GET register a router with method "GET".
func (r *Router) GET(path string, h Handler, mhs ...Handler) {
	r.push("GET", path, h, mhs...)
}

// POST register a router with method "POST".
func (r *Router) POST(path string, h Handler, mhs ...Handler) {
	r.push("POST", path, h, mhs...)
}

// HEAD register a router with method "HEAD".
func (r *Router) HEAD(path string, h Handler, mhs ...Handler) {
	r.push("HEAD", path, h, mhs...)
}

// OPTIONS register a router with method "OPTIONS".
func (r *Router) OPTIONS(path string, h Handler, mhs ...Handler) {
	r.push("OPTIONS", path, h, mhs...)
}

// PUT register a router with method "PUT".
func (r *Router) PUT(path string, h Handler, mhs ...Handler) {
	r.push("PUT", path, h, mhs...)
}

// PATCH register a router with method "PATCH".
func (r *Router) PATCH(path string, h Handler, mhs ...Handler) {
	r.push("PATCH", path, h, mhs...)
}

// DELETE register a router with method "DELETE".
func (r *Router) DELETE(path string, h Handler, mhs ...Handler) {
	r.push("DELETE", path, h, mhs...)
}

// push register router with httprouter's method `(*httprouter.Router).Handler`.
func (r *Router) push(method, path string, h Handler, mhs ...Handler) {
	r.r.Handler(method, path, Handler(
		func(c *Context) error {
			for _, h := range r.prev {
				if err := h(c); err != nil {
					ErrorHandler(c, err)
					return nil
				}
			}
			for _, h := range mhs {
				if err := h(c); err != nil {
					ErrorHandler(c, err)
					return nil
				}
			}
			if err := h(c); err != nil {
				ErrorHandler(c, err)
				return nil
			}
			for _, h := range r.next {
				if err := h(c); err != nil {
					ErrorHandler(c, err)
					return nil
				}
			}
			return nil
		},
	))
}
