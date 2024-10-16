package usecases

import (
	"context"
	"testing"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
	"github.com/stretchr/testify/assert"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/mocks/mockpiyographql"
	"github.com/golang/mock/gomock"
)

func TestUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockpiyographql.NewMockClient(ctrl)
	target := NewUserUsecase(logger.NewLogger(), mockClient)

	t.Run("get user", func(t *testing.T) {
		ID := "123"
		// モックの振る舞いを設定
		mockClient.EXPECT().GetUser(context.Background(), ID).
			Return(&models.User{ID: "123", Name: "Test User"}, nil)

		// テストケースを実行
		user, err := target.Get(context.Background(), ID)

		assert.NoError(t, err)
		assert.Equal(t, "123", user.ID)
	})

	t.Run("get user 2", func(t *testing.T) {
		ID := "aaa"
		// モックの振る舞いを設定
		mockClient.EXPECT().GetUser(context.Background(), ID).
			Return(&models.User{ID: "aaa", Name: "Test User"}, nil)

		// テストケースを実行
		user, err := target.Get(context.Background(), ID)

		assert.NoError(t, err)
		assert.Equal(t, "aaa", user.ID)
	})
}
