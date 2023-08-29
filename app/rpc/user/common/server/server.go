package server

import (
	"context"
	"tiant-micro/app/rpc/user/common/service"
	"tiant-micro/app/rpc/user/common/svc"
	"tiant-micro/app/rpc/user/pb"
)

type Server struct {
	pb.UnimplementedUserServer
	Svc *svc.ServiceContext
}

func NewServer() *Server{
	return &Server{
		Svc: svc.NewSVC(),
	}
}

func (s *Server) UserProfile(ctx context.Context, req *pb.UserProfileRequest) (*pb.UserProfileReply, error) {
	return service.UserProfile(ctx, s.Svc, req)
}
