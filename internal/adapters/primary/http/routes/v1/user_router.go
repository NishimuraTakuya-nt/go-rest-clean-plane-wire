package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
)

type UserRouter struct {
	userHandler *handlers.UserHandler
}

func NewUserRouter(userHandler *handlers.UserHandler) *UserRouter {
	return &UserRouter{userHandler: userHandler}
}

func (r *UserRouter) SetupUserRoutes(mux *http.ServeMux) {
	mux.Handle("/users", r.userHandler)
	mux.Handle("/user", r.userHandler)
}
