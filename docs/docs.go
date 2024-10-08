// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Ortelius Google Group",
            "email": "ortelius-dev@googlegroups.com"
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
        "/msapi/license": {
            "post": {
                "description": "Create a new License and persist it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "license"
                ],
                "summary": "Create a License",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/msapi/swagger": {
            "post": {
                "description": "Create a new Swagger and persist it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "swagger"
                ],
                "summary": "Create a Swagger",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/msapi/textfile": {
            "post": {
                "description": "Create a new Textfile and persist it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "textfile"
                ],
                "summary": "Create a Textfile",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "11.0.0",
	Host:             "localhost:8080",
	BasePath:         "/msapi/textfile",
	Schemes:          []string{},
	Title:            "Ortelius v11 Textfile Microservice",
	Description:      "RestAPI for the Readme, License and Swagger objects.  Only for new objects.  Retrieval will be done directly against the db by other microservices.\n![Release](https://img.shields.io/github/v/release/ortelius/scec-textfile?sort=semver)\n![license](https://img.shields.io/github/license/ortelius/.github)\n\n![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-textfile/build-push-chart.yml)\n[![MegaLinter](https://github.com/ortelius/scec-textfile/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-textfile/actions?query=workflow%3AMegaLinter+branch%3Amain)\n![CodeQL](https://github.com/ortelius/scec-textfile/workflows/CodeQL/badge.svg)\n[![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-textfile/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-textfile)\n\n![Discord](https://img.shields.io/discord/722468819091849316)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
