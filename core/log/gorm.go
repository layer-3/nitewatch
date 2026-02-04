package log

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// GormLogger implements gorm.io/gorm/logger.Interface
type GormLogger struct {
	LogLevel      gormlogger.LogLevel
	SlowThreshold time.Duration
}

// NewGormLogger creates a new GormLogger
func NewGormLogger(lvl LogLevel) gormlogger.Interface {
	var gormLevel gormlogger.LogLevel

	switch lvl {
	case LevelDebug:
		gormLevel = gormlogger.Info
	case LevelInfo, LevelWarn:
		gormLevel = gormlogger.Warn
	case LevelError:
		gormLevel = gormlogger.Error
	default:
		gormLevel = gormlogger.Warn
	}

	return &GormLogger{
		LogLevel:      gormLevel,
		SlowThreshold: 200 * time.Millisecond,
	}
}

// LogMode sets the log level
func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs info
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		Info().Msgf(msg, data...)
	}
}

// Warn logs warn
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		Warn().Msgf(msg, data...)
	}
}

// Error logs error
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		Error().Msgf(msg, data...)
	}
}

// Trace logs trace
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	var event *zerolog.Event

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		event = Error().Err(err)
	} else if elapsed > l.SlowThreshold && l.SlowThreshold != 0 {
		event = Warn().Dur("elapsed", elapsed).Str("slow_sql", "true")
	} else if l.LogLevel >= gormlogger.Info {
		event = Info().Dur("elapsed", elapsed)
	} else {
		// If level is not Info but we are tracing, and no error/slow, maybe we shouldn't log?
		// But Trace is called... usually controlled by LogMode.
		// If LogMode is Info, we land above.
		// If LogMode is Warn, we land here only if error or slow.
		// If LogMode is Error, we land here only if error.
		return
	}

	event.
		Str("sql", sql).
		Int64("rows", rows).
		Msg("GORM")
}
