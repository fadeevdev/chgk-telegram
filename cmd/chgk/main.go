package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/database"
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

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DbUser,
		cfg.Postgres.Password, cfg.Postgres.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	goose.SetBaseFS(database.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
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
