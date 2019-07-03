package forgotpasswordpb

const (
	Swagger = `{
		"swagger": "2.0",
		"info": {
		  "title": "api-proto/onboarding/forgotpassword/forgotpassword.proto",
		  "version": "version not set"
		},
		"schemes": [
		  "http",
		  "https"
		],
		"consumes": [
		  "application/json"
		],
		"produces": [
		  "application/json"
		],
		"paths": {
		  "/v1/user/password/forgot": {
			"get": {
			  "operationId": "ForgotPassowrdUserService",
			  "responses": {
				"200": {
				  "description": "A successful response.",
				  "schema": {
					"$ref": "#/definitions/forgotpasswordpbForgotPasswordResponse"
				  }
				}
			  },
			  "parameters": [
				{
				  "name": "email",
				  "in": "query",
				  "required": false,
				  "type": "string"
				},
				{
				  "name": "phone",
				  "in": "query",
				  "required": false,
				  "type": "string"
				}
			  ],
			  "tags": [
				"ForgotPasswordService"
			  ]
			}
		  }
		},
		"definitions": {
		  "forgotpasswordpbForgotPasswordResponse": {
			"type": "object",
			"properties": {
			  "message": {
				"type": "string"
			  },
			  "code": {
				"type": "integer",
				"format": "int32"
			  },
			  "verify_code": {
				"type": "string"
			  }
			}
		  }
		}
	  }`
)
