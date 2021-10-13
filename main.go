package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/fudute/proto_idl/helloservice"
	"google.golang.org/grpc"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (*HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (reply *pb.HelloReply, err error) {
	reply = &pb.HelloReply{Message: fmt.Sprintf("Hello %s, %s", req.Name, time.Now())}
	return
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloService{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
