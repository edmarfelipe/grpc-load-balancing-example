run-server-1:
	go run server/main.go -port 1101

run-server-2:
	go run server/main.go -port 1102

run-client:
	go run client/*.go

build-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shared/hello.proto