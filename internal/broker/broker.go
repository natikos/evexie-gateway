package broker

import (
	"sync"
)

type IBroker interface {
	Connect() error
	Close() error
	Publish(message []byte, routingKey string) error
}

var (
	Broker IBroker
	once   sync.Once
)

func GetBrokerInstance() (IBroker, error) {
	var err error

	once.Do(func() {
		Broker = NewRabbitMQBroker()
		err = Broker.Connect()
	})

	if err != nil {
		return nil, err
	}

	return Broker, nil
}
