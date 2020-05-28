package rmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func (proc *Rmq) Publish(payload string) error {
	conn, err := proc.connect()

	if err != nil {
		return err
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println("failed to close the connection")
		}
	}()

	return proc.publishToRmq(conn, payload)
}

func (proc *Rmq) publishToRmq(conn *amqp.Connection, jsonMessage string) error  {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer func() {
		err = ch.Close()
		if err != nil {
			log.Println(fmt.Sprintf("failed to close the channel: %s", err))
		}
	}()

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         []byte(jsonMessage),
		Expiration:   proc.config.Expiration,
	}

	err = ch.Publish(
		proc.config.Exchange,
		proc.config.RoutingKey,
		false,
		false,
		msg)

	if err != nil {
		return err
	}

	return nil
}
