package ctx

import "net/http"

// Handler is the alias of func(*Context) error
type Handler func(*Context) error

// Handlers is the alias of []Handler
type Handlers []Handler

// NewHttpHandler convert Handler into a http.HandlerFunc.
// NOTE: To use this method, you need realize that the context inside can not
// inherit other context. You can only use this method when this handler is the
// beginning of a handler chain or you really understand what you are doing.
func (h Handler) NewHttpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := getContext(w, r)
		err := h(ctx)
		if err != nil {
			ErrorHandler(ctx, err)
		}
	}
}

// ServeHTTP implements the http.Handler interface.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handle := h.NewHttpHandler()
	handle.ServeHTTP(w, r)
}

// Run runs the handlers. If c.abort is true, return nil.
func (hs Handlers) Run(c *Context) error {
	if c.abort {
		return nil
	}
	for _, h := range hs {
		if err := h(c); err != nil {
			return err
		}
		if c.abort {
			return nil
		}
	}
	return nil
}
