package main

import "errors"

// command struct to track name and args passed to a command
type command struct {
	Name string
	Args []string
}

// struct to a map of all the commands
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// Method to register a new command in the commands map
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// Method to run a commands using name and function callback from commands map
func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if ok {
		return handler(s, cmd)
	}
	return errors.New("command not found")
}
