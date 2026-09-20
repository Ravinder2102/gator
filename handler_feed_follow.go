package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ravinder2102/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}
	feedURL := cmd.Args[0]
	// get current user
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}
	feed, err := s.db.GetFeedByName(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed_follow: %w", err)
	}
	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Printf("No feed follows found for %s.\n", currentUser.Name)
	}

	fmt.Printf("Feed follows found for %s:\n", currentUser.Name)
	for _, feed := range feedFollows {
		fmt.Printf("%s\n", feed.FeedName)
	}
	return nil
}
func printFeedFollow(user, feed string) {
	fmt.Printf("User: %s\n", user)
	fmt.Printf("Feed: %s\n", feed)
}
