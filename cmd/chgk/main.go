package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/database"
	"gitlab.ozon.dev/fadeevdev/homework-2/migrations"
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
	ctx := context.Background()

	db, err := database.New(ctx, &cfg.Postgres)

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "."); err != nil {
		panic(err)
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
