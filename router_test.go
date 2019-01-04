package ctx

import (
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUse(t *testing.T) {
	resetRouter()
	Use(func(*Context) error { return nil })
	assert.Len(t, routerIns.prev, 1)
}

func TestPrev(t *testing.T) {
	resetRouter()
	Prev(func(*Context) error { return nil })
	assert.Len(t, routerIns.prev, 1)
}

func TestNext(t *testing.T) {
	resetRouter()
	Next(func(*Context) error { return nil })
	assert.Len(t, routerIns.next, 1)
}

func TestRouterMethod(t *testing.T) {
	// TODO:
}

func resetRouter() {
	routerIns.r = httprouter.New()
	routerIns.next = nil
	routerIns.prev = nil
	routerIns.s = newServer(":8080", routerIns.r)
}
