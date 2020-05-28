package main

import (
	"github.com/rudestan/hassalexarmq/pkg/msghandler"
	"github.com/rudestan/hassalexarmq/pkg/rmq"
)

func main() {
	rabbitMQ := rmq.NewRmq(rmq.NewConfigFromEnv())
	rabbitMQ.Consume(msghandler.NewHandler())
}
