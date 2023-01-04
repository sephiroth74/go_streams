package streams

func Map[I any, O any](data []I, f func(I) O) []O {
	mapped := make([]O, len(data))
	for i, e := range data {
		mapped[i] = f(e)
	}
	return mapped
}

func MapNotNull[I any, O any](data []I, f func(I) (O, error)) []O {
	var mapped []O
	for _, e := range data {
		m, err := f(e)
		if err == nil {
			mapped = append(mapped, m)
		}
	}
	return mapped
}

func Any[T any](data []T, f func(T) bool) bool {
	for _, e := range data {
		if f(e) {
			return true
		}
	}
	return false
}

func All[T any](data []T, f func(T) bool) bool {
	for _, e := range data {
		if !f(e) {
			return false
		}
	}
	return true
}

func None[T any](data []T, f func(T) bool) bool {
	for _, e := range data {
		if f(e) {
			return false
		}
	}
	return true
}

func First[T any](data []T, f func(T) bool) (*T, error) {
	for _, e := range data {
		if f(e) {
			return &e, nil
		}
	}
	return nil, ErrNotFound
}

func FirstNotNull[T any](data []T, f func(T) bool) *T {
	for _, e := range data {
		if f(e) {
			return &e
		}
	}
	return nil
}

func SumOf[T any](data []T, f func(T) int) int {
	total := 0
	for _, e := range data {
		total += f(e)
	}
	return total
}
