package db

import (
	"context"
	"fmt"
	"log"

	"github.com/f6o/napos/hands"
)

func SaveHandHistory(card *hands.Card, userId String) error {
	client := NewClient()
	if err := client.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	user, err := client.User.FindFirst(User.ID.Equals(userId)).Exec(ctx)
	if err != nil {
		return fmt.Errorf("User not found: %v", err)
	}

	history, err := client.HandHistory.CreateOne(
		HandHistory.User.Link(User.ID.Equals(user.ID)),
		HandHistory.Suit.Set(int(card.Suit.Number())),
		HandHistory.Rank.Set(int(card.Rank)),
	).Exec(ctx)

	if err != nil {
		return fmt.Errorf("Failed to save hand history")
	}

	log.Printf("Hand history saved: %v", history)

	return nil
}

func CreateUser(displayName string) error {
	client := NewClient()
	if err := client.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	user, err := client.User.CreateOne(
		User.DisplayName.Set(displayName),
	).Exec(ctx)
	if err != nil {
		return fmt.Errorf("Failed to create user")
	}

	log.Printf("User created: %v", user)
	return nil
}

func ListUsers() ([]UserModel, error) {
	client := NewClient()
	if err := client.Connect(); err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	users, err := client.User.FindMany().Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to find users")
	}

	return users, nil
}
