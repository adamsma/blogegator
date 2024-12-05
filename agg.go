package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/adamsma/blogegator/internal/database"
	"github.com/adamsma/blogegator/internal/rss"
	"github.com/google/uuid"
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
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

}

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return errors.New("unable to fetch next feed to agg")
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

	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		msg := fmt.Errorf("unable to fetch feed %s: %w", nextFeed.Name, err)
		fmt.Println(msg)
		return nil
	}

	fmt.Printf("Aggregating post for '%s':\n", nextFeed.Name)
	newPosts := 0
	for _, item := range rssFeed.Channel.Item {

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: item.PubDate,
			FeedID:      nextFeed.ID,
		}

		_, err = s.db.CreatePost(context.Background(), params)
		if err != nil {

			if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"posts_url_key\"") {
				continue
			}

			return fmt.Errorf("unable to save post: %w", err)
		}

		newPosts++

	}

	fmt.Printf("New Posts Captured: %d\n", newPosts)

	return nil
}
