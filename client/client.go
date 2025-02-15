package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
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

func printFreq(client hands.DealerClient) {
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

func addUser(client hands.GameManagerClient) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter display name: ")
	displayName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("could not read display name: %v", err)
	}
	displayName = displayName[:len(displayName)-1] // Remove the newline character

	user := &hands.User{
		DisplayName: displayName,
	}

	if _, err := client.AddUser(context.Background(), &hands.AddUserRequest{User: user}); err != nil {
		log.Fatalf("could not add user: %v", err)
	}
}

func listUsers(client hands.GameManagerClient) {
	users, err := client.ListUsers(context.Background(), &hands.ListUsersRequest{})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}

	log.Println("Users:")
	for _, user := range users.Users {
		log.Printf("%10s %s", user.Id, user.DisplayName)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("expected 'freq' subcommand")
	}

	subcommand := os.Args[1]

	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("127.0.0.1:3333", opt)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	switch subcommand {
	case "print-freq":
		client := hands.NewDealerClient(conn)
		printFreq(client)
	case "add-user":
		client := hands.NewGameManagerClient(conn)
		addUser(client)
	case "list-users":
		client := hands.NewGameManagerClient(conn)
		listUsers(client)
	default:
		log.Fatalf("unknown subcommand: %s", subcommand)
	}
}
