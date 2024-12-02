package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"strings"

	"github.com/adamsma/blogegator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.args) < 1 {
		return errors.New("usage: follow <url>")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("feed not found. Use 'register' command to add")
		}

		return fmt.Errorf("unable to retrieve feed: %v", err)
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to retrieve user information: %v", err)
	}

	args := database.CreateFeedFollowParams {
		ID: uuid.New(),      
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("unable to follow feed: %v", err)
	}
	
	fmt.Println("Now Following:")
	fmt.Printf("* User: %s\n", feedFollow.UserName)
	fmt.Printf("* Feed: %s\n", feedFollow.FeedName)
	
	return nil
}

func handlerFollowing(s *state, cmd command) error {

	follows, err := s.db.GetFeedFollowsForUser(
		context.Background(), s.config.CurrentUserName,
	)
	
	if err != nil {
		return fmt.Errorf("unable to retrieve feeds being followed: %v", err)
	}

	if len(follows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Println("You are following:")
	for _, feed := range follows {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}