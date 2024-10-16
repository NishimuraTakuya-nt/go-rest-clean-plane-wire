// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/routes"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/routes/v1"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/secondary/piyographql"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/usecases"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/auth"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/config"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
	"net/http"
)

import (
	_ "github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/docs/swagger"
)

// Injectors from wire.go:

func InitializeAPI(cfg *config.Config) (http.Handler, error) {
	tokenService := auth.NewTokenService(cfg)
	authUsecase := usecases.NewAuthUsecase(tokenService)
	loggerLogger := logger.NewLogger()
	healthcheckHandler := handlers.NewHealthcheckHandler(loggerLogger)
	healthcheckRouter := v1.NewHealthcheckRouter(healthcheckHandler)
	authHandler := handlers.NewAuthHandler(loggerLogger, authUsecase)
	authRouter := v1.NewAuthRouter(authHandler)
	client := piyographql.NewClient(loggerLogger)
	userUsecase := usecases.NewUserUsecase(loggerLogger, client)
	userHandler := handlers.NewUserHandler(loggerLogger, userUsecase)
	userRouter := v1.NewUserRouter(userHandler)
	productHandler := handlers.NewProductHandler()
	productRouter := v1.NewProductRouter(productHandler)
	orderHandler := handlers.NewOrderHandler()
	orderRouter := v1.NewOrderRouter(orderHandler)
	handler := routes.NewRouter(cfg, authUsecase, healthcheckRouter, authRouter, userRouter, productRouter, orderRouter)
	return handler, nil
}
