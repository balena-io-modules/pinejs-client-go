package pinejs

import "fmt"

// FilterType specifies the OData filter you wish to use.
type FilterType int

const (
	Expand FilterType = iota
	Filter
	Select
)

// String returns the OData $-prefixed name for the filter.
func (ft FilterType) String() string {
	switch ft {
	case Expand:
		return "$expand"
	case Filter:
		return "$filter"
	case Select:
		return "$select"
	}

	panic(fmt.Sprintf("Unrecognised filter type %d", ft))
}

type ODataFilter struct {
	Type    FilterType
	Content []string
}

// Filters is a collection of OData filters.
type Filters []ODataFilter

func (fs Filters) toMap() map[string][]string {
	ret := make(map[string][]string)

	for _, f := range fs {
		name := f.Type.String()
		ret[name] = append(ret[name], f.Content...)
	}

	return ret
}

func parseFilter(filter, val interface{}) ODataFilter {
	var strs []string

	switch v := val.(type) {
	case string:
		strs = []string{v}
	case []string:
		strs = v
	case nil:
		panic("invalid type")
	}

	return ODataFilter{filter.(FilterType), strs}
}

// Filterfy is a convenience function for inputting filters.
//
// Use it where Filters are expected as Filterfy(pinejs.Expand, []string {"foo", "bar"},
// pinejs.Select, "bar", etc.)
//
// Values can either be specified as a string or an array of strings.
func Filterfy(pairs ...interface{}) Filters {
	if len(pairs) < 2 {
		return nil
	}

	var ret Filters
	for i := 0; i < len(pairs)-1; i += 2 {
		ret = append(ret, parseFilter(pairs[i], pairs[i+1]))
	}

	return ret
}
