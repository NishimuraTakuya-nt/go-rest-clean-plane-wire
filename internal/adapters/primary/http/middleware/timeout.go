package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

func Timeout(duration time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), duration)
			defer cancel()
			log := logger.NewLogger()

			rw, ok := w.(*ResponseWriter)
			if !ok {
				log.WarnContext(ctx, "ResponseWriter is not of type *ResponseWriter")
				return
			}

			done := make(chan bool)
			go func() {
				next.ServeHTTP(rw, r.WithContext(ctx))
				done <- true
			}()

			select {
			case <-done:
				return
			case <-ctx.Done():
				log.ErrorContext(r.Context(), "Request timed out")
				rw.WriteError(apperrors.NewTimeoutError("Request timed out", ctx.Err()))
			}
		})
	}
}
