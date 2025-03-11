package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/natretsel/blog_aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the name")
	}
	//Query if there is an existing user with the same name
	user, _ := s.db.GetUser(context.Background(), cmd.args[0])
	if user.Name != "" {
		return fmt.Errorf("user with name %v already exists", user.Name)
	}
	// Create if it doesn't exit
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User %v successfully created.\n", user.Name)
	return nil
}
