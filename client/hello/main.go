package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api-key", "223"))

	res, err := echoClient.Echo(ctx, &pb.EchoMessage{
		Msg: "Hello",
	})

	if err != nil {
		code, ok := status.FromError(err)

		if ok {
			log.Fatalf("\nError code: %s\nMessage: %s", code.Code(), code.Message())
		}

		log.Fatal(err)
	}

	fmt.Printf("echoed from server: %s\n", res.EchoMsg)
}
