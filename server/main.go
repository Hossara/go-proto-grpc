package main

import (
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

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, api.NewOrderGRPCServer())
	pb.RegisterFSServer(grpcServer, api.NewFsGRPCApi("./files"))

	log.Println("Starting gRPC Server on port 8080")
	grpcServer.Serve(l)
}
