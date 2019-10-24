package main

/**
登录服务
*/
import (
	"google.golang.org/grpc"
	"other/core/login/service"
	"other/gateway/app/protobuf/app_gateway"

	"fmt"
	"log"
	"net"
)

func init() {
	//初始化全局gorm对象
	//common.InitGorm(configure.Get())
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8082))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	app_gateway.RegisterYourServiceServer(grpcServer, &service.LoginServ{})
	grpcServer.Serve(lis)
}
