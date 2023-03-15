package main

import (
	sse "github.com/alexandrevicenzi/go-sse"
	config "github.com/ivankrtv/sse-service-event-board/config"
	"github.com/ivankrtv/sse-service-event-board/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/cors"
	"log"
	"net/http"
)

var conf *config.Config

func sendEvent(s *sse.Server, m amqp.Delivery) {
	msg := string(m.Body)
	s.SendMessage(conf.SSE.NewEventRout, sse.SimpleMessage(msg))
}

func main() {
	conf = config.NewConfig()
	queue := rabbit.InitRabbit(conf.Rabbit)
	s := sse.NewServer(nil)

	go func() {
		for msg := range queue {
			sendEvent(s, msg)
		}
	}()

	serv := http.NewServeMux()
	serv.Handle(conf.SSE.NewEventRout, s)

	handler := cors.AllowAll().Handler(serv)
	err := http.ListenAndServe(conf.App.Port, handler)
	if err != nil {
		log.Print(err)
		log.Fatal("Server can not start on :8000")
	}

	for {
	}
}
