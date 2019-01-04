package ctx

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var routerIns *router

func init() {
	routerIns = new(router)
}

// Router is the http router in ctx. It's the entry of your app.
type router struct {
	r    *httprouter.Router
	prev Handlers
	next Handlers
	s    *server
}

// Run runs the app, default port is '8080'.
func Run(addr ...string) {
	port := ":8080"
	if len(addr) != 0 {
		port = addr[0]
	}
	if routerIns.r == nil {
		log.Fatalf("%s nil router\n", "[ctx]")
	}
	log.Printf("%s listen at%s.\n", "[ctx]", port)
	if routerIns.s == nil {
		routerIns.s = newServer(port, routerIns.r)
	}
	if err := routerIns.s.s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("%s server error: %v\n", "[ctx]", err)
	}
}

// Shutdown shutdown the server gracefully, when t <= 0, it wait for all request
// finished. Othercase it will shutdown right after t.
func Shutdown(t time.Duration) {
	if t <= 0 {
		routerIns.s.s.Shutdown(context.Background())
		return
	}
	c, cancel := context.WithTimeout(context.Background(), t)
	routerIns.s.s.Shutdown(c)
	cancel()
}

// Use is a alias of Prev, it register `hs` as a banch of prev handler.
// `hs` will be execute before the handler.
func Use(hs ...Handler) {
	Prev(hs...)
}

// Prev register `hs` as a banch of prev handler. `hs` will be execute before
// the handler.
func Prev(hs ...Handler) {
	if len(hs) == 0 {
		return
	}
	if routerIns.prev == nil {
		routerIns.prev = hs
		return
	}
	routerIns.prev = append(routerIns.prev, hs...)
}

// Next register `hs` as a banch of next handler. `hs` will be execute after the
// handler.
func Next(hs ...Handler) {
	if len(hs) == 0 {
		return
	}
	if routerIns.next == nil {
		routerIns.next = hs
		return
	}
	routerIns.next = append(routerIns.next, hs...)
}

// GET register a router with method "GET".
func GET(path string, h Handler, mhs ...Handler) {
	routerIns.push("GET", path, h, mhs...)
}

// POST register a router with method "POST".
func POST(path string, h Handler, mhs ...Handler) {
	routerIns.push("POST", path, h, mhs...)
}

// HEAD register a router with method "HEAD".
func HEAD(path string, h Handler, mhs ...Handler) {
	routerIns.push("HEAD", path, h, mhs...)
}

// OPTIONS register a router with method "OPTIONS".
func OPTIONS(path string, h Handler, mhs ...Handler) {
	routerIns.push("OPTIONS", path, h, mhs...)
}

// PUT register a router with method "PUT".
func PUT(path string, h Handler, mhs ...Handler) {
	routerIns.push("PUT", path, h, mhs...)
}

// PATCH register a router with method "PATCH".
func PATCH(path string, h Handler, mhs ...Handler) {
	routerIns.push("PATCH", path, h, mhs...)
}

// DELETE register a router with method "DELETE".
func DELETE(path string, h Handler, mhs ...Handler) {
	routerIns.push("DELETE", path, h, mhs...)
}

// push register router with httprouter's method `(*httprouter.Router).Handler`.
func (r *router) push(method, path string, h Handler, mhs ...Handler) {
	if routerIns.r == nil {
		routerIns.r = httprouter.New()
	}
	routerIns.r.Handler(method, path, Handler(
		func(c *Context) error {
			if err := r.prev.Run(c); err != nil {
				ErrorHandler(c, err)
				return nil
			}
			if err := Handlers(mhs).Run(c); err != nil {
				ErrorHandler(c, err)
				return nil
			}
			if err := Handlers([]Handler{h}).Run(c); err != nil {
				ErrorHandler(c, err)
				return nil
			}
			if err := r.next.Run(c); err != nil {
				ErrorHandler(c, err)
				return nil
			}
			return nil
		},
	))
}
