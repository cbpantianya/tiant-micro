package router

import (
	"tiant-micro/app/api/common/server"
	"tiant-micro/app/api/common/service"
)

func ServiceUserRegister(s *server.Server) {
	s.Engine.GET("/user",service.UserService(s) )
}