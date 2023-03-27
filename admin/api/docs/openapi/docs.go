// Code generated by swaggo/swag. DO NOT EDIT
package openapi

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.linkedin.com/in/emilianokohmann/",
        "contact": {
            "name": "Emiliano Kohmann",
            "url": "https://www.linkedin.com/in/emilianokohmann/",
            "email": "emikohmann@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/services": {
            "get": {
                "description": "Return the services information fetching information from the configuration.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Return the list of current services.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.ListServicesResponseHTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.APIErrorHTTP": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Some information not found"
                },
                "status": {
                    "type": "integer",
                    "example": 404
                }
            }
        },
        "http.ListServicesResponseHTTP": {
            "type": "object",
            "properties": {
                "services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.ServiceResponseHTTP"
                    }
                }
            }
        },
        "http.ServiceResponseHTTP": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9999",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Admin API",
	Description:      "This is an API that handles the admin commands.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
