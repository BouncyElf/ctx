package ctx

import "net/http"

// Handler is the alias of func(*Context) error
type Handler func(*Context) error

// Handlers is the alias of []Handler
type Handlers []Handler

// NewHttpHandler convert Handler into a http.HandlerFunc.
// To use this method, you need realize that the context inside can not inherit
// other context. You can only use this method when this handler is the beginning
// of a handler chain or you really understand this web framework.
func (h Handler) NewHttpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(w, r)
		err := h(ctx)
		if err != nil {
			ErrorHandler(ctx, err)
		}
	}
}

// ServeHTTP implements the http.Handler interface.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handle := h.NewHttpHandler()
	handle(w, r)
}
