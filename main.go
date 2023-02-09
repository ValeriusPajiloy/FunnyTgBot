package main

import (
	"context"
	"errors"
	"log"
	"os"
	tgclient "tgbot/internal/clients/telegram"
	"tgbot/internal/consumer/event_consumer"
	"tgbot/internal/events/telegram"
	"tgbot/internal/modules/notification"
	mongodb "tgbot/internal/storage/mongo_db"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

var ErrTokenNotExist = errors.New("the env [TG_TOKEN] does not exist")

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	var exist bool

	var mongoConnectionString string

	if mongoConnectionString, exist = os.LookupEnv("MONGO_CONN_STRING"); !exist {
		log.Fatal("", ErrTokenNotExist)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionString))
	if err != nil {
		log.Fatal("mongoDb not work", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("mongoDb not work", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	repo := mongodb.NewRepository(client)

	var token string
	if token, exist = os.LookupEnv("TG_TOKEN"); !exist {
		log.Fatal("", ErrTokenNotExist)
	}
	tgClient := tgclient.NewClient(tgBotHost, token, repo)
	notification := notification.NewNotification(&tgClient, repo)
	eventsWorker := telegram.NewWorker(&tgClient, notification)

	log.Print("service started")

	consumer := event_consumer.NewConsumer(eventsWorker, eventsWorker, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
