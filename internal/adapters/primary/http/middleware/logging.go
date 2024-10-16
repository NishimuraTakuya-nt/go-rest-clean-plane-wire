package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

// Logging logs details of each HTTP request
func Logging() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log := logger.NewLogger()

			// リクエスト開始時のログ
			log.InfoContext(r.Context(), "Request started")

			// ResponseWriter のラッピングを一度だけ行う（後続のMiddlewareではこれを使い回す）
			rw := NewResponseWriter(w)

			// 次のハンドラを呼び出し
			next.ServeHTTP(rw, r)

			// リクエスト終了時のログ
			duration := time.Since(start)
			log.InfoContext(r.Context(),
				"Request completed",
				slog.Int("status", rw.StatusCode),
				slog.Int64("bytes", rw.Length),
				slog.String("duration", duration.String()),
			)
		})
	}
}
