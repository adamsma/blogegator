package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {

	if len(cmd.args) > 0 {
		fmt.Println("arguments ignored: command accepts no arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable retrieve users: %v", err)
	}

	if len(users) == 0 {
		fmt.Println("<no users found>")
	}

	for _, user := range users {

		var indicator string
		if user.Name == s.config.CurrentUserName {
			indicator = " (current)"
		} else {
			indicator = ""
		}

		fmt.Printf("* %s%s\n", user.Name, indicator)
	}

	return nil

}
