package main

import (
	"tiant-micro/app/api/common/router"
	"tiant-micro/app/api/common/server"
)

func main() {
	sv := server.NewServer()
	router.ServiceHealthRegister(sv)
	router.ServiceUserRegister(sv)
	sv.Run()
}
