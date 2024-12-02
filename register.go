package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adamsma/blogegator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("usage: register <username>")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), newUser)

	if err != nil {

		if strings.Contains(
			err.Error(), "violates unique constraint \"users_name_key\"",
		) {
			fmt.Println("unable to register user: username already exists")
			os.Exit(1)
		}

		return fmt.Errorf("unable to register new user: %w", err)

	}

	fmt.Println("New User Created!")

	// debugging information
	fmt.Printf("  Name: %v\n", user.Name)
	fmt.Printf("  ID: %v\n", user.ID)
	fmt.Printf("  Created At: %v\n", user.CreatedAt.Format(time.DateTime))
	fmt.Printf("  Updated At: %v\n", user.UpdatedAt.Format(time.DateTime))

	handlerLogin(s, command{name: "login", args: cmd.args})

	return nil

}
