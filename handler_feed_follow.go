package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ravinder2102/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}
	feedURL := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
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

func handlerFollowing(s *state, cmd command, user database.User) error {

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Printf("No feed follows found for %s.\n", user.Name)
		return nil
	}

	fmt.Printf("Feed follows found for %s:\n", user.Name)
	for _, feed := range feedFollows {
		fmt.Printf("%s\n", feed.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow record: %w", err)
	}

	fmt.Printf("%s unfollowed successfully\n", feed.Name)
	return nil
}

func printFeedFollow(user, feed string) {
	fmt.Printf("User: %s\n", user)
	fmt.Printf("Feed: %s\n", feed)
}
