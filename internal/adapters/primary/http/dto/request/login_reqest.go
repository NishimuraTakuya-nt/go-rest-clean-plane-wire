package request

// LoginRequest
// @Description LoginRequest is a struct that represents the request of login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
