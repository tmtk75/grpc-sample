# README
This repository gives you some examples for gRPC.

* A server in golang
* Clients in
  - golang
  - Python
  - node

Prerequisites
* protobuf
* python3
* yarn

```
$ brew install protobuf python3 yarn
```

## Getting Started
In golang, it generates server and client codes.
```
$ go get google.golang.org/grpc
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ make 
...
$ ls server client
$ ./server

$ ./client add Jane 18
$ ./client list 
2017/05/07 21:34:00 name:"Jane" age:18
```

In python
```
$ make pip-install
$ make gen-py
$ source .env

(.e) $ python client.py list
name: "Jane"
age: 18
```

In node
```
$ yarn
$ node client.js list
{ name: 'Jane', age: 18 }
```
