package config

import "github.com/spf13/viper"

type Config struct {
	AppName            string `mapstructure:"APP_NAME"`
	AppVersion         string `mapstructure:"APP_VERSION"`
	Environment        string `mapstructure:"ENVIRONMENT"`
	LogOutPutPath      string `mapstructure:"LOG_OUTPUT_PATH"`
	LogErrorOutPutPath string `mapstructure:"LOG_ERROR_OUTPUT_PATH"`
	LogLevel           string `mapstructure:"LOG_LEVEL"`
	SentryDsn          string `mapstructure:"SENTRY_DSN"`
	HereApiKey         string `mapstructure:"HERE_API_KEY"`
	HereApiUrl         string `mapstructure:"HERE_API_URL"`
	HereApiTimeout     int    `mapstructure:"HERE_API_TIMEOUT"`
	AssistApiUrl       string `mapstructure:"ASSIST_API_URL"`
	AssistApiTimeout   int    `mapstructure:"ASSIST_API_TIMEOUT"`
	GeocoderApiUrl     string `mapstructure:"GEOCODER_API_URL"`
	GeocoderApiTimeout int    `mapstructure:"GEOCODER_API_TIMEOUT"`
	KafkaBrokers       string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopics        string `mapstructure:"KAFKA_TOPICS"`
	KafkaGroup         string `mapstructure:"KAFKA_GROUP"`
	KafkaUser          string `mapstructure:"KAFKA_USER"`
	KafkaPassword      string `mapstructure:"KAFKA_PASSWORD"`
	HttpPort           int    `mapstructure:"HTTP_PORT"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.SetDefault("APP_NAME", "gbox-workers-passage")
	viper.SetDefault("APP_VERSION", "main")
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("LOG_OUTPUT_PATH", "stdout")
	viper.SetDefault("LOG_ERROR_OUTPUT_PATH", "stderr")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("SENTRY_DSN", "")
	viper.SetDefault("HERE_API_KEY", "")
	viper.SetDefault("HERE_API_URL", "https://routematching.hereapi.com/v8/match/routelinks")
	viper.SetDefault("HERE_API_TIMEOUT", "60")
	viper.SetDefault("ASSIST_API_URL", "http://localhost:9060/api")
	viper.SetDefault("ASSIST_API_TIMEOUT", "5")
	viper.SetDefault("GEOCODER_API_URL", "http://api-internal232.dev.gbox.pl/geocoder-address/reverseGeocode")
	viper.SetDefault("GEOCODER_API_TIMEOUT", "1")
	viper.SetDefault("KAFKA_BROKERS", "localhost:9092")
	viper.SetDefault("KAFKA_TOPICS", "assist.passage")
	viper.SetDefault("KAFKA_GROUP", "passage")
	viper.SetDefault("KAFKA_USER", "api")
	viper.SetDefault("KAFKA_PASSWORD", "dev")
	viper.SetDefault("HTTP_PORT", "9876")

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
