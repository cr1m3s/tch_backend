// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "[TODO]",
        "contact": {
            "name": "API Support",
            "url": "[TODO]",
            "email": "[TODO]"
        },
        "license": {
            "name": "[TODO]",
            "url": "[TODO]"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/auth/login": {
            "post": {
                "description": "requires email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "POST request for login",
                "parameters": [
                    {
                        "description": "request info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "requires username and password for registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register"
                ],
                "summary": "POST request for registration",
                "parameters": [
                    {
                        "description": "user info for sign in",
                        "name": "user_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/queries.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/auth/reset-password": {
            "post": {
                "description": "requires registred email address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reset-password"
                ],
                "summary": "POST request to update password",
                "parameters": [
                    {
                        "description": "user email for update",
                        "name": "reset-password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/open/advertisements/getall": {
            "get": {
                "description": "endpoint for getting all advertisements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisements-getall"
                ],
                "summary": "GET request to get all advertisements",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/open/advertisements/getbyid/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint to get advertisement based on it's id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "open/advertisements/getbyid/{id}"
                ],
                "summary": "GET request to get advertisement by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "advertisement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/protected/advertisement-create": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for advertisement creation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-create"
                ],
                "summary": "POST request to create advertisement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "advertisement information",
                        "name": "advertisement-create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AdvertisementInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/protected/advertisement-delete": {
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for advertisement deletion by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-delete"
                ],
                "summary": "PATCH request to delete advertisement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "advertisement id",
                        "name": "advertisement-delete",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Id"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/protected/advertisement-filter": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for getting specific advertisements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-filter"
                ],
                "summary": "POST request to get advertisement based on params in filter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "advertisement filter",
                        "name": "advertisement-filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AdvertisementFilter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/protected/advertisement-patch": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for advertisement update",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-patch"
                ],
                "summary": "PATCH request to update advertisement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "advertisement information",
                        "name": "advertisement-patch",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AdvertisementUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/protected/create-password": {
            "patch": {
                "description": "requires token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "create-password"
                ],
                "summary": "PATCH request to create new password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "new user password",
                        "name": "create-password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/protected/user-patch": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "requires valid token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-patch"
                ],
                "summary": "PATCH request to update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "user info for update",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/queries.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/protected/userinfo": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "requires valid token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userinfo"
                ],
                "summary": "Get request to see user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AdvertisementFilter": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "maxexp": {
                    "type": "string"
                },
                "minexp": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.AdvertisementInput": {
            "type": "object",
            "properties": {
                "attachment": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "experience": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "mobile_phone": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "telegram": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.AdvertisementUpdate": {
            "type": "object",
            "properties": {
                "attachment": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "experience": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "language": {
                    "type": "string"
                },
                "mobile_phone": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "telegram": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.EmailRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "models.Id": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.InLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UserPassword": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "queries.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Study marketplace API",
	Description:      "Marketplace to connect students and teachers",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
