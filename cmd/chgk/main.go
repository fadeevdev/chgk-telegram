package main

import (
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	b, err := os.ReadFile("./config.yml")

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}

	newServer := chgk.New(cfg)
	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(TimeoutInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChgkServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
}
