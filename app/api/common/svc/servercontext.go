package svc

import (
	"os"
	"tiant-micro/app/api/common/config"
	"tiant-micro/app/rpc/user/pb"
	"tiant-micro/pkg/argv"
	"tiant-micro/pkg/zlog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Config *config.Config
	Logger *zap.SugaredLogger
	User pb.UserClient
}

func NewSVC() *ServiceContext{
	// Read config from here
	path, err := argv.DecodeFilePath(os.Args)
	if err != nil {
		panic(err)
	}
	cfg, err := config.NewConfig(path)
	if err != nil {
		panic(err)
	}

	log, err := zlog.NewDevZLogger()
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c := pb.NewUserClient(conn)

	return &ServiceContext{
		Config: cfg,
		Logger: log,
		User: c,
	}
}