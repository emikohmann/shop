package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewLogger(t *testing.T) {
	ctx := context.Background()

	debugLevel := NewLogger(logrus.DebugLevel)
	debugLevel.Debug(ctx, "Test log")
	debugLevel.Info(ctx, "Test log")
	debugLevel.Warn(ctx, "Test log")
	debugLevel.Error(ctx, "Test log")

	infoLevel := NewLogger(logrus.InfoLevel)
	infoLevel.Debug(ctx, "Test log")
	infoLevel.Info(ctx, "Test log")
	infoLevel.Warn(ctx, "Test log")
	infoLevel.Error(ctx, "Test log")

	warnLevel := NewLogger(logrus.WarnLevel)
	warnLevel.Debug(ctx, "Test log")
	warnLevel.Info(ctx, "Test log")
	warnLevel.Warn(ctx, "Test log")
	warnLevel.Error(ctx, "Test log")

	errorLevel := NewLogger(logrus.ErrorLevel)
	errorLevel.Debug(ctx, "Test log")
	errorLevel.Info(ctx, "Test log")
	errorLevel.Warn(ctx, "Test log")
	errorLevel.Error(ctx, "Test log")
}
