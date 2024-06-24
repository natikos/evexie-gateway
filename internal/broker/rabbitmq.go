package broker

import (
	"api-gateway/internal/utils"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBroker struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

var _ IBroker = (*RabbitMQBroker)(nil)

func NewRabbitMQBroker() IBroker {
	return &RabbitMQBroker{}
}

func (r *RabbitMQBroker) Connect() error {
	brokerConnectionUrl := os.Getenv("BROKER_URL")
	conn, err := amqp.Dial(brokerConnectionUrl)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	if err != nil {
		conn.Close()
	}

	r.connection = conn
	r.channel = ch

	return nil
}

func (r *RabbitMQBroker) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.connection.Close()
}

func (r *RabbitMQBroker) Publish(message []byte, routingKey string) error {
	return r.channel.Publish(
		"evexie",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}
