package main

import (
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/tmtk75/grpc-sample/proto"
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
	grpclog.Println(*p)
	a.m.Lock()
	defer a.m.Unlock()
	a.people = append(a.people, p)
	return &pb.NoContent{}, nil
}

func (a *ab) ListPerson(_ *pb.NoArgs, stream pb.AddressBook_ListPersonServer) error {
	a.m.Lock()
	defer a.m.Unlock()
	grpclog.Printf("number of people: %v", len(a.people))
	for _, p := range a.people {
		err := stream.Send(p)
		if err != nil {
			return err
		}
	}
	return nil
}
