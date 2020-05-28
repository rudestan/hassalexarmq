package main

import (
    "errors"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/rudestan/hassalexarmq/pkg/alexakit"
    "github.com/rudestan/hassalexarmq/pkg/rmq"
    "log"
    "os"
    "strings"
)

func HandleAlexaRequest(request alexakit.Request) (alexakit.Response, error) {
    payload, err := request.ToJson()

    if err != nil {
        return alexakit.NewPlainTextSpeechResponse(alexakit.SpeechTextFailed), err
    }

    if shouldSkip(request) {
        log.Println("Request filtered: " + payload)

        return alexakit.Response{}, nil
    }

    cfg := rmq.NewConfigFromEnv()

    if cfg.Host == "" {
        return alexakit.NewPlainTextSpeechResponse(alexakit.SpeechTextFailed),
            errors.New("failed posting a message to rabbit mq, host is not configured")
    }

    rabbitMQ := rmq.NewRmq(cfg)

    err = rabbitMQ.Publish(payload)
    if err != nil {
        return alexakit.NewPlainTextSpeechResponse(alexakit.SpeechTextFailed), err
    }

    return alexakit.NewPlainTextSpeechResponse(alexakit.SpeechTextConfirmation), nil
}

func shouldSkip(request alexakit.Request) bool  {
    intentsFilter := os.Getenv("INTENTS_FILTER")

    if intentsFilter == "" {
        return false
    }

    requestIntent := request.Body.Intent.Name

    if requestIntent == "" {
        return true
    }

    intents := strings.Split(intentsFilter, "|")

    for _, intent := range intents {
        if intent == requestIntent {
            return false
        }
    }

    return true
}

func main() {
    lambda.Start(HandleAlexaRequest)
}
