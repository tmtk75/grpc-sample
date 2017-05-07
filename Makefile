.DEFAULT_GOAL := goal

goal: server client

server: server.go
	go build -o server server.go

client: client.go
	go build -o client client.go

gen: ./proto/*.proto
	protoc --proto_path ./proto \
		--go_out=plugins=grpc:./proto \
		proto/*.proto 
