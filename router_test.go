package ctx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUse(t *testing.T) {
	r := New()
	r.Use(func(*Context) error { return nil })
	assert.Len(t, r.prev, 1)
}

func TestPrev(t *testing.T) {
	r := New()
	r.Prev(func(*Context) error { return nil })
	assert.Len(t, r.prev, 1)
}

func TestNext(t *testing.T) {
	r := New()
	r.Next(func(*Context) error { return nil })
	assert.Len(t, r.next, 1)
}
