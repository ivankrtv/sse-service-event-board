package rabbit

import (
	"fmt"
	"github.com/ivankrtv/sse-service-event-board/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func InitRabbit(conf config.RabbitConf) <-chan amqp.Delivery {
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", conf.Username, conf.Password, conf.Host, conf.Port)

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatal("Error connection to Rabbit server")
	}
	//defer conn.Close()

	ch, err1 := conn.Channel()
	if err1 != nil {
		log.Fatal("Cannot connect to chanel")
	}
	//defer ch.Close()

	q, err2 := ch.QueueDeclare(
		"event-board",
		false,
		false,
		false,
		false,
		nil,
	)
	if err2 != nil {
		log.Fatal("Failed queue declare")
	}

	queue, err3 := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err3 != nil {
		log.Fatal("Failed consume queue")
	}

	return queue
}
