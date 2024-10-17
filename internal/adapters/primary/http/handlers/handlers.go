package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/apperrors"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/logger"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the API"})
	if err != nil {
		logger.NewLogger().Warn("Failed to encode response", "error", err)
		return
	}
}

func writeError(w http.ResponseWriter, err error) {
	if ew, ok := w.(apperrors.ErrorWriter); ok {
		ew.WriteError(err)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func writeJSONResponse(ctx context.Context, w http.ResponseWriter, data any) {
	select {
	case <-ctx.Done():
		// コンテキストがキャンセルされている場合
		return
	default:
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.NewLogger().ErrorContext(ctx, "Failed to encode response", "error", err)
		writeError(w, apperrors.NewInternalError("Failed to encode response", err))
	}
}

// todo 実装の場所変える。処理がこれで良いかも使う時にブラッシュアップする
// getPaginationParams はクエリパラメータからオフセットとリミットを取得し、バリデーションを行います
func getPaginationParams(r *http.Request) (*int, *int, error) {
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	var offset, limit int

	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			return nil, nil, errors.New("invalid offset parameter")
		}
	}

	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 || limit > 100 {
			return nil, nil, errors.New("invalid limit parameter")
		}
	}

	return &offset, &limit, nil
}
