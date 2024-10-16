package piyographql

import (
	"context"
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

type Client interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
	ListUser(ctx context.Context, offset, limit *int) ([]models.User, error)
}

type client struct {
	log logger.Logger
	// クライアントの設定など
}

func NewClient(log logger.Logger) Client {
	return &client{
		log: log,
	}
}

func (c *client) GetUser(_ context.Context, ID string) (*models.User, error) {
	// ここでは簡易的に固定のユーザーを返していますが、
	// 実際には取得する処理を実装します

	return &models.User{
		ID:        ID,
		Name:      "example_user",
		Roles:     []string{"teamA:editor", "teamB:viewer"},
		Email:     "user@example.com",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now(),
	}, nil
}

func (c *client) ListUser(ctx context.Context, offset, limit *int) ([]models.User, error) {
	c.log.InfoContext(ctx, "client ListUser", "offset", offset, "limit", limit)
	return []models.User{
		{
			ID:    "1",
			Name:  "example_user1",
			Roles: []string{"teamA:editor", "teamB:viewer"},
		},
		{
			ID:    "2",
			Name:  "example_user2",
			Roles: []string{"teamA:viewer", "teamB:editor"},
		},
	}, nil
}
