package shared

import (
	"encoding/json"
	"net/http"
	"strings"
)

type FieldMask []string

func ParseFieldMask(r *http.Request) FieldMask {
	values := r.URL.Query()["field_mask"]

	if len(values) == 0 {
		return nil
	}

	var fields FieldMask

	for _, value := range values {
		parts := strings.Split(value, ",")

		for _, part := range parts {
			field := strings.TrimSpace(part)

			if field != "" {
				fields = append(fields, field)
			}
		}
	}

	return fields
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

func (fm FieldMask) ValidatePatch(allowedFields map[string]struct{}) *APIError {
	if fm.HasAll() {
		return &APIError{
			StatusCode: http.StatusBadRequest,
			ErrorBody: ErrorBody{
				Code:    "invalid_field_mask",
				Message: "field_mask=* is not allowed for PATCH",
				Field:   "field_mask",
			},
		}
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

func ExtractJSONFields(body []byte) (map[string]struct{}, *APIError) {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, &APIError{
			StatusCode: http.StatusBadRequest,
			ErrorBody: ErrorBody{
				Code:    "invalid_json",
				Message: "Invalid JSON body",
			},
		}
	}

	fields := make(map[string]struct{}, len(raw))

	for field := range raw {
		fields[field] = struct{}{}
	}

	return fields, nil
}

func InferPatchFieldMaskFromBody(
	bodyFields map[string]struct{},
	allowedFields map[string]struct{},
) (FieldMask, *APIError) {
	var fieldMask FieldMask

	for field := range bodyFields {
		if _, allowed := allowedFields[field]; !allowed {
			return nil, &APIError{
				StatusCode: http.StatusBadRequest,
				ErrorBody: ErrorBody{
					Code:    "invalid_patch_field",
					Message: "Field cannot be updated",
					Field:   field,
				},
			}
		}

		fieldMask = append(fieldMask, field)
	}

	return fieldMask, nil
}

func ParsePatchFieldMask(
	r *http.Request,
	body []byte,
	allowedFields map[string]struct{},
) (FieldMask, *APIError) {
	fieldMask := ParseFieldMask(r)

	bodyFields, apiErr := ExtractJSONFields(body)
	if apiErr != nil {
		return nil, apiErr
	}

	if fieldMask.IsEmpty() {
		return InferPatchFieldMaskFromBody(bodyFields, allowedFields)
	}

	if apiErr := fieldMask.ValidatePatch(allowedFields); apiErr != nil {
		return nil, apiErr
	}

	return fieldMask, nil
}
