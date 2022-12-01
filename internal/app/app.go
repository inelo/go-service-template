package app

import (
	"go-service-template/internal/config"
	"go.uber.org/zap"
)

type Application struct {
}

func NewApplication(conf config.Config, logger *zap.Logger) Application {
	return Application{}
}
