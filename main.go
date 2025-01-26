package main

import (
	"log"
	"net"

	hands "github.com/f6o/napos/hands"
	"google.golang.org/grpc"
)

type handServer {
	Unimplemented
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalln("error")
	}

	grpcServer := grpc.NewServer()
	hands.RegisterDealerServer(grpcServer,)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed")
	}
}
