package db

import (
	"calendly/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var _client *mongo.Client
var _ctx *context.Context

func Init() {
	c := config.GetConfig()
	client, err := mongo.NewClient(options.Client().ApplyURI(c.GetString("db.uri")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	_client = client
	_ctx = &ctx
}

func GetDB() *mongo.Database {
	return _client.Database("a")
}
