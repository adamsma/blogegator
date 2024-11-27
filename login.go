package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("command login expects 'username' argument")
	}

	name := cmd.args[0]

	err := s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't login current user: %w", err)
	}

	fmt.Printf("Current user set to: %s\n", name)

	return nil

}
