package reflex

import (
	"reflect"
	"strings"

	"github.com/best4tires/kit/convert"
	"github.com/best4tires/kit/srv"
)

func Filter[T any](ts []T, fs []srv.Filter) []T {
	accept := func(t T) bool {
		for _, f := range fs {
			v, ok := findFieldValue(t, f.Name)
			if !ok {
				return false
			}
			//convert filter value to type of any
			match := false
			switch baseKind(v) {
			case reflect.Bool:
				fv := convert.ToBool(f.Value)
				switch f.Comparator {
				case srv.FilterComparatorGreater:
					match = compareBool(Convert[bool](v), fv) == 1
				case srv.FilterComparatorLess:
					match = compareBool(Convert[bool](v), fv) == -1
				case srv.FilterComparatorEqual, srv.FilterComparatorLike:
					match = compareBool(Convert[bool](v), fv) == 0
				default:
					match = false
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fv, ok := convert.ToInt(f.Value)
				if !ok {
					match = false
				}
				switch f.Comparator {
				case srv.FilterComparatorGreater:
					match = Convert[int](v) > fv
				case srv.FilterComparatorLess:
					match = Convert[int](v) < fv
				case srv.FilterComparatorEqual, srv.FilterComparatorLike:
					match = Convert[int](v) == fv
				default:
					match = false
				}
			case reflect.Float32, reflect.Float64:
				fv, ok := convert.ToFloat(f.Value)
				if !ok {
					match = false
				}
				switch f.Comparator {
				case srv.FilterComparatorGreater:
					match = Convert[float64](v) > fv
				case srv.FilterComparatorLess:
					match = Convert[float64](v) < fv
				case srv.FilterComparatorEqual, srv.FilterComparatorLike:
					match = Convert[float64](v) == fv
				default:
					match = false
				}
			default:
				//is string anyway
				sv := convert.ToString(v)
				switch f.Comparator {
				case srv.FilterComparatorGreater:
					match = sv > f.Value
				case srv.FilterComparatorLess:
					match = sv < f.Value
				case srv.FilterComparatorEqual:
					match = sv == f.Value
				case srv.FilterComparatorLike:
					match = strings.Contains(strings.ToLower(sv), strings.ToLower(f.Value))
				default:
					match = false
				}
			}
			if !match {
				return false
			}
		}
		return true
	}

	var fts []T
	for _, t := range ts {
		if accept(t) {
			fts = append(fts, t)
		}
	}
	return fts
}
