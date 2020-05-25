package main

import (
    "context"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/rudestan/hassalexarmq/pkg/alexakit"
    "github.com/rudestan/hassalexarmq/pkg/rmq"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    cfg := rmq.NewConfigFromEnv()
    rabbitMQ := rmq.NewRmq(cfg)

    err := rabbitMQ.Publish(request.Body)
    if err != nil {
        return ApiGatewayProxyResponseForAlexa("failed posting a message to rabbit mq")
    }

    return ApiGatewayProxyResponseForAlexa(alexakit.SpeechTextConfirmation)
}

func ApiGatewayProxyResponseForAlexa(text string) (events.APIGatewayProxyResponse, error) {
    responseBody, err := json.Marshal(alexakit.NewPlainTextSpeechResponse(text))

    if err != nil {
        return events.APIGatewayProxyResponse{ Body: alexakit.RawFailedResponse(), StatusCode: 200}, err
    }

    return events.APIGatewayProxyResponse{ Body: string(responseBody), StatusCode: 200}, nil
}

func main() {
    lambda.Start(handleRequest)
}