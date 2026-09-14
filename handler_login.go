package main

import (
	"context"
	"fmt"
)

// Login user using given arg
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	// Check if user exists in the DB
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	// SetUser
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}
