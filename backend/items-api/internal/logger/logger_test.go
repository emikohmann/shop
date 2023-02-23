package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewLogger(t *testing.T) {
	debugLevel := NewLogger(logrus.DebugLevel)
	debugLevel.WithField("level", "debug").Debug("Test log")
	debugLevel.WithField("level", "debug").Info("Test log")
	debugLevel.WithField("level", "debug").Warn("Test log")
	debugLevel.WithField("level", "debug").Error("Test log")

	infoLevel := NewLogger(logrus.InfoLevel)
	infoLevel.WithField("level", "info").Debug("Test log")
	infoLevel.WithField("level", "info").Info("Test log")
	infoLevel.WithField("level", "info").Warn("Test log")
	infoLevel.WithField("level", "info").Error("Test log")

	warnLevel := NewLogger(logrus.WarnLevel)
	warnLevel.WithField("level", "warn").Debug("Test log")
	warnLevel.WithField("level", "warn").Info("Test log")
	warnLevel.WithField("level", "warn").Warn("Test log")
	warnLevel.WithField("level", "warn").Error("Test log")

	errorLevel := NewLogger(logrus.ErrorLevel)
	errorLevel.WithField("level", "error").Debug("Test log")
	errorLevel.WithField("level", "error").Info("Test log")
	errorLevel.WithField("level", "error").Warn("Test log")
	errorLevel.WithField("level", "error").Error("Test log")
}
