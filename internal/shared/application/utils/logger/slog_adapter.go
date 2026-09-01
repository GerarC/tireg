package logger

import (
	"log/slog"
	"os"

	"github.com/gerarc/tireg/internal/shared/application/utils/config"
	"github.com/gerarc/tireg/internal/shared/application/utils/logger/util/constant"
)

type SlogAdapter struct {
	logger *slog.Logger
}

func NewSlogAdapter(appConfig *config.Config) Logger {
	var logLevel slog.Level
	switch appConfig.LogLevel {
	case constant.LevelDebug:
		logLevel = slog.LevelDebug
	case constant.LevelWarn:
		logLevel = slog.LevelWarn
	case constant.LevelError:
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	handler := NewCustomTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05.000"))
			}

			return a
		},
	})

	return &SlogAdapter{logger: slog.New(handler)}
}

func (s *SlogAdapter) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *SlogAdapter) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s *SlogAdapter) Error(msg string, err error, args ...any) {
	allArgs := append(args, slog.String("error", err.Error()))
	s.logger.Error(msg, allArgs...)
}

func (s *SlogAdapter) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *SlogAdapter) With(args ...any) Logger {
	return &SlogAdapter{logger: s.logger.With(args...)}
}
