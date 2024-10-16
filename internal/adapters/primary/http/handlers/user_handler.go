package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/dto/response"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/middleware"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/usecases"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/user/"):
		h.Get(w, r)

	case r.Method == http.MethodGet && r.URL.Path == "/users":
		h.List(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/user":
		h.Create(w, r)
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/user"):
		h.Update(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/user"):
		h.Delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

// Get godoc
// @Summary Get a user by ID
// @Description Get details of a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Security ApiKeyAuth
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /user/{id} [get]
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger()
	requestID := middleware.GetRequestID(ctx)

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		log.Warn("Invalid user ID in request", "path", r.URL.Path, "request_id", requestID)
		writeError(w, apperrors.NewBadRequestError("Invalid user ID", nil))
		return
	}
	userID := parts[2]

	user, err := h.userUsecase.Get(ctx, userID)
	if err != nil {
		log.Error("Failed to get user", "error", err, "user_id", userID, "request_id", requestID)
		writeError(w, err)
		return
	}

	res := response.ToUserResponse(user)

	writeJSONResponse(w, res, requestID)
}

// List godoc
// @Summary List users
// @Description Get a list of users with pagination
// @Tags user
// @Accept  json
// @Produce  json
// @Param offset query int false "Offset for pagination" default(0) minimum(0)
// @Param limit query int false "Limit for pagination" default(10) maximum(100)
// @Success 200 {object} response.ListUserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users [get]
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.GetLogger()
	requestID, _ := ctx.Value(middleware.RequestIDKey).(string)

	// クエリパラメータの取得とバリデーション
	offset, limit, err := getPaginationParams(r)
	if err != nil {
		log.Warn("Invalid pagination parameters", "error", err, "request_id", requestID)
		writeError(w, apperrors.NewBadRequestError(err.Error(), err))
		return
	}

	// ユーザーリストの取得
	users, err := h.userUsecase.List(ctx, offset, limit)
	if err != nil {
		log.Error("Failed to get user list", "error", err, "request_id", requestID)
		writeError(w, err)
		return
	}

	res := response.ToListUserResponse(users, *offset, *limit)

	writeJSONResponse(w, res, requestID)
}

func (h *UserHandler) Create(w http.ResponseWriter, _ *http.Request) {
	// ユーザー作成処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Create user"})
}

func (h *UserHandler) Update(w http.ResponseWriter, _ *http.Request) {
	// ユーザー更新処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Update user"})
}

func (h *UserHandler) Delete(w http.ResponseWriter, _ *http.Request) {
	// ユーザー削除処理
	// nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete user"})
}
