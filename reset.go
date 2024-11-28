package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	if len(cmd.args) > 0 {
		fmt.Println("arguments ignored: command accepts no arguments")
	}

	err := s.db.ClearUsers(context.Background())
	if err != nil {
		return errors.New("unable to reset user table")
	}

	fmt.Println("User table reset successfully")

	return nil

}
