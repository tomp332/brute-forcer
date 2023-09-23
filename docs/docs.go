// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/brute": {
            "post": {
                "description": "Start brute force task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brute Force"
                ],
                "summary": "Start new brute force action",
                "parameters": [
                    {
                        "description": "IBruteForceCreate",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.IBruteForceCreate"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.IBruteForceRead"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/creds": {
            "get": {
                "description": "Get credentials from the database",
                "tags": [
                    "Creds"
                ],
                "summary": "Get credentials",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit the number of results",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset number",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.IReadCredentials"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update credentials in the database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Creds"
                ],
                "summary": "Update credentials",
                "parameters": [
                    {
                        "description": "ICredentialsCreate",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.IUpdateCredentials"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.IReadCredentials"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    }
                }
            },
            "post": {
                "description": "Add credentials to the database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Creds"
                ],
                "summary": "Add credentials",
                "parameters": [
                    {
                        "description": "ICredentialsCreate",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ICredentialsCreate"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ICredentialsCreate"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    }
                }
            },
            "delete": {
                "description": "IDeleteParams credentials from the database",
                "tags": [
                    "Creds"
                ],
                "summary": "Delete credentials by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the credentials to delete",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerError"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "API HealthCheck",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Check"
                ],
                "summary": "API Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Health"
                        }
                    }
                }
            }
        },
        "/slaves": {
            "get": {
                "description": "Get slaves from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Slaves"
                ],
                "summary": "Get slaves",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit the number of results",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset number",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SlaveDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Add slaves to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Slaves"
                ],
                "summary": "Add slaves",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SlaveDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Health": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/models.ServiceStatus"
                }
            }
        },
        "models.IBruteForceCreate": {
            "type": "object",
            "properties": {
                "algorithm": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "numSlaves": {
                    "type": "integer"
                }
            }
        },
        "models.IBruteForceRead": {
            "type": "object",
            "properties": {
                "algorithm": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "numSlaves": {
                    "type": "integer"
                },
                "plainText": {
                    "type": "string"
                }
            }
        },
        "models.ICredentialsCreate": {
            "type": "object",
            "properties": {
                "hash": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.IReadCredentials": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.IUpdateCredentials": {
            "type": "object",
            "properties": {
                "hash": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ServerError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.ServiceStatus": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "ONLINE",
                "PENDING",
                "ERROR",
                "ShuttingDown"
            ]
        },
        "models.SlaveDTO": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Brute Forcer API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
