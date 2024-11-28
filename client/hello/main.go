package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()

	client, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	echoClient := pb.NewFSClient(client)

	res, err := echoClient.Echo(context.Background(), &pb.EchoMessage{
		Msg: "Hello",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("echoed from server: %s\n", res.EchoMsg)
}
