package middleware

import (
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
			requestID := GetRequestID(r.Context())
			log := logger.GetLogger()

			rw, ok := w.(*ResponseWriter)
			if !ok {
				log.Warn("ResponseWriter is not of type *ResponseWriter")
				return
			}

			defer func() {
				if r := recover(); r != nil {
					var panicErr error
					switch err := r.(type) {
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
					log.Error("Panic occurred",
						"error", panicErr,
						"stack", string(stack),
						"request_id", requestID,
					)

					// クライアントへのレスポンス用エラー
					clientErr := apperrors.NewInternalError("Unexpected error occurred", panicErr)

					// エラーハンドリング
					handleError(rw, clientErr, requestID)
				}
			}()

			next.ServeHTTP(rw, r)

			if rw.Err != nil {
				handleError(rw, rw.Err, requestID)
			}
		})
	}
}

func handleError(rw *ResponseWriter, err error, requestID string) {
	var res response.ErrorResponse
	var statusCode int

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
		logger.GetLogger().Error("Failed to encode error response", "error", err, "request_id", requestID)
	}
}
