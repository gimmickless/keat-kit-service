package db_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/benweissmann/memongo"
	"github.com/gimmickless/keat-kit-service/internal/app"
	"github.com/gimmickless/keat-kit-service/internal/transport/egress/db"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

const (
	ctxTimeout     = 10 // seconds
	mongoDefaultDB = "default"
)

var testDB *mongo.Database
var catgRepo app.ICategoryRepo
var ingredRepo app.IIngredientRepo
var kitRepo app.IKitRepo

func TestMain(m *testing.M) {

	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatalf("failed to start in-memory mongo db: %v", err)
	}
	mongoURI := mongoServer.URI()
	defer mongoServer.Stop()

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("failed to create mongo client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), ctxTimeout*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect mongo: %v", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("failed to ping mongo: %v", err)
	}
	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("failed to disconnect mongo: %v", err)
		}
	}()

	logger := otelzap.New(zap.NewNop()).Sugar()
	testDB = client.Database(mongoDefaultDB)
	catgRepo = db.NewCategoryRepository(logger, testDB)
	ingredRepo = db.NewIngredientRepository(logger, testDB)
	kitRepo = db.NewKitRepository(logger, testDB)
	os.Exit(m.Run())
}
