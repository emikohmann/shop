package logger

import (
	"context"
	"github.com/emikohmann/shop/backend/items-api/internal/tracing"
	"github.com/sirupsen/logrus"
	"time"
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

func (logger *Logger) extraArgs(ctx context.Context) []interface{} {
	extra := make([]interface{}, 0)
	requestID := tracing.GetRequestID(ctx)
	if requestID != "" {
		extra = append(extra, tracing.RequestIDKey, requestID)
	}
	return extra
}

func (logger *Logger) Logf(ctx context.Context, level logrus.Level, format string, args ...interface{}) {
	logger.inner.Logf(level, format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Tracef(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Tracef(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Debugf(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Infof(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Printf(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Printf(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Warnf(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Warningf(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Warningf(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Errorf(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Fatalf(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Panicf(ctx context.Context, format string, args ...interface{}) {
	logger.inner.Panicf(format, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Log(ctx context.Context, level logrus.Level, args ...interface{}) {
	logger.inner.Log(level, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) LogFn(ctx context.Context, level logrus.Level, fn logrus.LogFunction) {
	logger.inner.LogFn(level, fn)
}

func (logger *Logger) Trace(ctx context.Context, args ...interface{}) {
	logger.inner.Trace(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Debug(ctx context.Context, args ...interface{}) {
	logger.inner.Debug(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Info(ctx context.Context, args ...interface{}) {
	logger.inner.Info(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Print(ctx context.Context, args ...interface{}) {
	logger.inner.Print(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Warn(ctx context.Context, args ...interface{}) {
	logger.inner.Warn(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Warning(ctx context.Context, args ...interface{}) {
	logger.inner.Warning(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Error(ctx context.Context, args ...interface{}) {
	logger.inner.Error(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Fatal(ctx context.Context, args ...interface{}) {
	logger.inner.Fatal(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Panic(ctx context.Context, args ...interface{}) {
	logger.inner.Panic(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) TraceFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.TraceFn(fn)
}

func (logger *Logger) DebugFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.DebugFn(fn)
}

func (logger *Logger) InfoFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.InfoFn(fn)
}

func (logger *Logger) PrintFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.PrintFn(fn)
}

func (logger *Logger) WarnFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.WarnFn(fn)
}

func (logger *Logger) WarningFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.WarningFn(fn)
}

func (logger *Logger) ErrorFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.ErrorFn(fn)
}

func (logger *Logger) FatalFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.FatalFn(fn)
}

func (logger *Logger) PanicFn(ctx context.Context, fn logrus.LogFunction) {
	logger.inner.PanicFn(fn)
}

func (logger *Logger) Logln(ctx context.Context, level logrus.Level, args ...interface{}) {
	logger.inner.Logln(level, append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Traceln(ctx context.Context, args ...interface{}) {
	logger.inner.Traceln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Debugln(ctx context.Context, args ...interface{}) {
	logger.inner.Debugln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Infoln(ctx context.Context, args ...interface{}) {
	logger.inner.Infoln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Println(ctx context.Context, args ...interface{}) {
	logger.inner.Println(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Warnln(ctx context.Context, args ...interface{}) {
	logger.inner.Warnln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Warningln(ctx context.Context, args ...interface{}) {
	logger.inner.Warningln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Errorln(ctx context.Context, args ...interface{}) {
	logger.inner.Errorln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Fatalln(ctx context.Context, args ...interface{}) {
	logger.inner.Fatalln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Panicln(ctx context.Context, args ...interface{}) {
	logger.inner.Panicln(append(args, logger.extraArgs(ctx))...)
}

func (logger *Logger) Exit(ctx context.Context, code int) {
	logger.inner.Exit(code)
}
