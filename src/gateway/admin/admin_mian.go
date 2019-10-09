package main

import (
	"common/gorm"
	"configure"
	"core/login/protobuf/login_pb"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

func init() {
	common.InitGorm(configure.Get())
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		v1.GET("/login", func(context *gin.Context) {
			conn, err := grpc.Dial(":8082",grpc.WithInsecure())
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			client := pb.NewLoginClient(conn)

			feature, err := client.Login(context, &pb.LoginParam{})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(feature.Msg)
		})
	}

	if err := router.Run(":8083"); err != nil {
		log.Fatal(err)
	}
}
