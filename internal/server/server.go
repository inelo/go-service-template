package server

import (
	"fmt"
	"go-service-template/internal/app"
	"go-service-template/internal/config"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StartHttp(config config.Config, logger *zap.Logger, application *app.Application) *http.Server {
	router := gin.New()

	// Middlewares
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	registerRoutes(router, application)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.HttpPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Info("http: web server shutdown complete")
			} else {
				logger.Sugar().Errorf("http: web server closed unexpect: %s", err)
			}
		}
	}()

	return s
}
