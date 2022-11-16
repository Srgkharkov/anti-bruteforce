package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	//"google.golang.org/grpc"
	"github.com/Srgkharkov/anti-bruteforce/pkg/delivery/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	grpcpb.UnimplementedAntiBrutforceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *grpcpb.Request) (*grpcpb.Result, error) {
	log.Printf("Received: %v", in.GetIp())
	ctx = ctx
	return &grpcpb.Result{Status: 3}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpcpb.RegisterAntiBrutforceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
