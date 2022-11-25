package slices

func Make[T any]() T {
	var t T
	return t
}

func Clone[T any](ts []T) []T {
	cts := make([]T, len(ts))
	copy(cts, ts)
	return cts
}

func ContainsFnc[T any](ts []T, fnc func(t T) bool) bool {
	for _, t := range ts {
		if fnc(t) {
			return true
		}
	}
	return false
}

func Contains[T comparable](ts []T, ct T) bool {
	for _, t := range ts {
		if t == ct {
			return true
		}
	}
	return false
}

func FindFnc[T any](ts []T, fnc func(t T) bool) (T, bool) {
	for _, t := range ts {
		if fnc(t) {
			return t, true
		}
	}
	return Make[T](), false
}

func FirstIndexOf[T comparable](ts []T, t T) int {
	for i, et := range ts {
		if et == t {
			return i
		}
	}
	return -1
}

func FirstIndexOfFnc[T any](ts []T, fnc func(t T) bool) int {
	for i, t := range ts {
		if fnc(t) {
			return i
		}
	}
	return -1
}

func RemoveIdx[T any](ts []T, idx int) []T {
	if idx < 0 || idx >= len(ts) {
		return ts
	}
	var rts []T
	for i, t := range ts {
		if i == idx {
			continue
		}
		rts = append(rts, t)
	}

	return rts
}

func RemoveFirst[T comparable](ts []T, t T) []T {
	idx := FirstIndexOf(ts, t)
	if idx < 0 {
		return ts
	}
	return RemoveIdx(ts, idx)
}

func RemoveFirstFnc[T any](ts []T, fnc func(t T) bool) []T {
	idx := FirstIndexOfFnc(ts, fnc)
	if idx < 0 {
		return ts
	}
	return RemoveIdx(ts, idx)
}

func Dedup[T comparable](ts []T) []T {
	var dts []T
	for _, t := range ts {
		if !Contains(dts, t) {
			dts = append(dts, t)
		}
	}
	return dts
}

func Map[T any, V any](ts []T, conv func(T) V) []V {
	vs := make([]V, len(ts))
	for i, t := range ts {
		vs[i] = conv(t)
	}
	return vs
}

func Repeat[T any](t T, count int) []T {
	ts := make([]T, count)
	for i := 0; i < count; i++ {
		ts[i] = t
	}
	return ts
}
