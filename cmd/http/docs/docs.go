// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/booking": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Request for a cab booking",
                "parameters": [
                    {
                        "description": "Booking request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.BookingReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.BookingID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrResp"
                        }
                    }
                }
            }
        },
        "/booking/{userID}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get the list of cab bookings previously made by the user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Booking"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrResp"
                        }
                    }
                }
            }
        },
        "/cabs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of cabs nearby from the given location",
                "parameters": [
                    {
                        "type": "number",
                        "description": "latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "longitude",
                        "name": "lon",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "distance in meters",
                        "name": "dist",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Cab"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Booking": {
            "type": "object",
            "required": [
                "from_lat",
                "from_lon",
                "pickup_time",
                "to_lat",
                "to_lon",
                "user_id"
            ],
            "properties": {
                "from_lat": {
                    "type": "number"
                },
                "from_lon": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "is_confirmed": {
                    "type": "integer"
                },
                "pickup_time": {
                    "type": "integer"
                },
                "to_lat": {
                    "type": "number"
                },
                "to_lon": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "main.BookingID": {
            "type": "object"
        },
        "main.BookingReq": {
            "type": "object",
            "required": [
                "from_lat",
                "from_lon",
                "pickup_time",
                "to_lat",
                "to_lon",
                "user_id"
            ],
            "properties": {
                "from_lat": {
                    "type": "number"
                },
                "from_lon": {
                    "type": "number"
                },
                "pickup_time": {
                    "type": "integer"
                },
                "to_lat": {
                    "type": "number"
                },
                "to_lon": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "main.Cab": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "number"
                },
                "lon": {
                    "type": "number"
                },
                "veh_no": {
                    "type": "string"
                }
            }
        },
        "main.ErrResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
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
	Host:        "localhost:8080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Cabs Servies API",
	Description: "APIs for booking and getting cabs.",
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
