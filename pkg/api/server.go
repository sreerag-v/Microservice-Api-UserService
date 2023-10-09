package http

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	handler "github.com/sreerag_v/Micro-Api-Auth/pkg/api/handler"
	pb "github.com/sreerag_v/Micro-Api-Auth/pkg/api/proto"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func StartGRPCServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("Grpc Server Starting Port :::::>", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()
	go StartGRPCServer(userHandler, "50056")
	// Use logger from Gin
	engine.Use(gin.Logger())

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3002")
}
