package main

import (
	"context"
	"core/login/protobuf/login_pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestGrpc(t *testing.T)  {
	conn, err := grpc.Dial(":8082",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLoginClient(conn)

	feature, err := client.Login(context.Background(), &pb.LoginParam{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feature.Msg)
}
