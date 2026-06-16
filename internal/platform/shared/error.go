package shared

type APIError struct {
	StatusCode int       `json:"-"`
	ErrorBody  ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func NewInvalidRequestError(message string, field string) *APIError {
	return &APIError{
		StatusCode: 400,
		ErrorBody: ErrorBody{
			Code:    "invalid_request",
			Message: message,
			Field:   field,
		},
	}
}

func NewNotFoundError(message string) *APIError {
	return &APIError{
		StatusCode: 404,
		ErrorBody: ErrorBody{
			Code:    "not_found",
			Message: message,
		},
	}
}

func NewInternalServerError(message string) *APIError {
	return &APIError{
		StatusCode: 500,
		ErrorBody: ErrorBody{
			Code:    "internal_server_error",
			Message: message,
		},
	}
}

func NewConflictError(message string) *APIError {
	return &APIError{
		StatusCode: 409,
		ErrorBody: ErrorBody{
			Code:    "conflict",
			Message: message,
		},
	}
}

func NewUnauthorizedError(message string) *APIError {
	return &APIError{
		StatusCode: 401,
		ErrorBody: ErrorBody{
			Code:    "unauthorized",
			Message: message,
		},
	}
}
