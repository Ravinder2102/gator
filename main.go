package main

import (
	"log"
	"os"

	"github.com/Ravinder2102/gator/internal/config"
)

// state struct keeps track of gator state
type state struct {
	cfg *config.Config
}

func main() {
	gatorConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Creates state and returns address
	gatorState := &state{
		cfg: &gatorConfig,
	}

	// Initialise commands struct and register command
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	// Check for correct cli usage
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Run command using the args from cli
	err = cmds.run(gatorState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("Error doing command")
	}

}
