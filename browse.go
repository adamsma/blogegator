package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/adamsma/blogegator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	if len(cmd.args) > 1 {
		return errors.New("usage: browse [limit]")
	}

	limit := 2

	if len(cmd.args) == 1 {
		limitOpt, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}

		limit = limitOpt

	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)},
	)
	if err != nil {
		return fmt.Errorf("unable to fetch posts for users: %w", err)
	}

	fmt.Printf("Posts found: %d\n", len(posts))

	for _, post := range posts {
		printPost(post)
	}

	return nil
}

func printPost(p database.GetPostsForUserRow) {
	fmt.Println("-----------------------------")
	fmt.Printf("* Title:        %s\n", p.Title)
	fmt.Printf("* Feed:         %s\n", p.FeedName)
	fmt.Printf("* Published:    %s\n", p.PublishedAt)
	fmt.Printf("* Link:         %s\n", p.Url)
	fmt.Printf("* Descriptions: %s\n", p.Description)
}
