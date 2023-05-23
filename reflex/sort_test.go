package reflex

import (
	"fmt"
	"testing"

	"github.com/best4tires/kit/slices"
	"github.com/best4tires/kit/srv"
	"github.com/best4tires/kit/testutil"
)

func TestSort(t *testing.T) {
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
		scs srv.SortComponents
		exp []testType
	}{
		{
			in: []testType{
				mk("s4", 4, true, 4.0),
				mk("s3", 3, false, 3.0),
				mk("s2", 2, false, 2.0),
				mk("s1", 1, true, 1.0),
			},
			scs: srv.SortComponents{
				{Name: "string-value", Order: srv.SortASC},
			},
			exp: []testType{
				mk("s1", 1, true, 1.0),
				mk("s2", 2, false, 2.0),
				mk("s3", 3, false, 3.0),
				mk("s4", 4, true, 4.0),
			},
		},
		{
			in: []testType{
				mk("s2", 4, true, 4.0),
				mk("s2", 2, false, 2.0),
				mk("s1", 3, false, 1.0),
				mk("s1", 1, true, 3.0),
			},
			scs: srv.SortComponents{
				{Name: "string-value", Order: srv.SortASC},
				{Name: "FloatValue", Order: srv.SortDESC},
			},
			exp: []testType{
				mk("s1", 1, true, 3.0),
				mk("s1", 3, false, 1.0),
				mk("s2", 4, true, 4.0),
				mk("s2", 2, false, 2.0),
			},
		},
		{
			in: []testType{
				mk("s2", 4, true, 4.0),
				mk("s2", 2, false, 2.0),
				mk("s1", 3, false, 1.0),
				mk("s1", 1, true, 3.0),
			},
			scs: srv.SortComponents{
				{Name: "string-value", Order: srv.SortASC},
				{Name: "bool-value", Order: srv.SortASC},
			},
			exp: []testType{
				mk("s1", 3, false, 1.0),
				mk("s1", 1, true, 3.0),
				mk("s2", 2, false, 2.0),
				mk("s2", 4, true, 4.0),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			in := slices.Clone(test.in)
			Sort(in, test.scs)
			testutil.AssertEqual(t, test.exp, in)
		})
	}
}
