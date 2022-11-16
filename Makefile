proto:
	protoc --proto_path=pkg/delivery/grpc --go_out=pkg/delivery/grpc --go_opt=paths=source_relative \
        --go-grpc_out=pkg/delivery/grpc --go-grpc_opt=paths=source_relative \
        pkg/delivery/grpc/grpc.proto