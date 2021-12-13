package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/text/language"
)

const (
	dbConnectionTimeout = 10 * time.Second
)

var (
	mongoURI        = os.Getenv("MONGODB_URI")
	mongoCampaignDB = os.Getenv("MONGODB_SRV_DB")
)

func initdb() (*mongo.Database, context.CancelFunc, func()) {
	nrMon := nrmongo.NewCommandMonitor(nil)
	opts := options.Client().ApplyURI(mongoURI).SetMonitor(nrMon)
	ctx, cancel := context.WithTimeout(context.Background(), dbConnectionTimeout)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	return client.Database(mongoCampaignDB), cancel, func() {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("failed to disconnect db: %v", err)
		}
	}
}

func initI18n() {
	// TODO: This is incomplete
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("en.toml")
}
