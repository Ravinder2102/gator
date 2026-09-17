package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ravinder2102/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	// get current user to use to insert in feed
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	// initialise feed params
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUser.ID,
	}

	// insert feed
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("ID:   		%s\n", feed.ID)
	fmt.Printf("Created: 	%v\n", feed.CreatedAt)
	fmt.Printf("Updated: 	%v\n", feed.UpdatedAt)
	fmt.Printf("Name: 		%s\n", feed.Name)
	fmt.Printf("URL: 		%s\n", feed.Url)
	fmt.Printf("UserID: 	%s\n", feed.UserID)
}
