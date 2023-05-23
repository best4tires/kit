package reflex

import (
	"github.com/best4tires/kit/slices"
	"github.com/best4tires/kit/srv"
)

func Query[T any](ts []T, meta srv.Meta) []T {
	qts := slices.Clone(ts)
	// first sort
	Sort(qts, meta.Sorts)
	//filter
	qts = Filter(qts, meta.Filters)
	// limit offset
	if meta.Skip > 0 {
		qts = qts[meta.Skip:]
	}
	if meta.Limit > 0 && len(qts) > meta.Limit {
		qts = qts[:meta.Limit]
	}
	return qts
}
