package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbConnectionTimeout = 10 * time.Second
)

var (
	mongoURI        = os.Getenv("MONGODB_URI")
	mongoCampaignDB = os.Getenv("MONGODB_SRV_DB")
)

func initdb() (*mongo.Database, context.CancelFunc, func()) {
	opts := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), dbConnectionTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect db: %w", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("failed to ping db: %w", err)
	}
	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("failed to disconnect db: %w", err)
		}
	}()

	return client.Database(mongoCampaignDB), cancel, func() {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("failed to disconnect db: %w", err)
		}
	}
}
