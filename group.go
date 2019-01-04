package ctx

type GroupRouter struct {
	prefix string
	r      *router
}

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

func (g *GroupRouter) Use(hs ...Handler) {
	g.Prev(hs...)
}

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

func (g *GroupRouter) GET(path string, h Handler, mhs ...Handler) {
	g.r.push("GET", g.prefix+path, h, mhs...)
}

func (g *GroupRouter) POST(path string, h Handler, mhs ...Handler) {
	g.r.push("POST", g.prefix+path, h, mhs...)
}

func (g *GroupRouter) HEAD(path string, h Handler, mhs ...Handler) {
	g.r.push("HEAD", g.prefix+path, h, mhs...)
}

func (g *GroupRouter) OPTIONS(path string, h Handler, mhs ...Handler) {
	g.r.push("OPTIONS", g.prefix+path, h, mhs...)
}

func (g *GroupRouter) PUT(path string, h Handler, mhs ...Handler) {
	g.r.push("PUT", g.prefix+path, h, mhs...)
}

func (g *GroupRouter) PATCH(path string, h Handler, mhs ...Handler) {
	g.r.push("PATCH", g.prefix+path, h, mhs...)
}

func (g *GroupRouter) DELETE(path string, h Handler, mhs ...Handler) {
	g.r.push("DELETE", g.prefix+path, h, mhs...)
}
