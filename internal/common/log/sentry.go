package log

import (
	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewSentryClient(SentryDsn string, Environment string, AppName string, AppVersion string) (*sentry.Client, error) {
	clientSentry, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:         SentryDsn,
		Environment: Environment,
		Release:     AppName + "@" + AppVersion,
	})

	return clientSentry, err
}

func ModifyToSentryLogger(log *zap.Logger, clientSentry *sentry.Client) (*zap.Logger, error) {
	cfg := zapsentry.Configuration{
		Level:             zapcore.ErrorLevel,
		EnableBreadcrumbs: false,
		BreadcrumbLevel:   zapcore.InfoLevel,
		Tags: map[string]string{
			"component": "system",
		},
	}

	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromClient(clientSentry))

	log = log.With(zapsentry.NewScope())

	return zapsentry.AttachCoreToLogger(core, log), err
}
