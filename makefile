gen_proto:
	protoc -I=./ --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./pb/order.proto ./pb/user.proto

gen_fs:
	protoc -I=./ --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./pb/fileserver.proto
