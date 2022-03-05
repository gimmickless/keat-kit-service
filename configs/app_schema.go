// pa
package configs

var App *AppSchema

type AppSchema struct {
	HTTPInbound   HTTPInboundConfig   `mapstructure:",squash"`
	HTTPOutbound  HTTPOutboundConfig  `mapstructure:",squash"`
	MongoDB       MongoDBConfig       `mapstructure:",squash"`
	Language      LanguageConfig      `mapstructure:",squash"`
	OpenTelemetry OpenTelemetryConfig `mapstructure:",squash"`
}
