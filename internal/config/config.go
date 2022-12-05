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
	HttpPort           int    `mapstructure:"HTTP_PORT"`
	DbUri              string `mapstructure:"DB_URI"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetDefault("APP_NAME", "gbox-go-service-template")
	viper.SetDefault("APP_VERSION", "main")
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("LOG_OUTPUT_PATH", "stdout")
	viper.SetDefault("LOG_ERROR_OUTPUT_PATH", "stderr")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("SENTRY_DSN", "")
	viper.SetDefault("HTTP_PORT", "9876")
	viper.SetDefault("DB_URI", "mongodb://dev:dev@127.0.0.1:27017/test?authSource=admin&x-migrations-collection=_migrations&x-advisory-locking=false")

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
