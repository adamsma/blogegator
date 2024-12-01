package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/adamsma/blogegator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.args) < 2 {
		return errors.New("usage: addFeed <name> <url>")
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to retrieve user information: %v", err)
	}

	feedParams := database.CreateFeedParams{
		ID: uuid.New(),      
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:   cmd.args[0],
		Url:    cmd.args[1],
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("unable to add feed: %v", err)
	}

	fmt.Println("Feed created successfully:")
	fStr, err := printFeed(feed, user)
	if err != nil {
		return fmt.Errorf("unable to print feed: %v", err)
	}
	fmt.Printf("%s\n", fStr)
	fmt.Println("=====================================")

	return nil
}

func printFeed(f database.Feed, u database.User) (string, error) {

	if f.UserID != u.ID {
		return "", errors.New("user id for feed must match user id for user information")
	}

	str := fmt.Sprintf("* ID:       %s\n", f.ID)
	str += fmt.Sprintf("* Created:  %v\n", f.CreatedAt)
	str += fmt.Sprintf("* Updated:  %v\n", f.UpdatedAt)
	str += fmt.Sprintf("* Name:     %s\n", f.Name)
	str += fmt.Sprintf("* URL:      %s\n", f.Url)
	str += fmt.Sprintf("* User:     %s\n", u.Name)

	return str, nil
}
