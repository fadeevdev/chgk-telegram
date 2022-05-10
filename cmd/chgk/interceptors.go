package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

func TimeoutInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	return handler(ctx, req)
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		token := md.Get("authorization")
		for _, t := range token {
			if t != "" {
				return handler(ctx, req)
			}
		}
	}
	return nil, errors.New("not authorized")
}
