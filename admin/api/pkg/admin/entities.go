package admin

import "api/pkg/admin/services"

type Service struct {
	Name    string
	Status  services.Status
	Health  services.Health
	Version string
}
