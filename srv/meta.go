package srv

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/best4tires/kit/convert"
)

// Meta holds various filter and query parameters passed with a http request
type Meta struct {
	Limit   int
	Skip    int
	Filters []Filter
	Sorts   SortComponents
}

type SortOrder string

const (
	SortNone SortOrder = "None"
	SortASC  SortOrder = "ASC"
	SortDESC SortOrder = "DESC"
)

func (o SortOrder) Valid() bool {
	return o == SortASC || o == SortDESC || o == SortNone
}

func (o SortOrder) String() string {
	return string(o)
}

func (o SortOrder) IfLess(less bool) bool {
	switch o {
	case SortASC:
		return less
	case SortDESC:
		return !less
	default:
		return less
	}
}

type SortField struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type SortConfiguration struct {
	Fields     []SortField    `json:"fields,omitempty"`
	Components SortComponents `json:"components,omitempty"`
}

type SortComponent struct {
	Name  string    `json:"name"`
	Order SortOrder `json:"order"`
}

func (sc SortComponent) String() string {
	return sc.Name + ":" + sc.Order.String()
}

type SortComponents []SortComponent

func (scs SortComponents) String() string {
	var sl []string
	for _, sc := range scs {
		sl = append(sl, sc.String())
	}
	return strings.Join(sl, "; ")
}

func (m Meta) SortOrder(name string) SortOrder {
	for _, sc := range m.Sorts {
		if sc.Name == name {
			return sc.Order
		}
	}
	return SortNone
}

func (m Meta) FilterValue(name string) interface{} {
	for _, f := range m.Filters {
		if f.Name == name {
			return f.Value
		}
	}
	return nil
}

func (m Meta) FilterValueBool(name string) bool {
	for _, f := range m.Filters {
		if f.Name == name {
			return convert.ToBool(f.Value)
		}
	}
	return false
}

func (m Meta) FilterValueString(name string) string {
	for _, f := range m.Filters {
		if f.Name == name {
			return fmt.Sprintf("%v", f.Value)
		}
	}
	return ""
}

func (m Meta) FilterValueInt(name string) int {
	for _, f := range m.Filters {
		if f.Name == name {
			v, ok := convert.ToInt(f.Value)
			if !ok {
				return 0
			} else {
				return v
			}
		}
	}
	return 0
}

func (m *Meta) ChangeFilterValueInt(name string, v int) {
	for i, f := range m.Filters {
		if f.Name == name {
			m.Filters[i].Value = fmt.Sprintf("%d", v)
			return
		}
	}
	m.Filters = append(m.Filters, Filter{
		Name:       name,
		Comparator: FilterComparatorEqual,
		Value:      fmt.Sprintf("%d", v),
	})
}

func (m Meta) FilterValueFloat(name string) float64 {
	for _, f := range m.Filters {
		if f.Name == name {
			v, ok := convert.ToFloat(f.Value)
			if !ok {
				return 0
			} else {
				return v
			}
		}
	}
	return 0
}

// FilterComparator types a comparator
type FilterComparator string

const (
	FilterComparatorEqual   FilterComparator = "eq"
	FilterComparatorLess    FilterComparator = "ls"
	FilterComparatorGreater FilterComparator = "gt"
	FilterComparatorLike    FilterComparator = "like"
)

func parseFilterComparator(c string) (FilterComparator, error) {
	switch FilterComparator(c) {
	case FilterComparatorEqual:
		return FilterComparatorEqual, nil
	case FilterComparatorLess:
		return FilterComparatorLess, nil
	case FilterComparatorGreater:
		return FilterComparatorGreater, nil
	case FilterComparatorLike:
		return FilterComparatorLike, nil
	default:
		return "", fmt.Errorf("invalid comparator")
	}
}

// Filter defines a general filter
type Filter struct {
	Name       string
	Comparator FilterComparator
	Value      string
}

func (f Filter) query() string {
	return fmt.Sprintf("filter=%s,%s,%s", f.Name, f.Comparator, f.Value)
}

func parseFilter(s string) (Filter, error) {
	sl := strings.Split(s, ",")
	if len(sl) != 3 {
		return Filter{}, fmt.Errorf("invalid filter format")
	}
	c, err := parseFilterComparator(sl[1])
	if err != nil {
		return Filter{}, err
	}
	return Filter{
		Name:       sl[0],
		Comparator: c,
		Value:      sl[2],
	}, nil
}

func (scs SortComponents) query() string {
	var sl []string
	for _, sc := range scs {
		sl = append(sl, fmt.Sprintf("%s:%s", sc.Name, sc.Order))
	}
	if len(sl) == 0 {
		return ""
	}
	return fmt.Sprintf("sort=%s", strings.Join(sl, ","))
}

// parseSort expects its values like "foo:asc,bar:desc ..."
func parseSort(s string) (SortComponents, error) {
	sl := strings.Split(s, ",")
	var cs SortComponents
	for _, c := range sl {
		csl := strings.Split(c, ":")
		if len(csl) == 0 || len(csl) > 2 {
			return SortComponents{}, fmt.Errorf("invalid sort component format %q", c)
		}
		name := csl[0]
		order := SortASC
		if len(csl) == 2 {
			order = SortOrder(csl[1])
		}
		cs = append(cs, SortComponent{
			Name:  name,
			Order: order,
		})
	}
	return cs, nil
}

func extractNumber(vs []string) (int, error) {
	if len(vs) == 0 {
		return 0, fmt.Errorf("no values")
	}
	n, err := strconv.ParseInt(vs[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// ParseMeta parses meta-data from a http request
func ParseMeta(r *http.Request) (Meta, error) {
	m := Meta{}
	var err error
	for key, values := range r.URL.Query() {
		switch key {
		case "limit":
			if m.Limit, err = extractNumber(values); err != nil {
				return m, err
			}
		case "skip":
			if m.Skip, err = extractNumber(values); err != nil {
				return m, err
			}
		case "filter":
			for _, s := range values {
				f, err := parseFilter(s)
				if err != nil {
					return m, err
				}
				m.Filters = append(m.Filters, f)
			}
		case "sort":
			if len(values) > 0 {
				m.Sorts, err = parseSort(values[0])
				if err != nil {
					return m, err
				}
			}
		}
	}
	return m, nil
}

func (m Meta) Query() string {
	var elts []string
	elts = append(elts, fmt.Sprintf("limit=%d", m.Limit))
	elts = append(elts, fmt.Sprintf("skip=%d", m.Skip))
	for _, f := range m.Filters {
		elts = append(elts, f.query())
	}
	sq := m.Sorts.query()
	if sq != "" {
		elts = append(elts, sq)
	}
	return strings.Join(elts, "&")
}
