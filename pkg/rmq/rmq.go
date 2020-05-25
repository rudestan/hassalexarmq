package rmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Rmq struct {
	config *Config
}

type MessageHandler interface {
	Handle(req string)
}

func NewRmq(config *Config) *Rmq  {
	return &Rmq{config:config}
}

func (proc *Rmq) connect() (*amqp.Connection, error)  {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
		proc.config.Login,
		proc.config.Password,
		proc.config.Host,
		proc.config.Port))

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func displayError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
