package main

import (
	"fmt"
)

func main() {

	fg := make([]*int, 2)
	output := []int{}
	g := 20
	fg = append(fg, &g)
	fmt.Println(fg)
	for _, i := range fg {
		if i != nil {
			output = append(output, *i)
		}
	}
	//
	fmt.Println(output)

	// key := 0
	// array := []int{-2, -1, 0, 1, 2, 3, 4, 5}
	// array = []int{15, -2, 2, -8, 1, 7, 10, 23}
	// tmp := []int{}

	// sum := array[0]
	// // var sum *int
	// // sum = &array[0]
	// tmp = append(tmp, array[0])
	// i := 1
	// k := i
	// for i = 1; i < len(array); i++ {
	// 	fmt.Println("----------", sum, array[i], i)
	// 	if sum > key {
	// 		fmt.Println("subarray not matches with given key value", tmp, sum, i)
	// 		tmp = []int{}
	// 		i = k + 1
	// 		sum = array[i]
	// 		tmp = append(tmp, array[i])
	// 		continue
	// 		// break
	// 	}
	// 	if sum == key {
	// 		fmt.Println("subarray found", tmp, len(tmp))
	// 		break
	// 	} else {
	// 		// fmt.Println(sum, array[i])

	// 		sum = sum + array[i]
	// 		tmp = append(tmp, array[i])
	// 	}
	// }

	//infinite for loop by changing the i value
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		i = 2
	// 	}
	// }

	// map1 := map[string][]interface{}{}
	// // err := json.Unmarshal(payload, &map1)
	// err := json.Unmarshal(input1, &map1)
	// if err != nil {
	// 	panic(err)
	// }
	// inString := ""
	// Mmap := map[string]string{}
	// for i, j := range map1 {
	// 	fmt.Println(i, "-----------------------------------------------", j)
	// 	if i == "input" {
	// 		fmt.Println("++++++++++")
	// 		inString = j[0].(string)
	// 	}
	// 	for v, ok := range j[0].(string) {
	// 		fmt.Println(reflect.TypeOf(j), v, string(ok))
	// 	}

	// 	// if v, ok := j.(map[string]string); ok {
	// 	// 	Mmap = v
	// 	// }
	// }

	// fmt.Println(Mmap, "--------------------", inString)
	// MaximumSubArray()

	// // Router := chi.NewRouter()
	// // router.

	// http://api.com/customers/id
	// request headers holds, authentication, content-type,user-agent

	// response side:
	// validating whether it is get request
	// checking the authentication
	// if authentication is failed - sends 400 status code with error response message
	// query the db with get requests url paramter
	// if data not found, statuscode 404, -- write the response message
	// if data is found, Convert db data into golang struct
	// sets content-type - application/json
	// decode the golang struct into json and sends into response with status code 200

	// POST request: http://api.com/customers/
	// request body, headers, authentication, content-type,user-agent

	// response side:
	// validating whether it is post request
	// checking the authentication
	// if authentication is failed - sends 400 status code with error response message
	// From posts request body, decode the data into gola
}

func MaximumSubArray() {
	key := 0
	array := []int{-2, -1, 0, 1, 2, 3, 4, 5}
	array = []int{15, -2, 2, -8, 1, 7, 10, 23}
	tmp := []int{}

	var sum *int
	sum = &array[0]
	tmp = append(tmp, array[0])
	for i := 1; i < len(array); i++ {
		if sum != nil && *sum > key {
			fmt.Println("subarray not matches with given key value", tmp, *sum)
			tmp = []int{}
			sum = nil
			// break
		}
		if sum != nil && *sum == key {
			fmt.Println("subarray found", tmp, len(tmp))
			break
		} else {
			// fmt.Println(sum, array[i])
			if sum == nil {
				sum = &array[i]
				tmp = append(tmp, array[i])
			}
			if sum != nil {
				*sum = *sum + array[i]
				tmp = append(tmp, array[i])
			}
		}
	}
}

// http://api.com/customers/id - GET
// http://api.com/customers/ - POST
// http://api.com/customers/id - PUT
// http://api.com/customers/id - Delete

var input1 = []byte(`{
	"input":"welcome to the world of <<params.provider>>. You will be working on <<params.project>>.",
	"params":[
		{
			"name":"provider",
			"value":"kubernetes"
		},
		{
			"name":"project",
			"value":"project-1"
		}
	]
}`)

