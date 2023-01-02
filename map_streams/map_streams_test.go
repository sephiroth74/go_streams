package map_streams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	m := map[string]string{
		"one": "this is one",
		"two": "this is two",
	}

	ok := Any(m, func(k string, v string) bool {
		return k == "two"
	})

	assert.True(t, ok)

	ok = Any(m, func(k string, v string) bool {
		return v == "hello world"
	})

	assert.False(t, ok)
}

func TestFirst(t *testing.T) {
	m := map[string]string{
		"one": "this is one",
		"two": "this is two",
	}

	key, err := First(m, func(k string, v string) (bool, error) {
		return k == "one", nil
	})

	assert.Nil(t, err)
	assert.NotNil(t, key)
	assert.Equal(t, "one", *key)
	assert.Equal(t, "this is one", m[*key])

	delete(m, *key)
	assert.Equal(t, 1, len(m))
}

func TestMap(t *testing.T) {
	m := map[string]string{
		"1": "one",
		"2": "two",
	}

	result := Map(m, func(k string, v string) TestStruct {
		return TestStruct{
			Key:   k,
			Value: v,
		}
	})

	assert.Equal(t, len(m), len(result))

	for _, v := range result {
		assert.Equal(t, m[v.Key], v.Value)
	}

}

type TestStruct struct {
	Key   string
	Value string
}
