// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Developer",
            "email": "intanmarsjaf@outlook.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/dare": {
            "get": {
                "description": "This is API to get get random activity",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Game of Shows Random Activity",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.BadRequestResponse"
                        }
                    }
                }
            }
        },
        "/v1/fact": {
            "get": {
                "description": "This is API to get get fact about math",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Game of Fact Math",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.BadRequestResponse"
                        }
                    }
                }
            }
        },
        "/v1/name": {
            "post": {
                "description": "This is API to get prediction from name input by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Game of Name Prediction",
                "parameters": [
                    {
                        "description": "Name Prediction",
                        "name": "Gamev1",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ExampleGameName"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.BadRequestResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.BadRequestResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "004"
                },
                "message": {
                    "type": "string",
                    "example": "[message]"
                },
                "result": {
                    "type": "string",
                    "example": "null"
                }
            }
        },
        "domain.ExampleGameName": {
            "type": "object",
            "properties": {
                "Name": {
                    "type": "string",
                    "example": "Admin"
                }
            }
        },
        "domain.SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "000"
                },
                "message": {
                    "type": "string",
                    "example": "[message]"
                },
                "result": {
                    "type": "object"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:1111",
	BasePath:    "/game",
	Schemes:     []string{"http"},
	Title:       "GOGAME",
	Description: "This page is API documentation for a little game like predict age by name",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
