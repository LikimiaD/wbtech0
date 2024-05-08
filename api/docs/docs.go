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
        "/get_order/{order_uid}": {
            "get": {
                "description": "Retrieves details of an order by order UID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get a single order by UID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order UID",
                        "name": "order_uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Detailed order information",
                        "schema": {
                            "$ref": "#/definitions/database.Order"
                        }
                    },
                    "400": {
                        "description": "Bad request if no order found or other request error"
                    },
                    "500": {
                        "description": "Internal server error if the order cannot be loaded"
                    }
                }
            }
        },
        "/get_orders": {
            "get": {
                "description": "Retrieves a list of all orders from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get list of all orders",
                "responses": {
                    "200": {
                        "description": "A list of orders",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Order"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request if specific error on getting orders"
                    },
                    "500": {
                        "description": "Internal server error if orders cannot be loaded"
                    }
                }
            }
        }
    },
    "definitions": {
        "database.Delivery": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
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
                "order_uid": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "zip": {
                    "type": "integer"
                }
            }
        },
        "database.Item": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "chrt_id": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "item_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nm_id": {
                    "type": "integer"
                },
                "order_uid": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "rid": {
                    "type": "string"
                },
                "sale": {
                    "type": "integer"
                },
                "size": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "database.Order": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "date_created": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "delivery": {
                    "$ref": "#/definitions/database.Delivery"
                },
                "delivery_service": {
                    "type": "string"
                },
                "entry": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "internal_signature": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.Item"
                    }
                },
                "locale": {
                    "type": "string"
                },
                "oof_shard": {
                    "type": "string"
                },
                "order_uid": {
                    "type": "string"
                },
                "payment": {
                    "$ref": "#/definitions/database.Payment"
                },
                "shard_key": {
                    "type": "string"
                },
                "sm_id": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "database.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "bank": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "custom_fee": {
                    "type": "integer"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "delivery_cost": {
                    "type": "integer"
                },
                "goods_total": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "order_uid": {
                    "type": "string"
                },
                "payment_dt": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "request_id": {
                    "type": "integer"
                },
                "transaction": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "WBTech Task0 Go API",
	Description:      "API Server for demonstrate work",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
