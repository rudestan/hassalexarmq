package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/rudestan/hassalexarmq/pkg/alexakit"
    "github.com/rudestan/hassalexarmq/pkg/rmq"
)

func HandleAlexaRequest(request alexakit.Request) (alexakit.Response, error) {
    payload, err := request.ToJson()

    if err != nil {
        return alexakit.NewPlainTextSpeechResponse(alexakit.SpeechTextFailed), err
    }

    cfg := rmq.NewConfigFromEnv()
    rabbitMQ := rmq.NewRmq(cfg)

    err = rabbitMQ.Publish(payload)
    if err != nil {
        return alexakit.NewPlainTextSpeechResponse("failed posting a message to rabbit mq"), err
    }

    return alexakit.NewPlainTextSpeechResponse(alexakit.SpeechTextConfirmation), nil
}

func main() {
    lambda.Start(HandleAlexaRequest)
}
