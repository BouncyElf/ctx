package ctx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	assert.Implements(t, (*http.Handler)(nil), Handler(
		func(*Context) error {
			return nil
		},
	))
}

func TestNewHttpHandler(t *testing.T) {
	assert.IsType(t, http.HandlerFunc(nil), (Handler(
		func(*Context) error {
			return nil
		},
	)).NewHttpHandler())
}

func TestHandlerChain(t *testing.T) {
	// TODO:
}
