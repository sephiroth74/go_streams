package map_streams

import (
	streams "github.com/sephiroth74/go_streams"
)

func Any[K comparable, T any](data map[K]T, f func(K, T) bool) bool {
	for k, v := range data {
		if f(k, v) {
			return true
		}
	}
	return false
}

func First[K comparable, T any](data map[K]T, f func(K, T) (bool, error)) (*K, error) {
	for k, v := range data {
		r, err := f(k, v)
		if err != nil {
			return nil, err
		}

		if r {
			return &k, nil
		}
	}
	return nil, streams.ErrNotFound
}

func FirstNotNull[K comparable, T any](data map[K]T, f func(K, T) bool) *K {
	for k, v := range data {
		if f(k, v) {
			return &k
		}
	}
	return nil
}

func Map[K comparable, T any, O any](data map[K]T, f func(K, T) O) []O {
	var result []O
	for k, v := range data {
		result = append(result, f(k, v))
	}
	return result
}
