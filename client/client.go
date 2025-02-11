package main

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/f6o/napos/hands"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getCard(client hands.DealerClient) *hands.Card {
	card, err := client.DealRandomCard(context.Background(), &hands.DealRandomCardRequest{})
	if err != nil {
		log.Fatalf("could not deal card: %v", err)
	}
	return card
}

func getUniqueCards(client hands.DealerClient) map[string]bool {
	dealtCards := make(map[string]bool)
	for {
		card := getCard(client)
		cardKey := fmt.Sprintf("%v-%d", card.Suit, card.Rank)
		if dealtCards[cardKey] {
			break
		}
		dealtCards[cardKey] = true
	}
	return dealtCards
}

func main() {
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("127.0.0.1:3333", opt)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := hands.NewDealerClient(conn)

	frequency := make(map[int]int)
	for i := 0; i < 100; i++ {
		dealtCards := getUniqueCards(client)
		length := len(dealtCards)
		frequency[length]++
	}

	log.Println("Frequency distribution of unique card lengths:")
	lengths := make([]int, 0, len(frequency))
	for length := range frequency {
		lengths = append(lengths, length)
	}
	sort.Ints(lengths)
	for _, length := range lengths {
		log.Printf("Length %d: %d times", length, frequency[length])
	}
}
