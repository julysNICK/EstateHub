package http

import (
	"context"
	"database/sql"
	nethttp "net/http"
	"time"
)

type HealthHandler struct {
	DB               *sql.DB
	ReadinessTimeout time.Duration
}

func NewHealthHandler(db *sql.DB, readinessTimeout time.Duration) *HealthHandler {
	return &HealthHandler{
		DB:               db,
		ReadinessTimeout: readinessTimeout,
	}
}

func (h *HealthHandler) Healthz(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != "GET" {
		ErrorJson(w, 405, "Method not allowed")
		return
	}

	WriteJson(w, 200, map[string]string{"status": "ok"})
}

func (h *HealthHandler) Readyz(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != "GET" {
		ErrorJson(w, 405, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.ReadinessTimeout)

	defer cancel()

	if err := h.DB.PingContext(ctx); err != nil {
		WriteJson(w, nethttp.StatusServiceUnavailable, map[string]any{
			"status": "not_ready",
			"checks": map[string]string{
				"database": "unavailable",
			},
		})
		return
	}

	WriteJson(w, nethttp.StatusOK, map[string]any{
		"status": "ready",
		"checks": map[string]string{
			"database": "available",
		},
	})
}
