package monitoring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPProfCountNormal(t *testing.T) {
	count := getCount("goroutine")
	assert.IsType(t, int(1), count)
	assert.NotEmpty(t, count)

	goroutine := countGoroutine()
	assert.IsType(t, int(1), goroutine)
	assert.NotEmpty(t, goroutine)

	create := countThreadCreate()
	assert.IsType(t, int(1), create)
	assert.NotEmpty(t, create)

	heap := countHeap()
	assert.IsType(t, int(1), heap)
	assert.NotNil(t, heap)

	block := countBlock()
	assert.IsType(t, int(1), block)
	assert.NotNil(t, block)

	mutex := countMutex()
	assert.IsType(t, int(1), mutex)
	assert.NotNil(t, mutex)
}

func TestGetPProfCountNameNotFound(t *testing.T) {
	count := getCount("foo")
	assert.IsType(t, int(1), count)
	assert.Equal(t, 0, count)
}
