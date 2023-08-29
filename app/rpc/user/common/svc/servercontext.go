package svc

import (
	"os"
	"tiant-micro/app/rpc/user/common/config"
	"tiant-micro/pkg/argv"
	"tiant-micro/pkg/zlog"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Config *config.Config
	Logger *zap.SugaredLogger
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

	return &ServiceContext{
		Config: cfg,
		Logger: log,
	}
}