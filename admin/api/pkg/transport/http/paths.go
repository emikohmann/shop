package http

import "fmt"

const (
	paramServiceID = "serviceID"
)

var (
	PathDocs         = "/docs/*any"
	PathListServices = "/services"
	PathGetService   = fmt.Sprintf("/services/:%s", paramServiceID)
)
