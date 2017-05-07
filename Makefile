./proto/addressbook.pb.go: ./proto/addressbook.proto
	protoc --proto_path ./proto \
		--go_out=plugins=grpc:./proto \
		proto/addressbook.proto 
