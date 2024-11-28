package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"proto/api"
	"proto/pb"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(logHandler), grpc.StreamInterceptor(logStreamHandler))

	pb.RegisterOrderServiceServer(grpcServer, api.NewOrderGRPCServer())
	pb.RegisterFSServer(grpcServer, api.NewFsGRPCApi("./files"))

	log.Println("Starting gRPC Server on port 8080")
	grpcServer.Serve(l)
}

var logHandler grpc.UnaryServerInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (resp any, err error) {

	/*	if info.FullMethod != pb.FS_Echo_FullMethodName {
		return next(ctx, req)
	}*/

	fmt.Printf("Unary Req type: %T - Method: %s\n", req, info.FullMethod)
	res, err := next(ctx, req)

	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	log.Printf("Unary Resp type: %T - Method: %s\n", res, info.FullMethod)

	return res, nil
}

var logStreamHandler grpc.StreamServerInterceptor = func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, next grpc.StreamHandler) error {
	fmt.Printf("Stram Server Handler Type: %T - Method: %s\n", srv, info.FullMethod)

	fmt.Printf("Is Server Stream %v\n", info.IsServerStream)
	fmt.Printf("Is Client Stream %v\n", info.IsClientStream)

	if err := next(srv, ss); err != nil {
		log.Printf("Error: %v\n", err)
		return err
	}

	// Server stream finished
	return nil
}
