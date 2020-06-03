package main

import (
	"github.com/rudestan/hassalexarmq/pkg/msghandler"
	"github.com/rudestan/hassalexarmq/pkg/rmq"
	"github.com/spf13/cast"
	"log"
	"os"
	"time"
)

const DEFAULT_RETRY_INTERVAL = 5 // seconds

func main() {
	rabbitMQ := rmq.NewRmq(rmq.NewConfigFromEnv())

	var retryInterval int

	envVal := os.Getenv("RETRY_INTERVAL")
	if envVal != "" {
		retryInterval = cast.ToInt(envVal)
	}

	if retryInterval == 0 {
		retryInterval = DEFAULT_RETRY_INTERVAL
	}

	consume(rabbitMQ, retryInterval)
}

func consume(rabbitMQ *rmq.Rmq, retryInterval int)  {
	for {
		err := rabbitMQ.Consume(msghandler.NewHandler())
		if err != nil {
			log.Println("retrying")
		}

		time.Sleep(time.Duration(retryInterval) * time.Second)
	}
}
