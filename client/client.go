package main

import (
	"context"
	"log"

	"github.com/f6o/napos/hands"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("127.0.0.1:3333", opt)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := hands.NewDealerClient(conn)

	// Deal a card
	// req := &hands.DealRequest{Seed: 12345}
	// card, err := client.DealCard(context.Background(), req)
	card, err := client.DealRandomCard(context.Background(), &hands.DealRandomCardRequest{})
	if err != nil {
		log.Fatalf("could not deal card: %v", err)
	}

	log.Printf("Dealt card: %v of rank %d", card.Suit, card.Rank)
}
