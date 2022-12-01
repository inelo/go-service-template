package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"go-service-template/internal/app"
	"go-service-template/internal/common/log"
	"go-service-template/internal/config"
	"go-service-template/internal/server"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var conf config.Config
var logger *zap.Logger

func init() {
	pathEnv := os.Getenv("VAULT_CONFIG_PATH")
	if len(pathEnv) == 0 {
		pathEnv = "."
	}

	cnf, err := config.LoadConfig(pathEnv)
	conf = cnf
	if err != nil {
		panic(fmt.Sprintf("Error: Cannot load config with path: %s, err: %s", pathEnv, err))
	}

	// Init logger
	logger = log.InitWithSentry(conf)

	logger.Info("Started " + conf.AppName)
}

func main() {
	defer logger.Sync()

	keepRunning := true

	consumerConf := sarama.NewConfig()

	consumerConf.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumerConf.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	consumerConf.Net.SASL.Enable = true
	consumerConf.Net.SASL.User = conf.KafkaUser
	consumerConf.Net.SASL.Password = conf.KafkaPassword

	appInst := app.NewApplication(conf, logger)

	ctx := context.TODO()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	server.StartHttp(conf, logger, &appInst)

	for keepRunning {
		select {
		case <-ctx.Done():
			logger.Sugar().Info("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			logger.Sugar().Info("terminating: via signal")
			keepRunning = false
		}
	}
}
