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
        "/user/count": {
            "post": {
                "description": "查询用户数量",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查询用户数量",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/curd.ParamSearch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-int64"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-user_User"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "查询用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查询用户",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "skip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyList-user_User"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "description": "这里写描述 get users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-user_User"
                        }
                    }
                }
            }
        },
        "/user/search": {
            "post": {
                "description": "查询用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查询用户",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/curd.ParamSearch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyList-user_User"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "获取用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-user_User"
                        }
                    }
                }
            },
            "post": {
                "description": "修改用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-user_User"
                        }
                    }
                }
            }
        },
        "/user/{id}/delete": {
            "get": {
                "description": "删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-user_User"
                        }
                    }
                }
            }
        },
        "/user/{id}/disable": {
            "get": {
                "description": "禁用用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "禁用用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-user_User"
                        }
                    }
                }
            }
        },
        "/user/{id}/enable": {
            "get": {
                "description": "启用用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "启用用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-user_User"
                        }
                    }
                }
            }
        },
        "/user/{id}/password": {
            "post": {
                "description": "修改密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "修改密码",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/curd.ReplyData-int"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "curd.ParamSearch": {
            "type": "object",
            "properties": {
                "filter": {
                    "type": "object",
                    "additionalProperties": true
                },
                "keyword": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "skip": {
                    "type": "integer"
                },
                "sort": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        },
        "curd.ReplyData-int": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "curd.ReplyData-int64": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "curd.ReplyData-user_User": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/user.User"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "curd.ReplyList-user_User": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.User"
                    }
                },
                "error": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "admin": {
                    "type": "boolean"
                },
                "cellphone": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "disabled": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
