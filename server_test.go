package ctx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	s := newServer("", nil)
	assert.Equal(t, s.s.Addr, ":8080")
}
