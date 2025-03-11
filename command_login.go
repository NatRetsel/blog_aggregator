package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error querying the database: %v", err)
	}
	if user.Name == "" {
		return fmt.Errorf("account with name %v does not exist", cmd.args[0])
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User %v has been set\n", s.cfg.Current_user_name)
	return nil
}
