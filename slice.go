package utils

func Contains[T comparable](val T, lst []T) bool {
	_, found := First(lst, func(t T) bool {
		return t == val
	})
	return found
}

func First[T any](lst []T, fn func(T) bool) (T, bool) {
	val := Reduce(lst, nil, func(r *T, t T) *T {
		if r != nil {
			return r
		}
		accept := fn(t)
		if accept {
			r = &t
		}
		return r
	})
	if val == nil {
		var t T
		return t, false
	} else {
		return *val, true
	}
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
	slice := make([]R, 0, len(lst))
	return Reduce(lst, slice, func(r []R, t T) []R {
		return append(r, fn(t))
	})
}

func Reduce[T any, R any](lst []T, d R, fn func(R, T) R) R {
	val, ok := Head(lst)
	if !ok {
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
