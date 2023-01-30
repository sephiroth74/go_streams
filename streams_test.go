package streams_test

import (
	streams "github.com/sephiroth74/go_streams"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	options := []int{2, 4, 6, 8, 10}
	result := streams.All(options, func(i int) bool {
		return i%2 == 0
	})
	assert.True(t, result)

	options = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result = streams.All(options, func(i int) bool {
		return i%2 == 0
	})

	assert.False(t, result)
}

func TestAny(t *testing.T) {
	options := []string{"Hello", "World"}
	result := streams.Any(options, func(s string) bool {
		return strings.EqualFold(s, "world")
	})
	assert.True(t, result)

	result = streams.Any(options, func(s string) bool {
		return strings.EqualFold(s, "worlds")
	})
	assert.False(t, result)
}

func TestDistinct(t *testing.T) {
	options := []string{"Hello", "World", "Another", "Hello", "World", "Another", "Test", "Another"}
	result := streams.Distinct(options)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "Hello", result[0])
	assert.Equal(t, "World", result[1])
	assert.Equal(t, "Another", result[2])
	assert.Equal(t, "Test", result[3])
}

func TestFilter(t *testing.T) {
	options := []string{"Hello", "World", "", "Hello", "World", "", "", ""}
	result := streams.Filter(options, func(i string) bool {
		return len(i) > 0
	})
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "Hello", result[0])
	assert.Equal(t, "World", result[1])
	assert.Equal(t, "Hello", result[2])
	assert.Equal(t, "World", result[3])
}
