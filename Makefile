proto:
	protoc --proto_path=internal/pkg/grpc --go_out=internal/pkg/grpc --go_opt=paths=source_relative \
        --go-grpc_out=internal/pkg/grpc --go-grpc_opt=paths=source_relative \
        internal/pkg/grpc/grpc.proto
build:
	docker compose up --build app-anti-bruteforce
