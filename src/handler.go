package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Lambda executed successfully!")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Lambda executed successfully!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
