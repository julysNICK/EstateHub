package property

import (
	nethttp "net/http"
)

type PropertyHandler struct {
	service *PropertyService
}

func NewPropertyHandler(service *PropertyService) *PropertyHandler {
	return &PropertyHandler{
		service: service,
	}
}

func (h *PropertyHandler) List(w nethttp.ResponseWriter, r *nethttp.Request) {
	return
}

func (h *PropertyHandler) GetByID(w nethttp.ResponseWriter, r *nethttp.Request) {
	return
}

func (h *PropertyHandler) Create(w nethttp.ResponseWriter, r *nethttp.Request) {
	return
}

func (h *PropertyHandler) Update(w nethttp.ResponseWriter, r *nethttp.Request) {
	return
}

func (h *PropertyHandler) Delete(w nethttp.ResponseWriter, r *nethttp.Request) {
	return
}
