package ctx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var h = func(c *Context) error {
	return c.String("1")
}

func TestGroupRouterPrefix(t *testing.T) {
	resetRouter()
	g1 := Group("/g1")
	g2 := g1.Group("/g2")
	assert.Equal(t, g1.prefix, "/g1")
	assert.Equal(t, g2.prefix, "/g1/g2")

	g1.GET("/test", h)
	g2.GET("/test", h)

	req := httptest.NewRequest(http.MethodGet, "/g1/test", nil)
	res := httptest.NewRecorder()
	routerIns.r.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)

	req = httptest.NewRequest(http.MethodGet, "/g1/g2/test", nil)
	res = httptest.NewRecorder()
	routerIns.r.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)

	req = httptest.NewRequest(http.MethodGet, "/test", nil)
	res = httptest.NewRecorder()
	routerIns.r.ServeHTTP(res, req)
	assert.Equal(t, 404, res.Code)
}

func TestGroupRouterMiddleware(t *testing.T) {
	resetRouter()

	g := Group("/g")
	g.GET("/test", h)

	Prev(h)
	assert.Len(t, g.r.prev, 0)
	assert.Len(t, routerIns.prev, 1)

	g.Use(h)
	assert.Len(t, g.r.prev, 1)
	assert.Len(t, routerIns.prev, 1)
}

func TestGroupRouterMethod(t *testing.T) {
	// TODO:
}
