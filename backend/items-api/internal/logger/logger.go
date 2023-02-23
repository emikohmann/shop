package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

// NewLogger returns a new instance of application logger
func NewLogger(level logrus.Level) *logrus.Logger {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.RFC3339
	formatter.FullTimestamp = true

	logger := logrus.New()
	logrus.SetFormatter(formatter)
	logger.SetLevel(level)

	return logger
}
