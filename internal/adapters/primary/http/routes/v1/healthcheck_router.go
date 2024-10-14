package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
)

type HealthcheckRouter struct {
	healthcheckHandler *handlers.HealthcheckHandler
}

func NewHealthcheckRouter(healthcheckHandler *handlers.HealthcheckHandler) *HealthcheckRouter {
	return &HealthcheckRouter{healthcheckHandler: healthcheckHandler}
}

func (r *HealthcheckRouter) SetupHealthcheckRoutes(mux *http.ServeMux) {
	mux.Handle("/healthcheck", r.healthcheckHandler)
}
