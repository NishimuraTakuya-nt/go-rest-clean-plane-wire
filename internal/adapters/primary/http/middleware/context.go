package middleware

import (
	"context"
	"net/http"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/common/contextkeys"
	"github.com/google/uuid"
)

func Context() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// リクエストIDをコンテキストに追加
			requestID := uuid.New().String()
			// nolint:staticcheck
			ctx := context.WithValue(r.Context(), contextkeys.RequestIDKey, requestID)
			w.Header().Set("X-Request-ID", requestID)

			// HTTPリクエスト情報をコンテキストに追加
			ctx = context.WithValue(ctx, contextkeys.HTTPRequestKey, r)

			// 次のハンドラを呼び出し
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(contextkeys.RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}
