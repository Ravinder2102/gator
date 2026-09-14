package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ravinder2102/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()
	// initialise user params for CreateUser
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	// Create user
	insertedUser, err := s.db.CreateUser(ctx, userParams)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	// SetUser in cfg
	err = s.cfg.SetUser(insertedUser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User created successfully:")
	printUser(insertedUser)
	return nil
}

func printUser(user database.User) {
	fmt.Printf("ID:   %v\n", user.ID)
	fmt.Printf("Name: %v\n", user.Name)
}
