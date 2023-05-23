package reflex

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"
)

func ZeroOf[T any]() T {
	var t T
	return t
}

func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

// CanConvert checks if v can be converted to T
func CanConvert[T any](v any) bool {
	rv := reflect.ValueOf(v)
	rty := TypeOf[T]()
	return rv.CanConvert(rty)
}

// Convert converts v to to T. Convert Panics if convert fails
func Convert[T any](v any) T {
	return reflect.ValueOf(v).Convert(TypeOf[T]()).Interface().(T)
}

func Compare[T constraints.Ordered](t1, t2 T) int {
	switch {
	case t1 < t2:
		return -1
	case t1 > t2:
		return 1
	default:
		return 0
	}
}

func compareBool(b1, b2 bool) int {
	switch {
	case !b1 && b2:
		return -1
	case b1 && !b2:
		return 1
	default:
		return 0
	}
}

// Compare compares to any-values and returns:
// strings.Compare of type-kind-names, if types of v1 and v2 are not equal
// 0: if values are equal; -1: if v1 < v2; 1: if v1 > v2
func CompareAny(v1, v2 any) int {
	ty1 := baseKind(v1)
	ty2 := baseKind(v2)
	if ty1 != ty2 {
		return strings.Compare(ty1.String(), ty2.String())
	}
	switch ty1 {
	case reflect.Bool:
		return compareBool(Convert[bool](v1), Convert[bool](v2))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Compare(Convert[int](v1), Convert[int](v2))
	case reflect.Float32, reflect.Float64:
		return Compare(Convert[float64](v1), Convert[float64](v2))
	default:
		return strings.Compare(fmt.Sprintf("%v", v1), fmt.Sprintf("%v", v2))
	}
}
