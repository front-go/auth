package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/front-go/auth/internal/repository"
	"github.com/front-go/auth/internal/service"
	"github.com/front-go/auth/pkg/auth"
)

func main() {
	repo := repository.NewRepository()

	srv := service.NewService(repo)

	grpcSrv := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcSrv, srv)
	lis, err := net.Listen("tcp", ":8095")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
