package user

import (
	"encoding/json"
	"errors"
	platformhttp "estatehub-api/internal/platform/http"
	"estatehub-api/internal/platform/shared"
	"io"
	nethttp "net/http"
	"strings"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(w nethttp.ResponseWriter, r *nethttp.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.Create(r.Context(), req)
	if err != nil {
		handleUserError(w, err)
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusCreated, resp)
}

func (h *UserHandler) List(w nethttp.ResponseWriter, r *nethttp.Request) {
	query := r.URL.Query()

	pageToken := query.Get("page_token")
	pageSize := query.Get("page_size")

	req, apiError := shared.ParsePageRequest(
		pageSize,
		pageToken,
	)

	if apiError != nil {
		platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, apiError)
		return
	}

	filter := ListUserRequest{
		Name:        query.Get("name"),
		Email:       query.Get("email"),
		PageRequest: req,
	}

	sort := query.Get("sort")
	direction := query.Get("direction")

	sortRequest, errSort := shared.ParsedSortRequest(sort, direction, allowedUserSorts)

	if errSort != nil {
		platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, errSort)
		return
	}

	resp, err := h.service.List(r.Context(), filter, sortRequest)

	if err != nil {
		platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, err)
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusOK, resp)
}

func (h *UserHandler) GetByID(w nethttp.ResponseWriter, r *nethttp.Request) {
	/* 	resp, err := h.service.GetByID(r.Context(), r.PathValue("id")) */

	ctx := r.Context()
	userID := r.PathValue("id")

	userm, err := h.service.GetByID(ctx, userID)
	if err != nil {
		handleUserError(w, err)
		return
	}

	fieldMask := shared.ParseFieldMask(r)

	resp, errMask := ApplyUserFieldMask(&userm, fieldMask)

	if errMask != nil {
		platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, errMask)
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusOK, resp)
}

func (h *UserHandler) Update(w nethttp.ResponseWriter, r *nethttp.Request) {
	var req UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ctx := r.Context()

		userID := strings.TrimPrefix(r.URL.Path, "/v1/users/")
		if userID == "" {
			platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, &shared.APIError{
				StatusCode: nethttp.StatusBadRequest,
				ErrorBody: shared.ErrorBody{
					Code:    "missing_user_id",
					Message: "Missing user id",
					Field:   "user_id",
				},
			})
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, &shared.APIError{
				StatusCode: nethttp.StatusBadRequest,
				ErrorBody: shared.ErrorBody{
					Code:    "invalid_body",
					Message: "Invalid request body",
				},
			})
			return
		}

		fieldMask, apiErr := shared.ParsePatchFieldMask(
			r,
			body,
			allowedUserPatchFieldMask,
		)
		if apiErr != nil {
			platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, apiErr)
			return
		}

		var req UpdateUserRequest
		if err := json.Unmarshal(body, &req); err != nil {
			platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, &shared.APIError{
				StatusCode: nethttp.StatusBadRequest,
				ErrorBody: shared.ErrorBody{
					Code:    "invalid_json",
					Message: "Invalid JSON body",
				},
			})
			return
		}

		user, apiErr := h.service.Update(ctx, userID, req, fieldMask)
		if apiErr != nil {
			platformhttp.ErrorJsonV2(w, nethttp.StatusBadRequest, apiErr)
			return
		}

		platformhttp.WriteJson(w, nethttp.StatusOK, user)
	}
}

func (h *UserHandler) Delete(w nethttp.ResponseWriter, r *nethttp.Request) {
	err := h.service.Delete(r.Context(), r.PathValue("id"))
	if err != nil {
		handleUserError(w, err)
		return
	}

	w.WriteHeader(nethttp.StatusNoContent)
}

func handleUserError(w nethttp.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrInvalidID),
		errors.Is(err, ErrInvalidName),
		errors.Is(err, ErrInvalidEmail),
		errors.Is(err, ErrInvalidPassword),
		errors.Is(err, ErrInvalidUserType),
		errors.Is(err, ErrInvalidCNPJ),
		errors.Is(err, ErrInvalidCPF),
		errors.Is(err, ErrInvalidCRECI):
		platformhttp.ErrorJson(w, nethttp.StatusBadRequest, err.Error())

	case errors.Is(err, ErrUserNotFound):
		platformhttp.ErrorJson(w, nethttp.StatusNotFound, err.Error())

	case errors.Is(err, ErrEmailAlreadyUsed):
		platformhttp.ErrorJson(w, nethttp.StatusConflict, err.Error())

	default:
		platformhttp.ErrorJson(w, nethttp.StatusInternalServerError, "internal server error")
	}
}
