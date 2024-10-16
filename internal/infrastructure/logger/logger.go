package logger

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/core/common/contextkeys"
)

// Logger インターフェース
type Logger interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	With(args ...any) Logger
}

// customLogger 構造体
type customLogger struct {
	logger *slog.Logger
}

// NewLogger カスタムロガーを作成
func NewLogger() Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     getLevelFromEnv(),
		ReplaceAttr: func(_ []string, att slog.Attr) slog.Attr {
			if att.Key == slog.SourceKey {
				const skip = 7
				_, file, line, ok := runtime.Caller(skip)
				if !ok {
					return att
				}
				v := fmt.Sprintf("%s:%d", filepath.Base(file), line)
				att.Value = slog.StringValue(v)
			}
			return att
		},
	})

	return &customLogger{
		logger: slog.New(handler),
	}
}

// addStandardFields 標準フィールドを追加
func (l *customLogger) addStandardFields(ctx context.Context, args []any) []any {
	if ctx != nil {
		if r, ok := ctx.Value(contextkeys.RequestIDKey).(string); ok {
			args = append(args, "request_id", r)
		}
		if u, ok := ctx.Value(contextkeys.UserIDKey).(string); ok {
			args = append(args, "user_id", u)
		}
		if r, ok := ctx.Value(contextkeys.HTTPRequestKey).(*http.Request); ok {
			args = append(args,
				"endpoint", r.URL.Path,
				"method", r.Method,
				"remote_ip", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)
		}
	}
	return args
}

func (l *customLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *customLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	args = l.addStandardFields(ctx, args)
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *customLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *customLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	args = l.addStandardFields(ctx, args)
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *customLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *customLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	args = l.addStandardFields(ctx, args)
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *customLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *customLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	args = l.addStandardFields(ctx, args)
	l.logger.ErrorContext(ctx, msg, args...)
}

// With 追加のフィールドを持つ新しいロガーを作成
func (l *customLogger) With(args ...any) Logger {
	return &customLogger{
		logger: l.logger.With(args...),
	}
}

// getLevelFromEnv
func getLevelFromEnv() slog.Level {
	levelStr := os.Getenv("LOG_LEVEL")
	switch levelStr {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
