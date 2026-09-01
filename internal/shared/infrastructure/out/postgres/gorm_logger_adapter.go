package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/gerarc/tireg/internal/shared/application/utils/logger"
)

type GormLoggerAdapter struct {
	logger logger.Logger
}

func NewGormLoggerAdapter(appLogger logger.Logger) gormlogger.Interface {
	return &GormLoggerAdapter{logger: appLogger}
}

func (gormLoggerAdapter *GormLoggerAdapter) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return gormLoggerAdapter
}

func (gormLoggerAdapter *GormLoggerAdapter) Info(ctx context.Context, msg string, data ...any) {
	gormLoggerAdapter.logger.Info(fmt.Sprintf(msg, data...))
}

func (gormLoggerAdapter *GormLoggerAdapter) Warn(ctx context.Context, msg string, data ...any) {
	gormLoggerAdapter.logger.Warn(fmt.Sprintf(msg, data...))
}

func (gormLoggerAdapter *GormLoggerAdapter) Error(ctx context.Context, msg string, data ...any) {
	gormLoggerAdapter.logger.Warn(fmt.Sprintf(msg, data...))
}

func (gormLoggerAdapter *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rowsAffected := fc()
	durationMS := time.Since(begin).Milliseconds()

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		gormLoggerAdapter.logger.Error("gorm query failed", err, "sql", sql, "rows", rowsAffected, "duration_ms", durationMS)
		return
	}

	gormLoggerAdapter.logger.Debug("gorm query", "sql", sql, "rows", rowsAffected, "duration_ms", durationMS)
}
