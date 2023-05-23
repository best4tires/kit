package reflex

import (
	"fmt"
	"testing"

	"github.com/best4tires/kit/testutil"
)

func TestCompareAny(t *testing.T) {
	tests := []struct {
		v1  any
		v2  any
		exp int
	}{
		{v1: 42, v2: 43, exp: -1},
		{v1: 53, v2: 47, exp: 1},
		{v1: 117, v2: 117, exp: 0},

		{v1: uint8(22), v2: int64(23), exp: -1},
		{v1: uint8(22), v2: int64(245632), exp: -1},
		{v1: int64(245632), v2: uint8(22), exp: 1},
		{v1: int8(-2), v2: int64(22432553), exp: -1},

		{v1: "foo", v2: "bar", exp: 1},
		{v1: "Foo", v2: "foo", exp: -1},
		{v1: "foo", v2: "foo", exp: 0},

		{v1: 1.3, v2: 1.31, exp: -1},
		{v1: 1.3, v2: "1.31", exp: -1}, // type "float" comes before type "string"
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			res := CompareAny(test.v1, test.v2)
			testutil.AssertEqual(t, test.exp, res)
		})
	}
}
