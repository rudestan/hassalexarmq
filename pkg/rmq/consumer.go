package rmq

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"log"
	"reflect"
	"time"
)

// Consume creates a new RMQ connection and starts listener to the preselect queue. Upon receiving
// a message handlerFunc will be executed with the received message payload.
func (proc* Rmq) Consume(handler interface{}) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", proc.config.Login, proc.config.Password, proc.config.Host, proc.config.Port))

	if err != nil {
		proc.logError(err, "failed to connect")

		return err
	}

	defer func() {
		if conn.IsClosed() {
			return
		}

		err := conn.Close()

		if err != nil {
			panic(err)
		}
	}()

	ch, err := proc.openChannel(conn)
	if err != nil {
		proc.logError(err, "could not create a channel")

		return err
	}

	defer func() {
		if conn.IsClosed() {
			return
		}

		err := ch.Close()

		if err != nil {
			panic(err)
		}
	}()

	err = proc.declareExchange(ch)
	if err != nil {
		proc.logError(err, "can not declare the exchange")

		return err
	}

	q, err := proc.declareQueue(ch)
	if err != nil {
		proc.logError(err, "failed to declare the queue")

		return err
	}

	err = proc.bindQueue(ch, q)
	if err != nil {
		proc.logError(err, "failed to bind the queue")

		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		proc.logError(err, "failed to register a consumer")

		return err
	}

	consumer := make(chan bool)

	go func() {
		for d := range msgs {
			handlerValue := reflect.ValueOf(handler)
			switch handlerValue.Interface().(type) {
			case MessageHandler:
				handler.(MessageHandler).Handle(cast.ToString(d.Body))
			default:
				log.Println("Wrong handler type provided!")
				return
			}
		}
	}()

	go func() {
		for {
			if conn.IsClosed() {
				consumer <- true
			}
			time.Sleep(2 * time.Second)
		}
	}()

	log.Printf("Consuming messages from %s:%d\n", proc.config.Host, proc.config.Port)

	<-consumer

	return nil
}

func (proc *Rmq) openChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (proc *Rmq) declareExchange(ch *amqp.Channel) error  {
	err := ch.ExchangeDeclare(
		proc.config.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (proc *Rmq) declareQueue(ch *amqp.Channel) (*amqp.Queue, error)  {
	q, err := ch.QueueDeclare(
		proc.config.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (proc *Rmq) bindQueue(ch *amqp.Channel, q *amqp.Queue) error  {
	err := ch.QueueBind(
		q.Name,
		proc.config.RoutingKey,
		proc.config.Exchange,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
