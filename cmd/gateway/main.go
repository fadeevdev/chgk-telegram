package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/gateway/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
)

func run(conf *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterChgkServiceHandlerFromEndpoint(ctx, mux, conf.GrpcServiceAddress, opts)
	if err != nil {
		return err
	}

	conf.Port = fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))

	log.Println("Gateway service listening on", conf.Port)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(conf.Port, mux)
}

func main() {
	b, err := os.ReadFile("./config.yml")

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	defer glog.Flush()

	if err := run(cfg); err != nil {
		glog.Fatal(err)
	}
}
