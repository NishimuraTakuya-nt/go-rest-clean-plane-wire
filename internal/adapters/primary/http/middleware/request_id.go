package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

var RequestIDKey = struct{}{}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			// nolint:staticcheck
			ctx := context.WithValue(r.Context(), RequestIDKey, requestID)

			w.Header().Set("X-Request-ID", requestID)

			// 次のハンドラを呼び出し
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}
