package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/front-go/auth/internal/config"
	"github.com/front-go/auth/internal/repository"
	"github.com/front-go/auth/internal/service"
	"github.com/front-go/auth/pkg/auth"
)

func main() {
	cfg := config.MustLoad()

	repo := repository.NewRepository(cfg)

	srv := service.NewService(repo)

	grpcSrv := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcSrv, srv)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Service.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
