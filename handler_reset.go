package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("Error deleting users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
