package pinejs

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
