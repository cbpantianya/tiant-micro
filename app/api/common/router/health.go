package router

import (
	"tiant-micro/app/api/common/server"
	"tiant-micro/app/api/common/service"
)

func ServiceHealthRegister(s *server.Server) {
	s.Engine.GET("/health", service.HealthService(s))
}
