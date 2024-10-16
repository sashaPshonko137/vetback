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
        "/animal": {
            "get": {
                "description": "Get many animals",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animal"
                ],
                "summary": "Get many animals",
                "responses": {}
            },
            "post": {
                "description": "create animal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animal"
                ],
                "summary": "create animal",
                "parameters": [
                    {
                        "description": "animal",
                        "name": "animalInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SwaggerAnimal"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/animal/{id}": {
            "get": {
                "description": "Get an animal by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animal"
                ],
                "summary": "Get animal",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "update animal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animal"
                ],
                "summary": "update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "animal",
                        "name": "animalInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AnimalToUpdate"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete an animal by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animal"
                ],
                "summary": "Delete animal",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/appointment": {
            "get": {
                "description": "Get many appointments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Get many appointments",
                "responses": {}
            },
            "post": {
                "description": "create appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "create appointment",
                "parameters": [
                    {
                        "description": "appointment",
                        "name": "appointmentInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SwaggerAppointment"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/appointment/{id}": {
            "get": {
                "description": "Get an appointment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Get appointment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "update appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "update appointment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "appointment",
                        "name": "appointmentInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AppointmentToUpdate"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete an appointment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "Delete appointment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/diagnosis": {
            "get": {
                "description": "Get much diagnosis",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnosis"
                ],
                "summary": "Get much diagnosis",
                "responses": {}
            },
            "post": {
                "description": "create diagnosis",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnosis"
                ],
                "summary": "create diagnosis",
                "parameters": [
                    {
                        "description": "diagnosis",
                        "name": "diagnosis",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SwaggerDiagnosis"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/diagnosis/{id}": {
            "get": {
                "description": "Get an diagnosis by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnosis"
                ],
                "summary": "Get diagnosis",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "update diagnosis",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnosis"
                ],
                "summary": "update diagnosis",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "diagnosis",
                        "name": "diagnosisInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DiagnosisToUpdate"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete an diagnosis by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnosis"
                ],
                "summary": "Delete diagnosis",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/session-info": {
            "get": {
                "description": "Get your session info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get session info",
                "responses": {}
            }
        },
        "/sign-in": {
            "post": {
                "description": "Sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserToLogin"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/sign-out": {
            "get": {
                "description": "Sign out",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign out",
                "responses": {}
            }
        },
        "/sign-up": {
            "post": {
                "description": "registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/story/{id}": {
            "get": {
                "description": "Generate and send PDF file",
                "produces": [
                    "application/pdf"
                ],
                "tags": [
                    "animal"
                ],
                "summary": "Generate and send PDF file",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/treatment": {
            "get": {
                "description": "Get much treatment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "summary": "Get much treatment",
                "responses": {}
            },
            "post": {
                "description": "create treatment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "summary": "create treatment",
                "parameters": [
                    {
                        "description": "treatment",
                        "name": "treatmentInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SwaggerTreatment"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/treatment/{id}": {
            "get": {
                "description": "Get an treatment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "summary": "Get treatment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "update treatment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "summary": "update treatment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "treatment",
                        "name": "treatmentInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TreatmentToUpdate"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete an treatment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "summary": "Delete treatment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user": {
            "get": {
                "description": "Get many users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get many users",
                "responses": {}
            },
            "put": {
                "description": "Update your account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update your account",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserToUpdate"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete your account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete your account",
                "responses": {}
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get an user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.AnimalToUpdate": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "breed": {
                    "type": "string"
                },
                "diagnosisId": {
                    "type": "integer"
                },
                "doctorId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "integer"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "model.AppointmentToUpdate": {
            "type": "object",
            "properties": {
                "animalId": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "diagnosisId": {
                    "type": "integer"
                },
                "doctorId": {
                    "type": "integer"
                },
                "ownerId": {
                    "type": "integer"
                },
                "reason": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.DiagnosisToUpdate": {
            "type": "object",
            "properties": {
                "doctorId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.SwaggerAnimal": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "breed": {
                    "type": "string"
                },
                "diagnosisId": {
                    "type": "integer"
                },
                "doctorId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "integer"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "model.SwaggerAppointment": {
            "type": "object",
            "properties": {
                "animalId": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "diagnosisId": {
                    "type": "integer"
                },
                "doctorId": {
                    "type": "integer"
                },
                "ownerId": {
                    "type": "integer"
                },
                "reason": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.SwaggerDiagnosis": {
            "type": "object",
            "properties": {
                "doctorId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.SwaggerTreatment": {
            "type": "object",
            "properties": {
                "animalId": {
                    "type": "integer"
                },
                "diagnosisId": {
                    "type": "integer"
                },
                "doctorId": {
                    "type": "integer"
                },
                "finish": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "model.TreatmentToUpdate": {
            "type": "object",
            "properties": {
                "animalId": {
                    "type": "integer"
                },
                "diagnosisId": {
                    "type": "integer"
                },
                "doctorId": {
                    "type": "integer"
                },
                "finish": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "specialization": {
                    "type": "string"
                }
            }
        },
        "model.UserToLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "model.UserToUpdate": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "specialization": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "14.89",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "VET API",
	Description:      "This is a server for owners and vets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
