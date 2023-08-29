package main

import (
	"fmt"
	"net"
	"tiant-micro/app/rpc/user/common/server"
	"tiant-micro/app/rpc/user/pb"
	"tiant-micro/pkg/interceptors"

	"google.golang.org/grpc"
)

func main() {

	sv := server.NewServer()
	rpcs := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
			interceptors.Logger(sv.Svc.Logger),
		),
	)
	pb.RegisterUserServer(rpcs, sv)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", sv.Svc.Config.Engine.Address, sv.Svc.Config.Engine.Port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("server listening at %v\n", listener.Addr())
	if err := rpcs.Serve(listener); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}

}
