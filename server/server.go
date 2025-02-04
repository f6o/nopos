package server

import (
	"context"
	"net"

	hands "github.com/f6o/napos/hands"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
)

type SimpleDealerServer struct {
	hands.UnimplementedDealerServer
}

func (dealer SimpleDealerServer) DealCard(ctx context.Context, req *hands.DealRequest) (*hands.Card, error) {
	rand.Seed(req.Seed)
	suits := []hands.Suits{
		hands.Suits_SUITS_SPADE,
		hands.Suits_SUITS_HEART,
		hands.Suits_SUITS_CLUB,
		hands.Suits_SUITS_DIAMOND,
	}
	card := &hands.Card{
		Suit: suits[rand.Intn(len(suits))],
		Rank: uint32(rand.Intn(13) + 1), // Rank from 1 to 13
	}
	return card, nil
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
