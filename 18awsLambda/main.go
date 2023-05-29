package main

import (
    "log"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
)

func main() {
    lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    log.Println("hello world!")

    response := events.APIGatewayProxyResponse{
        StatusCode: 200,
    }
    return response, nil
}