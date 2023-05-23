package reflex

import (
	"sort"

	"github.com/best4tires/kit/srv"
)

func sortLess[T any](t1, t2 T, scs srv.SortComponents) bool {
	for _, sc := range scs {
		if sc.Order == srv.SortNone {
			continue
		}
		v1, ok1 := findFieldValue(t1, sc.Name)
		v2, ok2 := findFieldValue(t2, sc.Name)
		switch {
		case !ok1 && !ok2:
			// equal regarding this sort-component
			continue
		case ok1 && !ok2:
			// consider to be less
			return sc.Order.IfLess(true)
		case !ok1 && ok2:
			// consider to be greater
			return sc.Order.IfLess(false)
		default:
			vc := CompareAny(v1, v2)
			switch {
			case vc == 0:
				// equal regarding this sort-component
				continue
			case vc < 0:
				// less
				return sc.Order.IfLess(true)
			default: // > 0
				//greater
				return sc.Order.IfLess(false)
			}
		}
	}
	// all equal
	return false
}

func Sort[T any](ts []T, scs srv.SortComponents) {
	if len(scs) == 0 {
		return
	}
	sort.SliceStable(ts, func(i, j int) bool {
		return sortLess(ts[i], ts[j], scs)
	})
}
