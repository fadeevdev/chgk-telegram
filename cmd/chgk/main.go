package main

import (
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	newServer := chgk.New(cfg)
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(TimeoutInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChgkServiceServer(grpcServer, newServer)

	log.Println("Running grpc service on port", ":"+cfg.Port)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
}
