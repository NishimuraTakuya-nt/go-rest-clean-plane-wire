package usecases

import (
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/auth"
)

type AuthUsecase interface {
	Login(userID string, roles []string) (string, error)
	Authenticate(tokenString string) (*models.User, error)
}

type authUsecase struct {
	tokenService auth.TokenService
}

func NewAuthUsecase(tokenService auth.TokenService) AuthUsecase {
	return &authUsecase{
		tokenService: tokenService,
	}
}

func (uc *authUsecase) Login(userID string, roles []string) (string, error) {
	return uc.tokenService.GenerateToken(userID, roles)
}

func (uc *authUsecase) Authenticate(tokenString string) (*models.User, error) {
	claims, err := uc.tokenService.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:    claims.UserID,
		Roles: claims.Roles,
	}, nil
}
