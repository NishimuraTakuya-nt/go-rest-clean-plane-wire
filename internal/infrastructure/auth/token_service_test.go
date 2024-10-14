package auth

import (
	"testing"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/mocks/mockauth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTokenService := mockauth.NewMockTokenService(ctrl)

	testCases := []struct {
		name     string
		userID   string
		roles    []string
		mockFunc func()
		want     string
		wantErr  bool
	}{
		{
			name:   "正常系: トークンが正しく生成される",
			userID: "user123",
			roles:  []string{"user", "admin"},
			mockFunc: func() {
				mockTokenService.EXPECT().
					GenerateToken("user123", []string{"user", "admin"}).
					Return("mocked.jwt.token", nil)
			},
			want:    "mocked.jwt.token",
			wantErr: false,
		},
		{
			name:   "異常系: エラーが返される",
			userID: "user456",
			roles:  []string{"user"},
			mockFunc: func() {
				mockTokenService.EXPECT().
					GenerateToken("user456", []string{"user"}).
					Return("", jwt.ErrSignatureInvalid)
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()

			got, err := mockTokenService.GenerateToken(tc.userID, tc.roles)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
