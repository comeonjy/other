package service

import (
	"context"
	"fmt"
	"log"
	"other/gateway/app/protobuf/app_gateway"
)

type LoginServ struct {

}

//Echo 回声服务器
func (s LoginServ) Echo(ctx context.Context,v *app_gateway.StringMessage) (*app_gateway.StringMessage, error){
	log.Println("login")
	return &app_gateway.StringMessage{
		Value: fmt.Sprintf("msg:%s",v.Value),
	},nil
}
