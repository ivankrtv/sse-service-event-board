package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type RabbitConf struct {
	Username string
	Password string
	Host     string
	Port     int
}

type SSE struct {
	NewEventRout string
}

type Application struct {
	Port string
}

type Config struct {
	Rabbit RabbitConf
	SSE    SSE
	App    Application
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	conf := Config{}

	// set rabbit env
	conf.Rabbit.Username = os.Getenv("RABBIT_USER")
	conf.Rabbit.Password = os.Getenv("RABBIT_PASSWORD")
	conf.Rabbit.Host = os.Getenv("RABBIT_HOST")
	rabbPort, err0 := strconv.Atoi(os.Getenv("RABBIT_PORT"))
	if err0 != nil {
		log.Fatal("Failed convert rabbit port from .env to int")
	}
	conf.Rabbit.Port = rabbPort

	// set SSE env
	conf.SSE.NewEventRout = os.Getenv("SSE_NEW_EVENT_ROUT")

	// set App env
	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	conf.App.Port = appPort

	return &conf
}
