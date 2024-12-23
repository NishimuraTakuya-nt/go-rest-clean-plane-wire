package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

type HealthcheckHandler struct {
	log logger.Logger
}

func NewHealthcheckHandler(log logger.Logger) *HealthcheckHandler {
	return &HealthcheckHandler{
		log: log,
	}
}

func (h *HealthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/healthcheck"):
		h.Get(w, r)
	default:
		http.NotFound(w, r)
	}
}

// Get godoc
// @Summary Health check endpoint
// @Description Get the health status of the API
// @Tags healthcheck
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} response.ErrorResponse
// @Router /healthcheck [get]
func (h *HealthcheckHandler) Get(w http.ResponseWriter, r *http.Request) {
	// healthcheck
	h.log.InfoContext(r.Context(), "Healthcheck ok")
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "healthcheck ok"})
}
