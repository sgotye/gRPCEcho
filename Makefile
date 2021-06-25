server-run:
	go run server/server.go

client-run:
	go run client/client.go

generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pingpong/pingpong.proto