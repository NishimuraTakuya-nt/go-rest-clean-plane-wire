package v1

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
)

type ProductRouter struct {
	productHandler *handlers.ProductHandler
}

func NewProductRouter(productHandler *handlers.ProductHandler) *ProductRouter {
	return &ProductRouter{productHandler: productHandler}
}

func (r *ProductRouter) SetupProductRoutes(mux *http.ServeMux) {
	mux.Handle("/products", r.productHandler)
	mux.Handle("/product/", r.productHandler)
}
