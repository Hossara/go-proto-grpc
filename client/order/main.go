package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	orderClient := pb.NewOrderServiceClient(client)

	resp, err := orderClient.GetOrder(context.Background(), &pb.GetOrderFilter{
		Id: 1,
	})

	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(resp.Orders, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
