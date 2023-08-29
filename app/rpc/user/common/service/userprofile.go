package service

import (
	"context"
	"tiant-micro/app/rpc/user/common/svc"
	"tiant-micro/app/rpc/user/pb"
	"time"
)

func UserProfile(ctx context.Context, stx *svc.ServiceContext, req *pb.UserProfileRequest) (*pb.UserProfileReply, error) {
	return &pb.UserProfileReply{
		UserId:   "90121",
		Nickname: "Tiant",
		Sex:      true,
		Email:    "tiant@example.com",
		CreateAt: time.Now().Unix(),
	}, nil
}
