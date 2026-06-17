package shared

import "net/http"

type FieldMask []string

func ParseFieldMask(r *http.Request) FieldMask {
	query := r.URL.Query()

	values := query["field_mask"]

	if len(values) == 0 {
		return nil
	}

	return FieldMask(values)
}

func (fm FieldMask) IsEmpty() bool {
	return len(fm) == 0
}

func (fm FieldMask) HasAll() bool {
	for _, field := range fm {
		if field == "*" {
			return true
		}
	}
	return false
}

func (fm FieldMask) ShouldReturnAll() bool {
	return fm.IsEmpty() || fm.HasAll()
}

func (fm FieldMask) Contains(field string) bool {
	for _, f := range fm {
		if f == field {
			return true
		}
	}
	return false
}

func (fm FieldMask) Validate(allowedFields map[string]struct{}) *APIError {
	if fm.ShouldReturnAll() {
		return nil
	}

	for _, field := range fm {
		if _, allowed := allowedFields[field]; !allowed {
			return &APIError{
				StatusCode: http.StatusBadRequest,
				ErrorBody: ErrorBody{
					Code:    "invalid_field_mask",
					Message: "Invalid field mask",
					Field:   "field_mask",
				},
			}
		}
	}
	return nil
}
