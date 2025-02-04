package server

import (
	"net"

	hands "github.com/f6o/napos/hands"
	"google.golang.org/grpc"
)

type SimpleDealerServer struct {
	hands.UnimplementedDealerServer
}

func (dealer SimpleDealerServer) StartServer() error {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	hands.RegisterDealerServer(grpcServer, dealer)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
