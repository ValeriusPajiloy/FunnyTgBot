package main

import (
	"errors"
	"flag"
	"log"
	"os"
	tgclient "tgbot/internal/clients/telegram"
	"tgbot/internal/consumer/event_consumer"
	"tgbot/internal/events/telegram"

	"github.com/joho/godotenv"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

var ErrTokenNotExist = errors.New("the env [TG_TOKEN] does not exist")

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	var token string
	var exist bool
	if token, exist = os.LookupEnv("APP_ADDR"); !exist {
		log.Fatal("", ErrTokenNotExist)
	}

	tgClient := tgclient.NewClient(tgBotHost, token)

	eventsWorker := telegram.NewWorker(&tgClient)

	log.Print("service started")

	consumer := event_consumer.NewConsumer(eventsWorker, eventsWorker, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

//Get token as flag
func mustToken() string {
	token := flag.String("token", "", "token for access to tgApi")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is null")
	}
	return *token
}
