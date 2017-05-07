.DEFAULT_GOAL := goal

goal: server client

server: server.go proto/addressbook.pb.go
	go build -o server server.go

client: client.go proto/addressbook.pb.go
	go build -o client client.go

proto/addressbook.pb.go: gen-go

gen-go: ./proto/*.proto
	protoc --proto_path ./proto \
		--go_out=plugins=grpc:./proto \
		proto/*.proto 
gen-py:
	.e/bin/python -m grpc_tools.protoc \
		--proto_path proto \
		--python_out=. \
		--grpc_python_out=. \
		./proto/*.proto

gen-js:
	node_modules/.bin/grpc_tools_node_protoc \
		--proto_path=./proto \
		--js_out=import_style=commonjs,binary:. \
		--grpc_out=. \
		--plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` \
		proto/*.proto

pip-install: .e/bin/activate
	.e/bin/pip install -r requirements.txt

.e/bin/activate:
	python3 -m venv .e

brew-install:
	brew install protobuf yarn python3

clean:
	rm -rf proto/*.pb.go *_pb*.* client server __pycache__

distclean: clean
	rm -rf .e node_modules
