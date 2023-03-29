package http

import (
	"api/internal/apierrors"
	"api/pkg/admin"
	"api/pkg/util"
	"time"
)

type APIErrorHTTP struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"Some information not found"`
}

type ListServicesResponseHTTP struct {
	Services []ServiceResponseHTTP `json:"services"`
}

type ServiceResponseHTTP struct {
	Name         string `json:"name,omitempty" example:"Items API"`
	Status       string `json:"status,omitempty" example:"RUNNING"`
	StatusDetail string `json:"status_detail,omitempty" example:"Up 2 hours"`
	Health       string `json:"health,omitempty" example:"HEALTHY"`
	Version      string `json:"version,omitempty" example:"0.0.1"`
	Port         int    `json:"port,omitempty" example:"8080"`
	Network      string `json:"network,omitempty" example:"shop_default"`
	CreationDate string `json:"creation_date,omitempty" example:"2023-03-29T03:36:19Z"`
}

type GetServiceResponseHTTP struct {
	Name           string                     `json:"name,omitempty" example:"Items API"`
	Status         string                     `json:"status,omitempty" example:"RUNNING"`
	StatusDetail   string                     `json:"status_detail,omitempty" example:"Up 2 hours"`
	Health         string                     `json:"health,omitempty" example:"HEALTHY"`
	Version        string                     `json:"version,omitempty" example:"0.0.1"`
	Port           int                        `json:"port,omitempty" example:"8080"`
	Network        string                     `json:"network,omitempty" example:"shop_default"`
	CreationDate   string                     `json:"creation_date,omitempty" example:"2023-03-29T03:36:19Z"`
	AdditionalInfo admin.DockerAdditionalInfo `json:"additional_info,omitempty"`
}

const (
	httpDateTimeLayout = time.RFC3339
)

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
			Name:         service.Name,
			Status:       service.Status.String(),
			StatusDetail: service.StatusDetail,
			Health:       service.Health.String(),
			Version:      service.Version,
			Port:         service.Port,
			Network:      service.Network,
			CreationDate: dateToString(service.CreationDate),
		})
	}
	return ListServicesResponseHTTP{
		Services: servicesHTTP,
	}
}

// GetServiceResponseToHTTP prepares the GetServiceResponse to be presented as HTTP
func GetServiceResponseToHTTP(response admin.GetServiceResponse) GetServiceResponseHTTP {
	return GetServiceResponseHTTP{
		Name:           response.Service.Name,
		Status:         response.Service.Status.String(),
		StatusDetail:   response.Service.StatusDetail,
		Health:         response.Service.Health.String(),
		Version:        response.Service.Version,
		Port:           response.Service.Port,
		Network:        response.Service.Network,
		CreationDate:   dateToString(response.Service.CreationDate),
		AdditionalInfo: response.DockerAdditionalInfo,
	}
}

// dateToString converts a date into a string only if the time is not empty
func dateToString(date time.Time) string {
	var result string
	if !util.IsEmpty(date) {
		result = date.Format(httpDateTimeLayout)
	}
	return result
}
