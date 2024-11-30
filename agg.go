package main

import (
	"context"
	"fmt"

	"github.com/adamsma/blogegator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) > 0 {
		fmt.Println("arguments ignored: command accepts no arguments")
	}

	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil

}
