package reflex

import (
	"reflect"
	"strings"
)

func findFieldValue(v any, field string) (any, bool) {
	rty := reflect.TypeOf(v)
	if rty.Kind() != reflect.Struct {
		return nil, false
	}
	rv := reflect.ValueOf(v)
	for i := 0; i < rty.NumField(); i++ {
		f := rty.Field(i)
		if f.Name == field {
			return rv.Field(i).Interface(), true
		}
		if jsonTag := f.Tag.Get("json"); jsonTag != "" {
			jsonName, _, _ := strings.Cut(jsonTag, ",")
			jsonName = strings.TrimSpace(jsonName)
			if jsonName == field {
				return rv.Field(i).Interface(), true
			}
		}
	}
	return nil, false
}

func baseKind(v any) reflect.Kind {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Bool:
		return reflect.Bool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.Int
	case reflect.Float32, reflect.Float64:
		return reflect.Float64
	default:
		return reflect.String
	}
}
