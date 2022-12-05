package main

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// OPTIONAL SECTION - REMOVE MIGRATIONS IF NOT NEEDED
	if conf.DbUri != "" {
		fmt.Println("Run migrations...")

		m, err := migrate.New(
			"file://db/migrations",
			conf.DbUri)

		if err != nil {
			panic(err)
		}

		m.Up()

		version, _, _ := m.Version()
		fmt.Println("Current migration version:", version)
	}

	// END OF MIGRATIONS SECTION

	appInst := app.NewApplication(conf, logger)
	server := server.StartHttp(conf, logger, &appInst)

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
		<-sigterm
		if err := server.Shutdown(context.Background()); err != nil {
			logger.Sugar().Infof("HTTP Server Shutdown Error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	<-idleConnectionsClosed
	logger.Sugar().Info("HTTP Server Shutdown")
}
