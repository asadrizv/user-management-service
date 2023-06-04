package publisher

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQPublisher(amqpURL string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{conn: conn, channel: channel}, nil
}

func (p *RabbitMQPublisher) Publish(exchange, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}

func (p *RabbitMQPublisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}

type UserUpdatedEvent struct {
	ID        string
	Name      string
	Email     string
	Timestamp int64
}
