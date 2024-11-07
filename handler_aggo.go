package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Brent-the-carpenter/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.name)
	}

	timeBetweenRequest, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...", timeBetweenRequest)

	timeChannel := time.NewTicker(timeBetweenRequest)
	defer timeChannel.Stop()
	for ; ; <-timeChannel.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %w", err)
	}
	fmt.Println("Found a feed to fetch!")
	scrapeFeed(s.db, nextFeed)
	return nil
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	feed, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("couldn't mark feed %s as fetched: %v", feed.Name, err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't collect feed %s: %v", feed.Name, err)
	}

	for _, rsPost := range rssFeed.Channel.Item {
		var pubDate time.Time
		// Adjust the format string to match your date format
		if t, err := time.Parse(time.RFC1123Z, rsPost.PubDate); err == nil {
			pubDate = t
			log.Printf("couldn't parse publish date into time.Time type: %v", err)
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			FeedID:      feed.ID,
			Title:       rsPost.Title,
			Url:         rsPost.Link,
			Description: sql.NullString{String: rsPost.Description, Valid: rsPost.Description != ""},
			PublishedAt: pubDate,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("couldn't save post: %s from: %s : %v\n", rsPost.Title, rssFeed.Channel.Title, err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
