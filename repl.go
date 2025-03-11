package main

import (
	"fmt"

	"github.com/natretsel/blog_aggregator/internal/config"
	"github.com/natretsel/blog_aggregator/internal/database"
)

type command struct {
	name string
	args []string
}

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
	fmt.Printf("Command %v registered successfully!\n", name)
}

func (c *commands) run(s *state, cmd command) error {
	err := c.commandMap[cmd.name](s, cmd)
	fmt.Printf("Executing command %v\n", cmd.name)
	if err != nil {
		return err
	}
	return nil
}
