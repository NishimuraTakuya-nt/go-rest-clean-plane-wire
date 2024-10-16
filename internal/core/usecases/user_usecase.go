package usecases

import (
	"context"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/secondary/piyographql"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

type UserUsecase interface {
	Get(ctx context.Context, ID string) (*models.User, error)
	List(ctx context.Context, offset, limit *int) ([]models.User, error)
}

type userUsecase struct {
	log           logger.Logger
	graphqlClient piyographql.Client
}

func NewUserUsecase(log logger.Logger, client piyographql.Client) UserUsecase {
	return &userUsecase{
		log:           log,
		graphqlClient: client,
	}
}

func (uc *userUsecase) Get(ctx context.Context, ID string) (*models.User, error) {
	// todo trace log
	return uc.graphqlClient.GetUser(ctx, ID)
}

func (uc *userUsecase) List(ctx context.Context, offset, limit *int) ([]models.User, error) {
	return uc.graphqlClient.ListUser(ctx, offset, limit)
}
