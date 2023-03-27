package admin

import (
	"api/internal/logger"
)

type service struct {
	logger *logger.Logger
}

func NewService(logger *logger.Logger) *service {
	return &service{
		logger: logger,
	}
}
