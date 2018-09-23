package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	Name        string
	Description string
}

//ErrInvalidRequestError is thrown when request is invalid
var (
	ErrInvalidRequestError = errors.New("No request body found")
	ErrInvalidRequestBody  = errors.New("Invalid request body")
)

// Handler is your Lambda function handler
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrInvalidRequestError
	}

	body := request.Body
	fmt.Println("Body", body)
	var input Input
	err := json.Unmarshal([]byte(body), &input)

	if err != nil {
		fmt.Println("Error", err)
		return events.APIGatewayProxyResponse{}, ErrInvalidRequestBody
	}

	fmt.Printf("Name :%s Description: %s", input.Name, input.Description)

	return events.APIGatewayProxyResponse{
		Body:       input.toString(),
		StatusCode: 200,
	}, nil
}

func (i Input) toString() string {
	return i.Name + " : " + i.Description
}

func main() {
	lambda.Start(Handler)

	// js := "{\"Name\" : \"Shabin\", \"Description\":\"Testing\"}"
	// var input Input
	// err := json.Unmarshal([]byte(js), &input)
	// fmt.Println(err)
	// fmt.Println(input.toString())

}
