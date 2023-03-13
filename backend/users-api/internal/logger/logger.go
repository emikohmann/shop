package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
	"users-api/internal/tracing"
)

type Logger struct {
    inner *logrus.Logger
}

// NewLogger returns a new instance of application Logger
func NewLogger(level logrus.Level) *Logger {
    formatter := new(logrus.TextFormatter)
    formatter.TimestampFormat = time.RFC3339
    formatter.FullTimestamp = true

    inner := logrus.New()
    logrus.SetFormatter(formatter)
    inner.SetLevel(level)

    return &Logger{
        inner: inner,
    }
}

func (logger *Logger) withExtra(ctx context.Context) *logrus.Entry {
    extra := make(map[string]interface{})
    requestID := tracing.GetRequestID(ctx)
    if requestID != "" {
        extra[tracing.RequestIDKey] = requestID
    }
    return logger.inner.WithFields(extra)
}

func (logger *Logger) Logf(ctx context.Context, level logrus.Level, format string, args ...interface{}) {
    logger.withExtra(ctx).Logf(level, format, args...)
}

func (logger *Logger) Tracef(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Tracef(format, args...)
}

func (logger *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Debugf(format, args...)
}

func (logger *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Infof(format, args...)
}

func (logger *Logger) Printf(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Printf(format, args...)
}

func (logger *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Warnf(format, args...)
}

func (logger *Logger) Warningf(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Warningf(format, args...)
}

func (logger *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Errorf(format, args...)
}

func (logger *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Fatalf(format, args...)
}

func (logger *Logger) Panicf(ctx context.Context, format string, args ...interface{}) {
    logger.withExtra(ctx).Panicf(format, args...)
}

func (logger *Logger) Log(ctx context.Context, level logrus.Level, args ...interface{}) {
    logger.withExtra(ctx).Log(level, args...)
}

func (logger *Logger) Trace(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Trace(args...)
}

func (logger *Logger) Debug(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Debug(args...)
}

func (logger *Logger) Info(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Info(args...)
}

func (logger *Logger) Print(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Print(args...)
}

func (logger *Logger) Warn(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Warn(args...)
}

func (logger *Logger) Warning(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Warning(args...)
}

func (logger *Logger) Error(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Error(args...)
}

func (logger *Logger) Fatal(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Fatal(args...)
}

func (logger *Logger) Panic(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Panic(args...)
}

func (logger *Logger) Logln(ctx context.Context, level logrus.Level, args ...interface{}) {
    logger.withExtra(ctx).Logln(level, args...)
}

func (logger *Logger) Traceln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Traceln(args...)
}

func (logger *Logger) Debugln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Debugln(args...)
}

func (logger *Logger) Infoln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Infoln(args...)
}

func (logger *Logger) Println(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Println(args...)
}

func (logger *Logger) Warnln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Warnln(args...)
}

func (logger *Logger) Warningln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Warningln(args...)
}

func (logger *Logger) Errorln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Errorln(args...)
}

func (logger *Logger) Fatalln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Fatalln(args...)
}

func (logger *Logger) Panicln(ctx context.Context, args ...interface{}) {
    logger.withExtra(ctx).Panicln(args...)
}
