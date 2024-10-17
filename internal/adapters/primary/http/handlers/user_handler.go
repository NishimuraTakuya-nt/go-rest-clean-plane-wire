package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/dto/request"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/adapters/primary/http/dto/response"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/usecases"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/pkg/validator"
)

type UserHandler struct {
	log         logger.Logger
	userUsecase usecases.UserUsecase
}

func NewUserHandler(log logger.Logger, userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		log:         log,
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

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		h.log.WarnContext(ctx, "Invalid user ID in request")
		writeError(w, apperrors.NewBadRequestError("Invalid user ID", nil))
		return
	}
	userID := parts[2]

	user, err := h.userUsecase.Get(ctx, userID)
	if err != nil {
		h.log.ErrorContext(ctx, "Failed to get user", "error", err)
		writeError(w, err)
		return
	}

	res := response.ToUserResponse(user)

	writeJSONResponse(ctx, w, res)
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

	// クエリパラメータの取得とバリデーション
	offset, limit, err := getPaginationParams(r)
	if err != nil {
		h.log.WarnContext(ctx, "Invalid pagination parameters", "error", err)
		writeError(w, apperrors.NewBadRequestError(err.Error(), err))
		return
	}

	// ユーザーリストの取得
	users, err := h.userUsecase.List(ctx, offset, limit)
	if err != nil {
		h.log.ErrorContext(ctx, "Failed to get user list", "error", err)
		writeError(w, err)
		return
	}

	res := response.ToListUserResponse(users, *offset, *limit)

	writeJSONResponse(ctx, w, res)
}

// Create godoc
// @Summary User create
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param request body request.UserRequest true "User information"
// @Security ApiKeyAuth
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /user [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// ユーザー作成処理
	ctx := r.Context()

	var req request.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.ErrorContext(ctx, "Failed to decode user request", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if validationErrors := validator.Validate(req); validationErrors != nil {
		writeError(w, validationErrors)
		return
	}

	res := response.UserResponse{
		ID:        "123",
		Name:      req.Name,
		Roles:     []string{"role:teamA:editor", "role:teamB:viewer"},
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	writeJSONResponse(ctx, w, res)
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
