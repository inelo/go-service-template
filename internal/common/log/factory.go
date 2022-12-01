package log

import (
	"go-service-template/internal/config"
	"go.uber.org/zap"
)

func InitWithSentry(config config.Config) *zap.Logger {
	logger, err := InitLogger(config.LogLevel, config.LogOutPutPath, config.LogErrorOutPutPath)

	if err != nil {
		panic("Error: Cannot initialize logger")
	}

	if len(config.SentryDsn) > 0 {
		clientSentry, err := NewSentryClient(config.SentryDsn, config.Environment, config.AppName, config.AppVersion)

		if err != nil {
			logger.Warn("failed new client sentry", zap.Error(err))
		}

		logger, err = ModifyToSentryLogger(logger, clientSentry)
		if err != nil {
			logger.Warn("failed to init sentry", zap.Error(err))
		}
	}

	return logger
}
