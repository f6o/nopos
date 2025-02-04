package main

import (
	"log"
	"net"

	hands "github.com/f6o/napos/hands"
	"google.golang.org/grpc"
)

type delaerServer struct {
	hands.UnimplementedDealerServer
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalln("error")
	}

	grpcServer := grpc.NewServer()
	dealerServer := delaerServer{}
	hands.RegisterDealerServer(grpcServer, dealerServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed")
	}
}
