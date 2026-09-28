package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Ravinder2102/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	log.Printf("Collecting feeds every %s...", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		ArgLimit, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = ArgLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Limit:  int32(limit),
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't list posts for user %s: %w", user.Name, err)
	}

	if len(posts) == 0 {
		fmt.Printf("No posts found for %s\n", user.Name)
		return nil
	}

	fmt.Printf("%d posts found for %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("Published At: %s\n", post.PublishedAt.Time.Format("Tue Feb 22"))
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("Description: %v\n", post.Description.String)
		fmt.Println("=======================================================")
	}

	return nil
}