var payload = []byte(`{
	"swagger": "2.0",
	"info": {
	  "description": "Your first API with Amazon API Gateway. This is a sample API that integrates via HTTP with our demo Pet Store endpoints",
	  "title": "PetStore"
	},
	"schemes": [
	  "https"
	],
	"paths": {
	  "/": {
		"get": {
		  "tags": [
			"pets"
		  ],
		  "description": "PetStore HTML web page containing API usage information",
		  "consumes": [
			"application/json"
		  ],
		  "produces": [
			"text/html"
		  ],
		  "responses": {
			"200": {
			  "description": "Successful operation",
			  "headers": {
				"Content-Type": {
				  "type": "string",
				  "description": "Media type of request"
				}
			  }
			}
		  },
		  "x-amazon-apigateway-integration": {
			"responses": {
			  "default": {
				"statusCode": "200",
				"responseParameters": {
				  "method.response.header.Content-Type": "'text/html'"
				},
				"responseTemplates": {
				  "text/html": "<html>\n    <head>\n        <style>\n        body {\n            color: #333;\n            font-family: Sans-serif;\n            max-width: 800px;\n            margin: auto;\n        }\n        </style>\n    </head>\n    <body>\n        <h1>Welcome to your Pet Store API</h1>\n        <p>\n            You have successfully deployed your first API. You are seeing this HTML page because the <code>GET</code> method to the root resource of your API returns this content as a Mock integration.\n        </p>\n        <p>\n            The Pet Store API contains the <code>/pets</code> and <code>/pets/{petId}</code> resources. By making a <a href=\"/$context.stage/pets/\" target=\"_blank\"><code>GET</code> request</a> to <code>/pets</code> you can retrieve a list of Pets in your API. If you are looking for a specific pet, for example the pet with ID 1, you can make a <a href=\"/$context.stage/pets/1\" target=\"_blank\"><code>GET</code> request</a> to <code>/pets/1</code>.\n        </p>\n        <p>\n            You can use a REST client such as <a href=\"https://www.getpostman.com/\" target=\"_blank\">Postman</a> to test the <code>POST</code> methods in your API to create a new pet. Use the sample body below to send the <code>POST</code> request:\n        </p>\n        <pre>\n{\n    \"type\" : \"cat\",\n    \"price\" : 123.11\n}\n        </pre>\n    </body>\n</html>"
				}
			  }
			},
			"passthroughBehavior": "when_no_match",
			"requestTemplates": {
			  "application/json": "{\"statusCode\": 200}"
			},
			"type": "mock"
		  }
		}
	  },
	  "/pets": {
		"get": {
		  "tags": [
			"pets"
		  ],
		  "summary": "List all pets",
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "type",
			  "in": "query",
			  "description": "The type of pet to retrieve",
			  "required": false,
			  "type": "string"
			},
			{
			  "name": "page",
			  "in": "query",
			  "description": "Page number of results to return.",
			  "required": false,
			  "type": "string"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "Successful operation",
			  "schema": {
				"$ref": "#/definitions/Pets"
			  },
			  "headers": {
				"Access-Control-Allow-Origin": {
				  "type": "string",
				  "description": "URI that may access the resource"
				}
			  }
			}
		  },
		  "x-amazon-apigateway-integration": {
			"responses": {
			  "default": {
				"statusCode": "200",
				"responseParameters": {
				  "method.response.header.Access-Control-Allow-Origin": "'*'"
				}
			  }
			},
			"requestParameters": {
			  "integration.request.querystring.page": "method.request.querystring.page",
			  "integration.request.querystring.type": "method.request.querystring.type"
			},
			"uri": "http://petstore.execute-api.us-east-1.amazonaws.com/petstore/pets",
			"passthroughBehavior": "when_no_match",
			"httpMethod": "GET",
			"type": "http"
		  }
		},
		"post": {
		  "tags": [
			"pets"
		  ],
		  "operationId": "CreatePet",
		  "summary": "Create a pet",
		  "consumes": [
			"application/json"
		  ],
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "in": "body",
			  "name": "NewPet",
			  "required": true,
			  "schema": {
				"$ref": "#/definitions/NewPet"
			  }
			}
		  ],
		  "responses": {
			"200": {
			  "description": "Successful operation",
			  "schema": {
				"$ref": "#/definitions/NewPetResponse"
			  },
			  "headers": {
				"Access-Control-Allow-Origin": {
				  "type": "string",
				  "description": "URI that may access the resource"
				}
			  }
			}
		  },
		  "x-amazon-apigateway-integration": {
			"responses": {
			  "default": {
				"statusCode": "200",
				"responseParameters": {
				  "method.response.header.Access-Control-Allow-Origin": "'*'"
				}
			  }
			},
			"uri": "http://petstore.execute-api.us-east-1.amazonaws.com/petstore/pets",
			"passthroughBehavior": "when_no_match",
			"httpMethod": "POST",
			"type": "http"
		  }
		},
		"options": {
		  "consumes": [
			"application/json"
		  ],
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"200": {
			  "description": "Successful operation",
			  "schema": {
				"$ref": "#/definitions/Empty"
			  },
			  "headers": {
				"Access-Control-Allow-Origin": {
				  "type": "string",
				  "description": "URI that may access the resource"
				},
				"Access-Control-Allow-Methods": {
				  "type": "string",
				  "description": "Method or methods allowed when accessing the resource"
				},
				"Access-Control-Allow-Headers": {
				  "type": "string",
				  "description": "Used in response to a preflight request to indicate which HTTP headers can be used when making the request."
				}
			  }
			}
		  },
		  "x-amazon-apigateway-integration": {
			"responses": {
			  "default": {
				"statusCode": "200",
				"responseParameters": {
				  "method.response.header.Access-Control-Allow-Methods": "'POST,GET,OPTIONS'",
				  "method.response.header.Access-Control-Allow-Headers": "'Content-Type,X-Amz-Date,Authorization,X-Api-Key'",
				  "method.response.header.Access-Control-Allow-Origin": "'*'"
				}
			  }
			},
			"passthroughBehavior": "when_no_match",
			"requestTemplates": {
			  "application/json": "{\"statusCode\": 200}"
			},
			"type": "mock"
		  }
		}
	  },
	  "/pets/{petId}": {
		"get": {
		  "tags": [
			"pets"
		  ],
		  "summary": "Info for a specific pet",
		  "operationId": "GetPet",
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "petId",
			  "in": "path",
			  "description": "The id of the pet to retrieve",
			  "required": true,
			  "type": "string"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "Successful operation",
			  "schema": {
				"$ref": "#/definitions/Pet"
			  },
			  "headers": {
				"Access-Control-Allow-Origin": {
				  "type": "string",
				  "description": "URI that may access the resource"
				}
			  }
			}
		  },
		  "x-amazon-apigateway-integration": {
			"responses": {
			  "default": {
				"statusCode": "200",
				"responseParameters": {
				  "method.response.header.Access-Control-Allow-Origin": "'*'"
				}
			  }
			},
			"requestParameters": {
			  "integration.request.path.petId": "method.request.path.petId"
			},
			"uri": "http://petstore.execute-api.us-east-1.amazonaws.com/petstore/pets/{petId}",
			"passthroughBehavior": "when_no_match",
			"httpMethod": "GET",
			"type": "http"
		  }
		},
		"options": {
		  "consumes": [
			"application/json"
		  ],
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "petId",
			  "in": "path",
			  "description": "The id of the pet to retrieve",
			  "required": true,
			  "type": "string"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "Successful operation",
			  "schema": {
				"$ref": "#/definitions/Empty"
			  },
			  "headers": {
				"Access-Control-Allow-Origin": {
				  "type": "string",
				  "description": "URI that may access the resource"
				},
				"Access-Control-Allow-Methods": {
				  "type": "string",
				  "description": "Method or methods allowed when accessing the resource"
				},
				"Access-Control-Allow-Headers": {
				  "type": "string",
				  "description": "Used in response to a preflight request to indicate which HTTP headers can be used when making the request."
				}
			  }
			}
		  },
		  "x-amazon-apigateway-integration": {
			"responses": {
			  "default": {
				"statusCode": "200",
				"responseParameters": {
				  "method.response.header.Access-Control-Allow-Methods": "'GET,OPTIONS'",
				  "method.response.header.Access-Control-Allow-Headers": "'Content-Type,X-Amz-Date,Authorization,X-Api-Key'",
				  "method.response.header.Access-Control-Allow-Origin": "'*'"
				}
			  }
			},
			"passthroughBehavior": "when_no_match",
			"requestTemplates": {
			  "application/json": "{\"statusCode\": 200}"
			},
			"type": "mock"
		  }
		}
	  }
	},
	"definitions": {
	  "Pets": {
		"type": "array",
		"items": {
		  "$ref": "#/definitions/Pet"
		}
	  },
	  "Empty": {
		"type": "object"
	  },
	  "NewPetResponse": {
		"type": "object",
		"properties": {
		  "pet": {
			"$ref": "#/definitions/Pet"
		  },
		  "message": {
			"type": "string"
		  }
		}
	  },
	  "Pet": {
		"type": "object",
		"properties": {
		  "id": {
			"type": "integer"
		  },
		  "type": {
			"type": "string"
		  },
		  "price": {
			"type": "number"
		  }
		}
	  },
	  "NewPet": {
		"type": "object",
		"properties": {
		  "type": {
			"$ref": "#/definitions/PetType"
		  },
		  "price": {
			"type": "number"
		  }
		}
	  },
	  "PetType": {
		"type": "string",
		"enum": [
		  "dog",
		  "cat",
		  "fish",
		  "bird",
		  "gecko"
		]
	  }
	},
	"x-amazon-apigateway-documentation": {
	  "version": "v2.1",
	  "createdDate": "2016-11-17T07:03:59Z",
	  "documentationParts": [
		{
		  "location": {
			"type": "API"
		  },
		  "properties": {
			"info": {
			  "description": "Your first API with Amazon API Gateway. This is a sample API that integrates via HTTP with our demo Pet Store endpoints"
			}
		  }
		},
		{
		  "location": {
			"type": "METHOD",
			"method": "GET"
		  },
		  "properties": {
			"tags": [
			  "pets"
			],
			"description": "PetStore HTML web page containing API usage information"
		  }
		},
		{
		  "location": {
			"type": "METHOD",
			"path": "/pets/{petId}",
			"method": "GET"
		  },
		  "properties": {
			"tags": [
			  "pets"
			],
			"summary": "Info for a specific pet"
		  }
		},
		{
		  "location": {
			"type": "METHOD",
			"path": "/pets",
			"method": "GET"
		  },
		  "properties": {
			"tags": [
			  "pets"
			],
			"summary": "List all pets"
		  }
		},
		{
		  "location": {
			"type": "METHOD",
			"path": "/pets",
			"method": "POST"
		  },
		  "properties": {
			"tags": [
			  "pets"
			],
			"summary": "Create a pet"
		  }
		},
		{
		  "location": {
			"type": "PATH_PARAMETER",
			"path": "/pets/{petId}",
			"method": "*",
			"name": "petId"
		  },
		  "properties": {
			"description": "The id of the pet to retrieve"
		  }
		},
		{
		  "location": {
			"type": "QUERY_PARAMETER",
			"path": "/pets",
			"method": "GET",
			"name": "page"
		  },
		  "properties": {
			"description": "Page number of results to return."
		  }
		},
		{
		  "location": {
			"type": "QUERY_PARAMETER",
			"path": "/pets",
			"method": "GET",
			"name": "type"
		  },
		  "properties": {
			"description": "The type of pet to retrieve"
		  }
		},
		{
		  "location": {
			"type": "REQUEST_BODY",
			"path": "/pets",
			"method": "POST"
		  },
		  "properties": {
			"description": "Pet object that needs to be added to the store"
		  }
		},
		{
		  "location": {
			"type": "RESPONSE",
			"method": "*",
			"statusCode": "200"
		  },
		  "properties": {
			"description": "Successful operation"
		  }
		},
		{
		  "location": {
			"type": "RESPONSE_HEADER",
			"method": "OPTIONS",
			"statusCode": "200",
			"name": "Access-Control-Allow-Headers"
		  },
		  "properties": {
			"description": "Used in response to a preflight request to indicate which HTTP headers can be used when making the request."
		  }
		},
		{
		  "location": {
			"type": "RESPONSE_HEADER",
			"method": "OPTIONS",
			"statusCode": "200",
			"name": "Access-Control-Allow-Methods"
		  },
		  "properties": {
			"description": "Method or methods allowed when accessing the resource"
		  }
		},
		{
		  "location": {
			"type": "RESPONSE_HEADER",
			"method": "*",
			"statusCode": "200",
			"name": "Access-Control-Allow-Origin"
		  },
		  "properties": {
			"description": "URI that may access the resource"
		  }
		},
		{
		  "location": {
			"type": "RESPONSE_HEADER",
			"method": "GET",
			"statusCode": "200",
			"name": "Content-Type"
		  },
		  "properties": {
			"description": "Media type of request"
		  }
		}
	  ]
	}
  }`)
