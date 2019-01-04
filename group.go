package ctx

// GroupRouter is a set of routes with common prefix and middleware.
type GroupRouter struct {
	prefix string
	r      *router
}

// Group returns a new GroupRouter with prefix and optinal middleware.
func Group(prefix string, hs ...Handler) *GroupRouter {
	r := &router{
		r:    routerIns.r,
		prev: routerIns.prev,
		next: routerIns.next,
		s:    nil,
	}
	return &GroupRouter{
		prefix: prefix,
		r:      r,
	}
}

// g.Group returns a new GroupRouter with prefix and
// optinal middleware based on g.
func (g *GroupRouter) Group(prefix string, hs ...Handler) *GroupRouter {
	r := &router{
		r:    routerIns.r,
		prev: g.r.prev,
		next: g.r.next,
		s:    nil,
	}
	return &GroupRouter{
		prefix: g.prefix + prefix,
		r:      r,
	}
}

// g.Use is same as Use, it register `hs` as a bench of prev handler in g.
func (g *GroupRouter) Use(hs ...Handler) {
	g.Prev(hs...)
}

// g.Prev is same as Prev, it register `hs` as a bench of prev handler in g.
func (g *GroupRouter) Prev(hs ...Handler) {
	if len(hs) == 0 {
		return
	}
	if g.r.prev == nil {
		g.r.prev = hs
		return
	}
	g.r.prev = append(g.r.prev, hs...)
}

// g.Next is same as Next, it register `hs` as a bench of next handler in g.
func (g *GroupRouter) Next(hs ...Handler) {
	if len(hs) == 0 {
		return
	}
	if g.r.next == nil {
		g.r.next = hs
		return
	}
	g.r.next = append(g.r.next, hs...)
}

// g.GET is same as GET, it register a GET route with g.prefix+path.
func (g *GroupRouter) GET(path string, h Handler, mhs ...Handler) {
	g.r.push("GET", g.prefix+path, h, mhs...)
}

// g.POST is same as POST, it register a POST route with g.prefix+path.
func (g *GroupRouter) POST(path string, h Handler, mhs ...Handler) {
	g.r.push("POST", g.prefix+path, h, mhs...)
}

// g.HEAD is same as HEAD, it register a HEAD route with g.prefix+path.
func (g *GroupRouter) HEAD(path string, h Handler, mhs ...Handler) {
	g.r.push("HEAD", g.prefix+path, h, mhs...)
}

// g.OPTIONS is same as OPTIONS, it register a OPTIONS route with g.prefix+path.
func (g *GroupRouter) OPTIONS(path string, h Handler, mhs ...Handler) {
	g.r.push("OPTIONS", g.prefix+path, h, mhs...)
}

// g.PUT is same as PUT, it register a PUT route with g.prefix+path.
func (g *GroupRouter) PUT(path string, h Handler, mhs ...Handler) {
	g.r.push("PUT", g.prefix+path, h, mhs...)
}

// g.PATCH is same as PATCH, it register a PATCH route with g.prefix+path.
func (g *GroupRouter) PATCH(path string, h Handler, mhs ...Handler) {
	g.r.push("PATCH", g.prefix+path, h, mhs...)
}

// g.DELETE is same as DELETE, it register a DELETE route with g.prefix+path.
func (g *GroupRouter) DELETE(path string, h Handler, mhs ...Handler) {
	g.r.push("DELETE", g.prefix+path, h, mhs...)
}
