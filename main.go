package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/fudute/proto_idl/helloservice"
	"google.golang.org/grpc"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (*HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (reply *pb.HelloReply, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknow host name"
		log.Println("unable to get hostname")
	}
	reply = &pb.HelloReply{Message: fmt.Sprintf("Hello %s, from %s at %s", req.Name, hostname, time.Now())}
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
