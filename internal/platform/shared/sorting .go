package shared

const (
	DefaultSortField = "created_at"
	DefaultSortOrder = "desc"
)

type SortRequest struct {
	Sort      string
	Direction string
	Column    string
}

func ParsedSortRequest(rawSort string, rawDirection string, allowedSorts map[string]string) (SortRequest, *APIError) {
	sort := rawSort
	if sort == "" {
		sort = DefaultSortField
	}

	column, ok := allowedSorts[sort]

	if !ok {
		return SortRequest{}, &APIError{
			StatusCode: 400,
			ErrorBody: ErrorBody{
				Code:    "invalid_sort",
				Message: "Invalid sort field",
				Field:   "sort",
			},
		}
	}

	direction := rawDirection

	if direction == "" {
		direction = DefaultSortOrder
	}

	if direction != "asc" && direction != "desc" {
		return SortRequest{}, &APIError{
			StatusCode: 400,
			ErrorBody: ErrorBody{
				Code:    "invalid_sort_direction",
				Message: "Invalid sort direction",
				Field:   "direction",
			},
		}
	}

	return SortRequest{
		Sort:      sort,
		Direction: direction,
		Column:    column,
	}, nil

}
