// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/etc": {
            "post": {
                "description": "사용자의 기타 정보를 업데이트한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Etc"
                ],
                "summary": "Update etc by user_id 사용자의 기타 정보를 업데이트한다",
                "operationId": "update-etc-by-user-id",
                "parameters": [
                    {
                        "description": "Etc",
                        "name": "etc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Etc"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Etc"
                        }
                    }
                }
            }
        },
        "/api/v1/etc/{user_id}": {
            "get": {
                "description": "사용자의 모든 기타 정보를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Etc"
                ],
                "summary": "Get all etc by user_id 사용자의 모든 기타 정보를 조회한다",
                "operationId": "get-all-etc-by-user-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Etc"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/etc/{user_id}/{key}": {
            "get": {
                "description": "사용자의 기타 정보를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Etc"
                ],
                "summary": "Get etc by user_id and key 사용자의 기타 정보를 조회한다",
                "operationId": "get-etc-by-user-id-and-key",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Key",
                        "name": "key",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Etc"
                        }
                    }
                }
            }
        },
        "/api/v1/profile": {
            "post": {
                "description": "사용자의 프로필 닉네임 정보를 업데이트한다. 중복시 에러",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Update profile nicname by ID 사용자의 프로필 닉네임 정보를 업데이트한다",
                "operationId": "update-profile-nicname-by-id",
                "parameters": [
                    {
                        "description": "Profile",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Profile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Profile"
                        }
                    }
                }
            }
        },
        "/api/v1/profile/{id}": {
            "get": {
                "description": "사용자의 프로필 정보를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get profile by ID 사용자의 프로필 정보를 조회한다",
                "operationId": "get-profile-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Profile"
                        }
                    }
                }
            }
        },
        "/api/v1/social": {
            "post": {
                "description": "사용자의 소셜 정보를 추가한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social"
                ],
                "summary": "Insert social 사용자의 소셜 정보를 추가한다",
                "operationId": "insert-social",
                "parameters": [
                    {
                        "description": "Social",
                        "name": "social",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SocialRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Social"
                        }
                    }
                }
            }
        },
        "/api/v1/social/delete": {
            "post": {
                "description": "소셜 정보를 삭제한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social"
                ],
                "summary": "Delete social by ID 소셜 정보를 삭제한다",
                "operationId": "delete-social-by-id",
                "parameters": [
                    {
                        "description": "Social",
                        "name": "social",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SocialDeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Social"
                        }
                    }
                }
            }
        },
        "/api/v1/social/{target_id}/{type}": {
            "get": {
                "description": "타겟의 소셜 정보 카운트를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social"
                ],
                "summary": "Get social count by target_id and type 타겟의 소셜 정보 카운트를 조회한다",
                "operationId": "get-social-count-by-target-id-and-type",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Target ID",
                        "name": "target_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/api/v1/social/{target_id}/{type}/{limit}/{offset}": {
            "get": {
                "description": "타겟의 소셜 정보를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social"
                ],
                "summary": "Get social by target_id and type 타겟의 소셜 정보를 조회한다",
                "operationId": "get-social-by-target-id-and-type",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Target ID",
                        "name": "target_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Social"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/social/{user_id}/{type}/{limit}/{offset}": {
            "get": {
                "description": "사용자의 소셜 정보를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social"
                ],
                "summary": "Get social by user_id and type 사용자의 소셜 정보를 조회한다",
                "operationId": "get-social-by-user-id-and-type",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Social"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/user/recovery": {
            "post": {
                "description": "사용자의 복구 코드를 통해 UUID를 변경한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Recovery user 사용자의 복구 코드를 통해 UUID를 변경한다.",
                "operationId": "recovery-user",
                "parameters": [
                    {
                        "description": "Recovery",
                        "name": "recovery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RecoveryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{user_id}/recovery": {
            "get": {
                "description": "사용자의 복구 코드를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get recovery code by user_id 사용자의 복구 코드를 조회한다",
                "operationId": "get-recovery-code-by-user-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Recovery"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{uuid}": {
            "get": {
                "description": "앱부팅시 UUID를 생성하고, UUID를 통해 사용자 정보를 조회한다. (초기화시 매번 호출)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by UUID 앱부팅시 UUID를 생성하고, UUID를 통해 사용자 정보를 조회한다",
                "operationId": "get-user-by-uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Etc": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.Profile": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Recovery": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "expired": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.RecoveryRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "model.Social": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "target_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "updated": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "vars": {
                    "type": "string"
                }
            }
        },
        "model.SocialDeleteRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.SocialRequest": {
            "type": "object",
            "properties": {
                "target_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "vars": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "updated": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "go-http-game-server API",
	Description:      "go-http-game-server API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
