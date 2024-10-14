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
	graphqlClient piyographql.Client
}

func NewUserUsecase(client piyographql.Client) UserUsecase {
	return &userUsecase{
		graphqlClient: client,
	}
}

func (uc *userUsecase) Get(ctx context.Context, ID string) (*models.User, error) {
	// todo trace log

	n := lenID(ID)
	logger.GetLogger().Info("lenID", "n", n)
	return uc.graphqlClient.GetUser(ctx, ID)
}

func lenID(ID string) int {
	return len(ID)
}

func (uc *userUsecase) List(ctx context.Context, offset, limit *int) ([]models.User, error) {
	return uc.graphqlClient.ListUser(ctx, offset, limit)
}
