package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}
