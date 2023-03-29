package http

import (
	"api/pkg/admin"
	"api/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// HTTPToGetServiceRequest turns the HTTP request into a ListItemsRequest
func HTTPToGetServiceRequest(ctx *gin.Context) (admin.GetServiceRequest, error) {
	serviceID := ctx.Param(paramServiceID)
	if util.IsEmpty(serviceID) {
		return admin.GetServiceRequest{}, fmt.Errorf("%s cannot be empty", paramServiceID)
	}
	return admin.GetServiceRequest{
		ID: serviceID,
	}, nil
}
