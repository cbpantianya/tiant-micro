package router

import (
	"tiant-micro/app/api/common/server"

	"github.com/gin-gonic/gin"
)

func ServiceHealthRegister(s *server.Server) {
	s.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "OK",
		})
	})
}
