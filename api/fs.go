package api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"path/filepath"
	"proto/pb"
)

type fsGRPCApi struct {
	pb.UnimplementedFSServer
	fileDir string
}

func NewFsGRPCApi(fileDir string) pb.FSServer {
	return &fsGRPCApi{fileDir: fileDir}
}

func (s *fsGRPCApi) Upload(req grpc.ClientStreamingServer[pb.Chunk, pb.UploadResponse]) error {
	var (
		filename string
		content  []byte
	)

	// Handle Stream
	for {
		chunk, err := req.Recv()

		if err != nil {
			log.Println(err.Error())
			req.SendAndClose(&pb.UploadResponse{
				Status: pb.UploadStatus_FAILED,
			})
		}

		if len(filename) == 0 {
			filename = chunk.GetFileName()

		}

		content = append(content, chunk.Data...)

		if chunk.Done {
			break
		}
	}

	if err := os.WriteFile(filepath.Join(s.fileDir, filename), content, os.ModePerm); err != nil {
		log.Println(err.Error())

		req.SendAndClose(&pb.UploadResponse{
			Status: pb.UploadStatus_FAILED,
		})
	}

	req.SendAndClose(&pb.UploadResponse{
		Status: pb.UploadStatus_SUCCESS,
	})

	return nil
}

func (s *fsGRPCApi) Download(req *pb.DownloadRequest, stream grpc.ServerStreamingServer[pb.Chunk]) error {
	data, err := os.ReadFile(filepath.Join(s.fileDir, req.GetFileName()))

	if err != nil {
		return err
	}

	chunkSize := 500
	start, end := 0, chunkSize
	for {
		if end >= len(data) {
			end = len(data)
		}

		done := end == len(data)

		dataToSend := data[start:end]

		stream.Send(&pb.Chunk{
			Data:     dataToSend,
			FileName: req.GetFileName(),
			Done:     done,
		})

		if done {
			return nil
		}

		start, end = end, end+chunkSize

	}
}

func (s *fsGRPCApi) Echo(ctx context.Context, req *pb.EchoMessage) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		EchoMsg: req.GetMsg(),
	}, nil
}
