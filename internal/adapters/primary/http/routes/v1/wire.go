package v1

import "github.com/google/wire"

var Set = wire.NewSet(
	NewHealthcheckRouter,
	NewAuthRouter,
	NewUserRouter,
	NewProductRouter,
	NewOrderRouter,
)
