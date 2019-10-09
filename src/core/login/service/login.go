package service

import (
	"context"
	"core/login/protobuf/login_pb"
	"log"
)

type LoginServ struct {

}

func (LoginServ)Login(ctx context.Context,param *pb.LoginParam) (*pb.LoginReply,error){
	log.Println("login")
	return &pb.LoginReply{
		Code:                 0,
		Msg:                  "111",
	},nil
}
