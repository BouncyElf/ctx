package ctx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.EqualValues(
		t,
		map[string]interface{}{
			"t": "t",
		}, Map{
			"t": "t",
		},
	)
}

func TestContextRequestMethod(t *testing.T) {
	// TODO:
}

func TestContextResponseMethod(t *testing.T) {
	// TODO:
}
