package main

import (
	"io"
	"log"
	"os"

	"github.com/jawher/mow.cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/tmtk75/grpc-sample/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	client := pb.NewAddressBookClient(conn)

	app := cli.App("client", "gRPC sample client")
	app.Command("list", "List all people", func(c *cli.Cmd) {
		c.Action = func() {
			list(client)
		}
	})
	app.Command("add", "Add a person", func(c *cli.Cmd) {
		var (
			name = c.String(cli.StringArg{Name: "NAME"})
			age  = c.Int(cli.IntArg{Name: "AGE"})
		)
		c.Spec = "NAME [AGE]"
		c.Action = func() {
			add(client, *name, *age)
		}
	})
	app.Run(os.Args)
}

func list(c pb.AddressBookClient) error {
	stream, err := c.ListPerson(context.Background(), new(pb.NoArgs))
	if err != nil {
		return err
	}

	for {
		p, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		grpclog.Println(p)
	}

	return nil
}

func add(c pb.AddressBookClient, name string, age int) error {
	_, err := c.AddPerson(context.Background(), &pb.Person{
		Name: name,
		Age:  int32(age),
	})
	if err != nil {
		return err
	}
	return nil
}
