package ctx

import (
	"sync"
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

func TestEnsure(t *testing.T) {
	counter, c, wg := 0, getContext(nil, nil), new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Ensure(func() {
				counter++
			})
		}()
	}
	wg.Wait()
	assert.Equal(t, 10, counter)
}
