package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
)

type OrderRouter struct {
	orderHandler *handlers.OrderHandler
}

func NewOrderRouter(orderHandler *handlers.OrderHandler) *OrderRouter {
	return &OrderRouter{orderHandler: orderHandler}
}

func (r *OrderRouter) SetupOrderRoutes(mux *http.ServeMux) {
	mux.Handle("/orders", r.orderHandler)
	mux.Handle("/order/", r.orderHandler)
}
