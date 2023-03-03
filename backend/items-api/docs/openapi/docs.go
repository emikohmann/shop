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
        "/items": {
            "post": {
                "description": "Store the item information against the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Store the item information.",
                "parameters": [
                    {
                        "description": "Item to save",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.SaveItemRequestHTTP"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/http.SaveItemResponseHTTP"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
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
        },
        "/items/{itemID}": {
            "get": {
                "description": "Return the item information fetching information from the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Return the item information.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the item to get",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.GetItemResponseHTTP"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates the item information against the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Updates the item information.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the item to get",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Item fields to update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UpdateItemRequestHTTP"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.UpdateItemResponseHTTP"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the item information against the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Delete the item information.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the item to delete",
                        "name": "itemID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.DeleteItemResponseHTTP"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.APIErrorHTTP"
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
        "http.DeleteItemResponseHTTP": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "http.GetItemResponseHTTP": {
            "type": "object",
            "properties": {
                "date_created": {
                    "type": "string",
                    "example": "2023-02-23T21:46:28.366Z"
                },
                "description": {
                    "type": "string",
                    "example": "The iPhone 13 display has rounded corners"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://www.macstation.com.ar/img/productos/2599-2.jpg"
                    ]
                },
                "is_active": {
                    "type": "boolean",
                    "example": true
                },
                "last_updated": {
                    "type": "string",
                    "example": "2023-02-23T21:46:28.366Z"
                },
                "name": {
                    "type": "string",
                    "example": "Iphone 13 128GB 4GB RAM"
                },
                "price": {
                    "type": "number",
                    "example": 729.99
                },
                "restrictions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stock": {
                    "type": "integer",
                    "example": 1
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"
                }
            }
        },
        "http.SaveItemRequestHTTP": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "The iPhone 13 display has rounded corners"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://www.macstation.com.ar/img/productos/2599-2.jpg"
                    ]
                },
                "is_active": {
                    "type": "boolean",
                    "example": true
                },
                "name": {
                    "type": "string",
                    "example": "Iphone 13 128GB 4GB RAM"
                },
                "price": {
                    "type": "number",
                    "example": 729.99
                },
                "restrictions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stock": {
                    "type": "integer",
                    "example": 1
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"
                }
            }
        },
        "http.SaveItemResponseHTTP": {
            "type": "object",
            "properties": {
                "date_created": {
                    "type": "string",
                    "example": "2023-02-23T21:46:28.366Z"
                },
                "description": {
                    "type": "string",
                    "example": "The iPhone 13 display has rounded corners"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://www.macstation.com.ar/img/productos/2599-2.jpg"
                    ]
                },
                "is_active": {
                    "type": "boolean",
                    "example": true
                },
                "last_updated": {
                    "type": "string",
                    "example": "2023-02-23T21:46:28.366Z"
                },
                "name": {
                    "type": "string",
                    "example": "Iphone 13 128GB 4GB RAM"
                },
                "price": {
                    "type": "number",
                    "example": 729.99
                },
                "restrictions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stock": {
                    "type": "integer",
                    "example": 1
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"
                }
            }
        },
        "http.UpdateItemRequestHTTP": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "The iPhone 13 display has rounded corners"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://www.macstation.com.ar/img/productos/2599-2.jpg"
                    ]
                },
                "is_active": {
                    "type": "boolean",
                    "example": true
                },
                "name": {
                    "type": "string",
                    "example": "Iphone 13 128GB 4GB RAM"
                },
                "price": {
                    "type": "number",
                    "example": 729.99
                },
                "restrictions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stock": {
                    "type": "integer",
                    "example": 1
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"
                }
            }
        },
        "http.UpdateItemResponseHTTP": {
            "type": "object",
            "properties": {
                "date_created": {
                    "type": "string",
                    "example": "2023-02-23T21:46:28.366Z"
                },
                "description": {
                    "type": "string",
                    "example": "The iPhone 13 display has rounded corners"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "https://www.macstation.com.ar/img/productos/2599-2.jpg"
                    ]
                },
                "is_active": {
                    "type": "boolean",
                    "example": true
                },
                "last_updated": {
                    "type": "string",
                    "example": "2023-02-23T21:46:28.366Z"
                },
                "name": {
                    "type": "string",
                    "example": "Iphone 13 128GB 4GB RAM"
                },
                "price": {
                    "type": "number",
                    "example": 729.99
                },
                "restrictions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stock": {
                    "type": "integer",
                    "example": 1
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Items API",
	Description:      "This is an API that handles the items information.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}