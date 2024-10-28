package emailpub

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *amqp091.Connection
	ch   *amqp091.Channel
	q    amqp091.Queue
}

func NewPublisher(user, pass, host, port string) (*Publisher, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	log.Println("connecting to rabbitmq on", connStr)

	conn, err := amqp091.Dial(connStr)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := declareQueue(ch)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		conn: conn,
		ch:   ch,
		q:    q,
	}, nil
}

func (p *Publisher) Close() {
	p.ch.Close()
	p.conn.Close()
}

func declareQueue(ch *amqp091.Channel) (amqp091.Queue, error) {
	return ch.QueueDeclare(
		"email", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
}
