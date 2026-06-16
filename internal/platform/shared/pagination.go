package shared

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"strconv"
)

const (
	DefaultPageSize = 20
	MaxPageSize     = 100
)

type PageRequest struct {
	PageSize  int
	PageToken string
}

type PageResult[T any] struct {
	Item          []T
	NextPageToken string
}

type pageTokenPayLoad struct {
	Offset int `json:"offset"`
}

func ParsePageRequest(rawPageSize string, rawPageToken string) (PageRequest, *APIError) {
	pageSize := DefaultPageSize

	if rawPageSize != "" {
		parsePageSize, err := strconv.Atoi(rawPageSize)

		if err != nil {
			return PageRequest{}, &APIError{
				StatusCode: nethttp.StatusBadRequest,
				ErrorBody: ErrorBody{
					Message: "invalid page size",
					Field:   "page_size",
					Code:    "invalid_page_size",
				},
			}
		}

		pageSize = parsePageSize
	}

	if pageSize < 1 || pageSize > MaxPageSize {
		return PageRequest{}, &APIError{
			StatusCode: nethttp.StatusBadRequest,
			ErrorBody:  ErrorBody{Message: fmt.Sprintf("page size must be between 1 and %d", MaxPageSize), Field: "page_size", Code: "invalid_page_size"},
		}
	}

	if rawPageToken != "" {
		if _, err := DecodePageToken(rawPageToken); err != nil {
			return PageRequest{}, &APIError{
				StatusCode: nethttp.StatusBadRequest,
				ErrorBody:  ErrorBody{Message: "invalid page token", Field: "page_token", Code: "invalid_page_token"},
			}
		}

	}

	return PageRequest{
		PageSize:  pageSize,
		PageToken: rawPageToken,
	}, nil
}

func EncodePageToken(offset int) (string, error) {
	if offset < 0 {
		return "", fmt.Errorf("invalid offset")
	}
	payload := pageTokenPayLoad{
		Offset: offset,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func DecodePageToken(token string) (int, error) {
	if token == "" {
		return 0, nil
	}

	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}

	var payload pageTokenPayLoad

	if err := json.Unmarshal(data, &payload); err != nil {
		return 0, err
	}
	if payload.Offset < 0 {
		return 0, fmt.Errorf("invalid page token")
	}

	return payload.Offset, nil
}
