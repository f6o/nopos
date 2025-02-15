package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/f6o/napos/db"
	hands "github.com/f6o/napos/hands"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
)

type SimpleDealerServer struct {
	hands.UnimplementedDealerServer
}

var suits []hands.Suits

func init() {
	suits = []hands.Suits{
		hands.Suits_SUITS_SPADE,
		hands.Suits_SUITS_HEART,
		hands.Suits_SUITS_CLUB,
		hands.Suits_SUITS_DIAMOND,
	}
}

func (dealer SimpleDealerServer) DealCard(ctx context.Context, req *hands.DealRequest) (*hands.Card, error) {
	rand.Seed(req.Seed)
	card := &hands.Card{
		Suit: suits[rand.Intn(len(suits))],
		Rank: uint32(rand.Intn(13) + 1), // Rank from 1 to 13
	}
	return card, nil
}

func (dealer SimpleDealerServer) DealRandomCard(ctx context.Context, req *hands.DealRandomCardRequest) (*hands.Card, error) {
	rand.Seed(uint64(time.Now().UnixNano()))
	card := &hands.Card{
		Suit: suits[rand.Intn(len(suits))],
		Rank: uint32(rand.Intn(13) + 1), // Rank from 1 to 13
	}
	err := db.SaveHandHistory(card, "user1")
	if err != nil {
		return card, err
	}
	return card, nil
}

func (dealer SimpleDealerServer) StartServer(port int) error {
	if port <= 0 || port > 65535 {
		return fmt.Errorf("invalid port number: %d", port)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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
