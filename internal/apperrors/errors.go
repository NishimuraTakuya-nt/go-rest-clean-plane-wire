package apperrors

import (
	"net/http"
)

// AppError はアプリケーション固有のエラーを表します。
type AppError struct {
	Type       ErrorType
	RawError   error
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

// ErrorType はエラーの種類を表す列挙型です。
type ErrorType string

const (
	ErrorTypeBadRequest         ErrorType = "BAD_REQUEST"
	ErrorTypeUnauthorized       ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden          ErrorType = "FORBIDDEN"
	ErrorTypeNotFound           ErrorType = "NOT_FOUND"
	ErrorTypeConflict           ErrorType = "CONFLICT"
	ErrorTypeRateLimit          ErrorType = "RATE_LIMIT"
	ErrorTypeInternal           ErrorType = "INTERNAL_ERROR"
	ErrorTypeExternalService    ErrorType = "EXTERNAL_SERVICE_ERROR"
	ErrorTypeServiceUnavailable ErrorType = "SERVICE_UNAVAILABLE"
	ErrorTypeTimeout            ErrorType = "TIMEOUT"
)

// NewAppError creates a new AppError
func NewAppError(errType ErrorType, rawErr error, statusCode int, message string) *AppError {
	return &AppError{
		Type:       errType,
		RawError:   rawErr,
		StatusCode: statusCode,
		Message:    message,
	}
}

// NewBadRequestError 400 Bad Request
func NewBadRequestError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeBadRequest, rawErr, http.StatusBadRequest, message)
}

// NewUnauthorizedError 401 Unauthorized
func NewUnauthorizedError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeUnauthorized, rawErr, http.StatusUnauthorized, message)
}

// NewForbiddenError 403 Forbidden
func NewForbiddenError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeForbidden, rawErr, http.StatusForbidden, message)
}

// NewNotFoundError 404 Not Found
func NewNotFoundError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeNotFound, rawErr, http.StatusNotFound, message)
}

// NewConflictError 409 Conflict
func NewConflictError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeConflict, rawErr, http.StatusConflict, message)
}

// NewRateLimitError 429 Too Many Requests
func NewRateLimitError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeRateLimit, rawErr, http.StatusTooManyRequests, message)
}

// NewInternalError 500 Internal Server Error
func NewInternalError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeInternal, rawErr, http.StatusInternalServerError, message)
}

// NewExternalServiceError 502 Bad Gateway
func NewExternalServiceError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeExternalService, rawErr, http.StatusBadGateway, message)
}

// NewServiceUnavailableError 503 Service Unavailable
func NewServiceUnavailableError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeServiceUnavailable, rawErr, http.StatusServiceUnavailable, message)
}

// NewTimeoutError 504 Gateway Timeout
func NewTimeoutError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeTimeout, rawErr, http.StatusGatewayTimeout, message)
}
