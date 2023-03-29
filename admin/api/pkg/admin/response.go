package admin

type ListServicesResponse struct {
	Services []Service
}

type GetServiceResponse struct {
	Service              Service
	DockerAdditionalInfo DockerAdditionalInfo
}
