package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
)

type AuthRouter struct {
	authHandler *handlers.AuthHandler
}

func NewAuthRouter(authHandler *handlers.AuthHandler) *AuthRouter {
	return &AuthRouter{authHandler: authHandler}
}

func (r *AuthRouter) SetupAuthRoutes(mux *http.ServeMux) {
	mux.Handle("/auth/", r.authHandler)
}
