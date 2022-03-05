package configs

import (
	"golang.org/x/text/language"
)

type HTTPInboundConfig struct {
	Port string `mapstructure:"HTTP_PORT"`
}

type HTTPOutboundConfig struct {
}

type MongoDBConfig struct {
	URI      string `mapstructure:"MONGODB_URI"`
	Database string `mapstructure:"MONGODB_CAMPAIGN_DB"`
}

type LanguageConfig struct {
	Default   language.Tag `mapstructure:"LANGUAGE_DEFAULT"`
	Languages []language.Tag
}

type OpenTelemetryConfig struct {
	TracerName string `mapstructure:"OTEL_TRACER_NAME"`
}
