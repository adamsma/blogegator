package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("command login expects 'username' argument")
	}

	name := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("login failed: unknown user")
		}

		return fmt.Errorf("login failed: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't login current user: %w", err)
	}

	fmt.Printf("Current user set to: %s\n", name)

	return nil

}
