package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var filePath = flag.String("file", "", "file to send")
var fileName = flag.String("filename", "", "file to receive")
var method = flag.String("method", "", "upload or download")

func main() {
	flag.Parse()

	if *method == "" {
		log.Fatal("No method specified")
	}

	client, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	fsClient := pb.NewFSClient(client)

	switch *method {
	case "upload":
		if *filePath == "" {
			log.Fatal("No file provided to upload!")
		}

		upload(fsClient)

	case "download":
		if *fileName == "" {
			log.Fatal("No file name provided to download!")
		}

		download(fsClient)

	default:
		log.Fatal("Invalid method")
	}
}

func upload(client pb.FSClient) {
	data, err := os.ReadFile(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Base(*filePath)
	stream, err := client.Upload(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	start, end := 0, 500
	chunkSize := 500
	for {
		if end >= len(data) {
			end = len(data)
		}

		done := end == len(data)

		dataToSend := data[start:end]

		stream.Send(&pb.Chunk{
			Data:     dataToSend,
			FileName: fileName,
			Done:     done,
		})

		if done {
			break
		}

		start, end = end, end+chunkSize

	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Upload status: ", resp.Status)
}

func download(client pb.FSClient) {
	var (
		filename string
		content  []byte
	)

	req, err := client.Download(context.Background(), &pb.DownloadRequest{
		FileName: *fileName,
	})

	if err != nil {
		log.Println(err.Error())
	}

	// Handle Stream
	for {
		chunk, err := req.Recv()

		if err != nil {
			log.Println(err.Error())

			req.CloseSend()
		}

		if len(filename) == 0 {
			filename = chunk.GetFileName()

		}
		content = append(content, chunk.Data...)

		if chunk.Done {
			break
		}
	}

	if err := os.WriteFile(filepath.Join("./", fmt.Sprintf("downlod-%s", filename)), content, os.ModePerm); err != nil {
		log.Println(err.Error())

		req.CloseSend()
	}

	req.CloseSend()
}
