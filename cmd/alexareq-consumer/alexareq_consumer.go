package main

import (
	"github.com/rudestan/hassalexarmq/pkg/msghandler"
	"github.com/rudestan/hassalexarmq/pkg/rmq"
)

func main()  {
	cfg := rmq.NewConfigFromEnv()
	rabbitMQ := rmq.NewRmq(cfg)
	msgHandler := msghandler.NewHandler("http://homeassistant.local:8123/api/alexa")

	rabbitMQ.Consume(msgHandler)
}
