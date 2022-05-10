package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
)

func run() error {
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

	log.Println("Gateway service listening on ", ":"+os.Getenv("PORT"))

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}

func main() {

	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
