package main
/**
登录服务
 */
import (
	"common/gorm"
	"configure"
	"core/login/protobuf/login_pb"
	"core/login/service"
	"google.golang.org/grpc"
	"fmt"
	"log"
	"net"
)

func init()  {
	//初始化全局gorm对象
	common.InitGorm(configure.Get())
}

func main()  {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8082))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLoginServer(grpcServer, &service.LoginServ{})
	grpcServer.Serve(lis)
}
