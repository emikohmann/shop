package admin

import (
	"api/pkg/admin/services"
	"time"
)

type Service struct {
	Name         string
	Status       services.Status
	StatusDetail string
	Health       services.Health
	Version      string
	Port         int
	Network      string
	CreationDate time.Time
}

type DockerAdditionalInfo map[string]interface{}
