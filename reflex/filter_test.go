package reflex

import (
	"fmt"
	"testing"

	"github.com/best4tires/kit/srv"
	"github.com/best4tires/kit/testutil"
)

func TestFilter(t *testing.T) {
	type testType struct {
		StringValue string  `json:"string-value"`
		IntValue    int     `json:"int-value"`
		BoolValue   bool    `json:"bool-value"`
		FloatValue  float64 `json:"float-value"`
	}
	mk := func(s string, n int, b bool, f float64) testType {
		return testType{s, n, b, f}
	}

	tests := []struct {
		in  []testType
		fs  []srv.Filter
		exp []testType
	}{
		{
			in: []testType{
				mk("s4", 4, true, 4.0),
				mk("s3", 3, false, 3.0),
				mk("s2", 2, false, 2.0),
				mk("s1", 1, true, 1.0),
			},
			fs: []srv.Filter{
				{Name: "string-value", Comparator: srv.FilterComparatorEqual, Value: "s2"},
			},
			exp: []testType{
				mk("s2", 2, false, 2.0),
			},
		},
		{
			in: []testType{
				mk("s4", 4, true, 4.0),
				mk("s3", 3, false, 3.0),
				mk("s2", 2, false, 2.0),
				mk("s1", 1, true, 1.0),
			},
			fs: []srv.Filter{
				{Name: "float-value", Comparator: srv.FilterComparatorGreater, Value: "2.1"},
			},
			exp: []testType{
				mk("s4", 4, true, 4.0),
				mk("s3", 3, false, 3.0),
			},
		},
		{
			in: []testType{
				mk("s4", 4, true, 4.0),
				mk("s3", 3, false, 3.0),
				mk("s2", 2, false, 2.0),
				mk("s1", 1, true, 1.0),
			},
			fs: []srv.Filter{
				{Name: "float-value", Comparator: srv.FilterComparatorGreater, Value: "1.9"},
				{Name: "bool-value", Comparator: srv.FilterComparatorEqual, Value: "false"},
			},
			exp: []testType{
				mk("s3", 3, false, 3.0),
				mk("s2", 2, false, 2.0),
			},
		},
		{
			in: []testType{
				mk("s4", 1, true, 4.0),
				mk("s3", 2, false, 3.0),
				mk("s2", 3, false, 2.0),
				mk("s1", 4, true, 1.0),
			},
			fs: []srv.Filter{
				{Name: "float-value", Comparator: srv.FilterComparatorGreater, Value: "1.9"},
				{Name: "int-value", Comparator: srv.FilterComparatorLess, Value: "3"},
			},
			exp: []testType{
				mk("s4", 1, true, 4.0),
				mk("s3", 2, false, 3.0),
			},
		},
		{
			in: []testType{
				mk("s4", 1, true, 4.0),
				mk("s3", 2, false, 3.0),
				mk("s2", 3, false, 2.0),
				mk("s1", 4, true, 1.0),
			},
			fs: []srv.Filter{
				{Name: "float-value", Comparator: srv.FilterComparatorGreater, Value: "1.9"},
				{Name: "int-value", Comparator: srv.FilterComparatorLess, Value: "3"},
				{Name: "bool-value", Comparator: srv.FilterComparatorEqual, Value: "true"},
			},
			exp: []testType{
				mk("s4", 1, true, 4.0),
			},
		},
		{
			in: []testType{
				mk("s4", 1, true, 4.0),
				mk("s3", 2, false, 3.0),
				mk("s2", 3, false, 2.0),
				mk("s1", 4, true, 1.0),
				mk("seb4", 1, true, 4.0),
				mk("se3", 2, false, 3.0),
				mk("sb22", 3, false, 2.0),
				mk("sebe1", 4, true, 1.0),
			},
			fs: []srv.Filter{
				{Name: "float-value", Comparator: srv.FilterComparatorGreater, Value: "1.9"},
				{Name: "int-value", Comparator: srv.FilterComparatorLess, Value: "3"},
				{Name: "bool-value", Comparator: srv.FilterComparatorEqual, Value: "true"},
			},
			exp: []testType{
				mk("s4", 1, true, 4.0),
				mk("seb4", 1, true, 4.0),
			},
		},
		{
			in: []testType{
				mk("s4", 1, true, 4.0),
				mk("s3", 2, false, 3.0),
				mk("s2", 3, false, 2.0),
				mk("s1", 4, true, 1.0),
				mk("seb4", 1, true, 4.0),
				mk("se3", 2, false, 3.0),
				mk("sb22", 3, false, 2.0),
				mk("sebe1", 4, true, 1.0),
			},
			fs: []srv.Filter{
				{Name: "float-value", Comparator: srv.FilterComparatorGreater, Value: "1.9"},
				{Name: "int-value", Comparator: srv.FilterComparatorLess, Value: "3"},
				{Name: "bool-value", Comparator: srv.FilterComparatorEqual, Value: "true"},
				{Name: "string-value", Comparator: srv.FilterComparatorLike, Value: "EB"},
			},
			exp: []testType{
				mk("seb4", 1, true, 4.0),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Filter(test.in, test.fs)
			testutil.AssertEqual(t, test.exp, res)
		})
	}
}
