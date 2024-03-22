package ping

import (
	"github.com/gin-gonic/gin"
)

func (p *ModulePing) InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"version": "v4",
		})
	})
}
