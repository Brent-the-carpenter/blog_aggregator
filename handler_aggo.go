package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/Brent-the-carpenter/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.name)
	}

	time_between_reqs := cmd.args[0]

	loop_time, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...", time_between_reqs)

	timeChannel := time.NewTicker(loop_time)
	defer timeChannel.Stop()
	for ; ; <-timeChannel.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func handlerBrowse(s *state, cmd command) error {
	var limit int32
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s <optional limit>", cmd.name)
	}
	if len(cmd.args) == 0 {
		limit = 2
	} else {
		fmtLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit specified: %w", err)
		}
		limit = int32(fmtLimit)
	}

	posts, err := s.db.GetPosts(context.Background(), limit)
	if err != nil {
		return err
	}

	for _, post := range posts {

		printPost(post)
	}

	return nil

}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %w", err)
	}

	fmt.Println("Found a feed to fetch!")

	feed, err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't collect feed %s , error: %w", feed.Name, err)

	}

	for _, rsPost := range rssFeed.Channel.Item {

		pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", rsPost.PubDate) // Adjust the format string to match your date format
		if err != nil {
			fmt.Printf("couldn't parse publish date into time.Time type: %v", err)
		}

		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       rsPost.Title,
			Url:         rsPost.Link,
			Description: sql.NullString{String: rsPost.Description, Valid: rsPost.Description != ""},
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})

		if err != nil {
			fmt.Printf("couldn't save post: %s from: %s : %v\n", rsPost.Title, rssFeed.Channel.Title, err)
		}
		fmt.Println()
		fmt.Println("===========")
		fmt.Printf("Post: %s saved!\n", post.Title)
		fmt.Println("===========")

	}

	return nil
}

func printPost(post database.GetPostsRow) {
	fmt.Println()
	fmt.Println("RSS feed:", post.FeedName)
	fmt.Println("Title:", post.Title)
	fmt.Println("URL:", post.Url)
	fmt.Println("Published:", post.PublishedAt)
	fmt.Println()
	fmt.Println("Description:", post.Description)

}
