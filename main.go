package main

import (
	"fmt"
	sse "github.com/alexandrevicenzi/go-sse"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"os"
	"strconv"
)

func initEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func initRabbit() <-chan amqp.Delivery {
	username := os.Getenv("RABBIT_USER")
	password := os.Getenv("RABBIT_PASSWORD")
	host := os.Getenv("RABBIT_HOST")
	port, err0 := strconv.Atoi(os.Getenv("RABBIT_PORT"))
	if err0 != nil {
		log.Fatal("Failed convert port to int")
	}

	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port)
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

func sendEvent(s *sse.Server, msg amqp.Delivery) {
	log.Print(msg)
}

func main() {
	initEnvs()
	queue := initRabbit()
	s := sse.NewServer(nil)
	//defer s.Shutdown()
	http.Handle("/new-events", s)

	go func() {
		for msg := range queue {
			sendEvent(s, msg)
		}
	}()

	for {
	}
}
