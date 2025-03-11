package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("error trying to reset database: %v", err)
	}
	fmt.Println("Reset successful")
	return nil
}
