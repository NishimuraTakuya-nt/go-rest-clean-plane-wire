package auth

import (
	"errors"
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(userID string, roles []string) (string, error)
	ValidateToken(tokenString string) (*models.Claims, error)
}

type tokenService struct {
	config *config.Config
}

func NewTokenService(cfg *config.Config) TokenService {
	return &tokenService{
		config: cfg,
	}
}

func (s *tokenService) GenerateToken(userID string, roles []string) (string, error) {
	claims := &models.Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecretKey))
}

func (s *tokenService) ValidateToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(_ *jwt.Token) (any, error) {
		return []byte(s.config.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
