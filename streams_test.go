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
