package ping

import (
	"gin-rush-template/internal/global/otel"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (p *ModulePing) InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		otel.CustomMetrics()
		c.JSON(200, gin.H{
			"message": "pong",
			"version": "v4",
		})
	})

	r.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})
}
