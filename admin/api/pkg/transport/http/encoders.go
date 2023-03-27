package http

import (
	"api/internal/apierrors"
	"api/pkg/admin"
)

type APIErrorHTTP struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"Some information not found"`
}

type ListServicesResponseHTTP struct {
	Services []ServiceResponseHTTP `json:"services"`
}

type ServiceResponseHTTP struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

// APIErrorToHTTP prepares the APIError to be presented as HTTP
func APIErrorToHTTP(apiError apierrors.APIError) APIErrorHTTP {
	return APIErrorHTTP{
		Status:  apiError.Status(),
		Message: apiError.Message(),
	}
}

// ListServicesResponseToHTTP prepares the ListServicesResponse to be presented as HTTP
func ListServicesResponseToHTTP(response admin.ListServicesResponse) ListServicesResponseHTTP {
	servicesHTTP := make([]ServiceResponseHTTP, 0)
	for _, service := range response.Services {
		servicesHTTP = append(servicesHTTP, ServiceResponseHTTP{
			Name:    service.Name,
			Status:  service.Status,
			Version: service.Version,
		})
	}
	return ListServicesResponseHTTP{
		Services: servicesHTTP,
	}
}
