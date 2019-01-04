package ctx

import "net/http"

type server struct {
	s *http.Server
}

func newServer(addr string, h http.Handler) *server {
	if addr == "" {
		addr = ":8080"
	}
	s := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	return &server{
		s: s,
	}
}

// TODO: add server conf
