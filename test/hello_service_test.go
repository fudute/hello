package test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/fudute/proto_idl/helloservice"
	"google.golang.org/grpc"
)

var address = "127.0.0.1:8080"

func TestHelloService(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()
	r, err := c.Hello(ctx, &pb.HelloRequest{Name: "fudute"})
	if err != nil {
		log.Fatalf("could not hello: %v", err)
	}
	log.Println(r.GetMessage())
}
