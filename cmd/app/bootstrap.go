package main

import (
	"context"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gimmickless/keat-kit-service/configs"
	"github.com/gimmickless/keat-kit-service/pkg/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"golang.org/x/text/language"
)

const (
	dbConnectionTimeout = 10 * time.Second
)

func initConfig() {
	viper.Set("LANGUAGES", []language.Tag{language.English, language.Turkish})
	viper.Set("LANGUAGE_DEFAULT", language.English)

	var appSchema configs.AppSchema

	// Read from .env file, if it exists
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("viper.ReadInConfig (ignore in test & prod but local): .env file not readable  %v", err)
	}

	// Bind environment variables
	viper.AutomaticEnv()
	config.BindEnvs(viper.GetViper(), appSchema)

	if err := viper.Unmarshal(&appSchema); err != nil {
		log.Fatalf("viper.Unmarshal: %v", err)
	}
	configs.App = &appSchema
}

func initdb() (*mongo.Database, context.CancelFunc, func()) {
	otelMon := otelmongo.NewMonitor()
	opts := options.Client().ApplyURI(configs.App.MongoDB.URI).SetMonitor(otelMon)
	ctx, cancel := context.WithTimeout(context.Background(), dbConnectionTimeout)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	return client.Database(configs.App.MongoDB.Database), cancel, func() {
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

func initTracer() *sdktrace.TracerProvider {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	//exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		log.Fatalf("failed to disconnect db: %v", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("my-service"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}
