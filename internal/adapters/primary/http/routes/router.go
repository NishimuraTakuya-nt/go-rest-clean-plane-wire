package routes

import (
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/usecases"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/handlers"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/middleware"
	v1 "github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/routes/v1"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/config"
)

func NewRouter(
	cfg *config.Config,
	authUsecase usecases.AuthUsecase,
	healthcheckRouter *v1.HealthcheckRouter,
	authRouter *v1.AuthRouter,
	userRouter *v1.UserRouter,
	productRouter *v1.ProductRouter,
	orderRouter *v1.OrderRouter,
) http.Handler {
	mux := http.NewServeMux()

	// ルートハンドラの登録
	mux.HandleFunc("/", handlers.HomeHandler)
	// Swagger 2.0
	mux.Handle("/swagger/2.0/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	// OAS 3.0
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger/openapi3.json"),
	))
	mux.Handle("/docs/swagger/", http.StripPrefix("/docs/swagger/", http.FileServer(http.Dir("./docs/swagger"))))

	// API v1 ルート
	apiV1 := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	healthcheckRouter.SetupHealthcheckRoutes(apiV1)
	authRouter.SetupAuthRoutes(apiV1)
	userRouter.SetupUserRoutes(apiV1)
	productRouter.SetupProductRoutes(apiV1)
	orderRouter.SetupOrderRoutes(apiV1)

	// CORSの設定
	corsConfig := middleware.DefaultCORSConfig()
	corsConfig.AllowOrigins = cfg.AllowedOrigins

	// ミドルウェアの適用
	handler := middleware.Chain(
		mux,
		middleware.Context(),
		middleware.CORS(corsConfig),
		middleware.Logging(),
		middleware.ErrorHandler(),
		middleware.Timeout(cfg.RequestTimeout),
		middleware.Authenticate(authUsecase), // fix 面倒なので一旦OFF
	)
	return handler
}
