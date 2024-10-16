package contextkeys

type contextKey string

const (
	RequestIDKey   contextKey = "requestID"
	HTTPRequestKey contextKey = "httpRequest"
	UserIDKey      contextKey = "userID"
)
