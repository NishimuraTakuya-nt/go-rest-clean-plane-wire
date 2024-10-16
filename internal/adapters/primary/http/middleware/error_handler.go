package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/dto/response"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

func ErrorHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logger.NewLogger()

			rw, ok := w.(*ResponseWriter)
			if !ok {
				log.Warn("ResponseWriter is not of type *ResponseWriter")
				return
			}

			defer func() {
				if re := recover(); re != nil {
					var panicErr error
					switch err := re.(type) {
					case string:
						panicErr = errors.New(err)
					case error:
						panicErr = err
					default:
						panicErr = fmt.Errorf("unknown panic: %v", err)
					}

					// スタックトレースを取得
					stack := debug.Stack()

					// 詳細なログを記録
					log.ErrorContext(r.Context(),
						"Panic occurred",
						"error", panicErr,
						"stack", string(stack),
					)

					// クライアントへのレスポンス用エラー
					clientErr := apperrors.NewInternalError("Unexpected error occurred", panicErr)

					// エラーハンドリング
					handleError(r.Context(), rw, clientErr)
				}
			}()

			next.ServeHTTP(rw, r)

			if rw.Err != nil {
				handleError(r.Context(), rw, rw.Err)
			}
		})
	}
}

func handleError(ctx context.Context, rw *ResponseWriter, err error) {
	var res response.ErrorResponse
	var statusCode int
	requestID := GetRequestID(ctx)

	switch e := err.(type) {
	case *apperrors.ValidationErrors:
		statusCode = http.StatusBadRequest
		details := make([]map[string]string, 0, len(*e))
		for _, fe := range *e {
			details = append(details, map[string]string{
				"field":   fe.Field,
				"message": fe.Message,
			})
		}
		res = response.ErrorResponse{
			StatusCode: statusCode,
			Type:       string(apperrors.ErrorTypeBadRequest),
			RequestID:  requestID,
			Message:    "Validation error",
			Details:    details,
		}

	case *apperrors.AppError:
		statusCode = e.StatusCode
		res = response.ErrorResponse{
			StatusCode: statusCode,
			Type:       string(e.Type),
			RequestID:  requestID,
			Message:    e.Message,
		}

	default:
		statusCode = http.StatusInternalServerError
		res = response.ErrorResponse{
			StatusCode: statusCode,
			Type:       string(apperrors.ErrorTypeInternal),
			RequestID:  requestID,
			Message:    "Internal server error",
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		logger.NewLogger().ErrorContext(ctx, "Failed to encode error response", "error", err)
	}
}
