package utils

func Contains[T comparable](val T, lst []T) bool {
	filtered := Filter(lst, func(t T) bool {
		return t == val
	})
	return len(filtered) > 0
}

func Filter[T any](lst []T, fn func(T) bool) []T {
	slice := make([]T, 0)
	return Reduce(lst, slice, func(r []T, t T) []T {
		accept := fn(t)
		if accept {
			r = append(r, t)
		}
		return r
	})
}

func Map[T any, R any](lst []T, fn func(T) R) []R {
	slice := make([]R, 0)
	return Reduce(lst, slice, func(r []R, t T) []R {
		return append(r, fn(t))
	})
}

func Reduce[T any, R any](lst []T, d R, fn func(R, T) R) R {
	val, hasLen := Head(lst)
	if !hasLen {
		return d
	}
	r := fn(d, val)
	return Reduce(Tail(lst), r, fn)
}

func Head[T any](lst []T) (T, bool) {
	var t T
	if len(lst) > 0 {
		return lst[0], true
	}
	return t, false
}

func Tail[T any](lst []T) []T {
	return lst[1:]
}
