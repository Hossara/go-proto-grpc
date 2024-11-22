package api

import (
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
			filename = chunk.FileName

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
