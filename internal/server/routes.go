package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"passage-worker/internal/app"
)

func registerRoutes(router *gin.Engine, application *app.Application) {
	// Docs link
	//router.LoadHTMLGlob("docs/*.html")
	//router.Static("/docs", "./docs")
	//router.GET("/docs", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})

	router.GET("/metrics", metricsHandler())
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func metricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
