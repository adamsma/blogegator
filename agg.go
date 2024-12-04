package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/adamsma/blogegator/internal/database"
	"github.com/adamsma/blogegator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return errors.New("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return errors.New("unable to parse time_between_reqs")
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.args[0])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatal("unable to fetch next feed to agg")
	}

	err = s.db.MarkFeedFetched(
		context.Background(),
		database.MarkFeedFetchedParams{
			LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
			ID:            nextFeed.ID,
		},
	)
	if err != nil {
		log.Printf("unable to mark feed %s as fetched: %v", nextFeed.Name, err)
	}

	feed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		msg := fmt.Errorf("unable to fetch feed %s: %w", nextFeed.Name, err)
		fmt.Println(msg)
		return
	}

	fmt.Printf("Items aggregated for '%s':\n", nextFeed.Name)
	for _, item := range feed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}

}
