package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Ravinder2102/gator/internal/config"
	"github.com/Ravinder2102/gator/internal/database"
	_ "github.com/lib/pq"
)

// state struct keeps track of gator state with config and database
type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	// Read config file at home dir
	gatorConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Open connection to DB
	db, err := sql.Open("postgres", gatorConfig.DbURL)
	if err != nil {
		log.Fatalf("error connecting to database")
	}
	defer db.Close()
	// Query handler object to handle sql cmds
	dbQueries := database.New(db)

	// Creates state and returns address
	gatorState := &state{
		cfg: &gatorConfig,
		db:  dbQueries,
	}

	// Initialise commands struct and register command
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	// Check for correct cli usage
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Run command using the args from cli
	err = cmds.run(gatorState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
