package server

import (
	"fmt"
	"tiant-micro/app/api/common/svc"
	"tiant-micro/pkg/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Server struct {
	Engine *gin.Engine
	Limiter *rate.Limiter
	Svc    *svc.ServiceContext
}

func NewServer() (server *Server) {
	gin.SetMode(gin.ReleaseMode)
	server = &Server{
		Engine: gin.New(),
		Svc: svc.NewSVC(),
	}
	server.Limiter = rate.NewLimiter(rate.Limit(server.Svc.Config.Engine.Limite), server.Svc.Config.Engine.Limite)
	server.Engine.Use(middleware.Logger(server.Svc.Logger),middleware.Limiter(server.Limiter),gin.Recovery())
	return
}

func (s *Server) Run() {
	err := s.Engine.Run(fmt.Sprintf("%s:%d",s.Svc.Config.Engine.Address, s.Svc.Config.Engine.Port))
	if err != nil {
		s.Svc.Logger.Panicw("failed to serve", "err", err)
	}
}
