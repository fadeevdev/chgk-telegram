package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/gateway/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
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

	log.Println("Running gateway service on port", ":"+conf.Port)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":"+conf.Port, mux)
}

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	defer glog.Flush()

	if err := run(cfg); err != nil {
		glog.Fatal(err)
	}
}
