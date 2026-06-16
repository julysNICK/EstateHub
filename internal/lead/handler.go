package lead

import (
	"encoding/json"
	"errors"
	platformhttp "estatehub-api/internal/platform/http"
	nethttp "net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w nethttp.ResponseWriter, r *nethttp.Request) {
	var req CreateLeadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.Create(r.Context(), req)

	if err != nil {
		switch {
		case errors.Is(err, ErrTitleRequired):
			platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "title is required")
		case errors.Is(err, ErrDescriptionRequired):
			platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "description is required")
		default:
			platformhttp.ErrorJson(w, nethttp.StatusInternalServerError, ErrFailedToCreateLead.Error())
		}
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusCreated, resp)
}

func (h *Handler) List(w nethttp.ResponseWriter, r *nethttp.Request) {
	resp, err := h.service.List(r.Context())

	if err != nil {
		platformhttp.ErrorJson(w, nethttp.StatusInternalServerError, ErrFailedToGetLeads.Error())
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusOK, resp)
}

func (h *Handler) GetByID(w nethttp.ResponseWriter, r *nethttp.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "id is required")
		return
	}

	resp, err := h.service.GetById(r.Context(), id)
	if err != nil {
		platformhttp.ErrorJson(w, nethttp.StatusInternalServerError, ErrFailedToGetLead.Error())
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusOK, resp)
}

func (h *Handler) Update(w nethttp.ResponseWriter, r *nethttp.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "id is required")
		return
	}

	var req UpdateLeadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrLeadNotFound):
			platformhttp.ErrorJson(w, nethttp.StatusNotFound, "lead not found")
		default:
			platformhttp.ErrorJson(w, nethttp.StatusInternalServerError, ErrFailedToUpdateLead.Error())
		}
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusOK, resp)
}

func (h *Handler) Delete(w nethttp.ResponseWriter, r *nethttp.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		platformhttp.ErrorJson(w, nethttp.StatusBadRequest, "id is required")
		return
	}

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, ErrLeadNotFound):
			platformhttp.ErrorJson(w, nethttp.StatusNotFound, "lead not found")
		default:
			platformhttp.ErrorJson(w, nethttp.StatusInternalServerError, ErrFailedToDeleteLead.Error())
		}
		return
	}

	platformhttp.WriteJson(w, nethttp.StatusNoContent, nil)
}
