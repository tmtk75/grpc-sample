package main

import (
	"golang.org/x/net/context"
	"log"
	"net"
	"sync"

	pb "github.com/tmtk75/grpc-sample/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddressBookServer(s, new(ab))
	s.Serve(l)
}

type ab struct {
	people []*pb.Person
	m      sync.Mutex
}

func (a *ab) AddPerson(ctx context.Context, p *pb.Person) (*pb.NoContent, error) {
	return &pb.NoContent{}, nil
}

func (a *ab) ListPerson(_ *pb.NoArgs, stream pb.AddressBook_ListPersonServer) error {
	return nil
}
