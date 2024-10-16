//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/routes"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/routes/v1"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/secondary/piyographql"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/usecases"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/auth"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/config"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (http.Handler, error) {
	wire.Build(
		logger.Set,
		piyographql.Set,
		auth.Set,
		usecases.Set,
		handlers.Set,
		v1.Set,
		routes.Set,
	)
	return nil, nil
}
