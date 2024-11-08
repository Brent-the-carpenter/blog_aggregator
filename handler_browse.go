package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Brent-the-carpenter/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.args) == 1 {
		if fmtLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = int32(fmtLimit)
		} else {

		}
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s <optional limit>", cmd.name)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		printPost(post)
	}

	return nil
}

func printPost(post database.GetPostsForUserRow) {
	fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
	fmt.Printf("--- %s ---\n", post.Title)
	fmt.Printf("    %v\n", post.Description.String)
	fmt.Printf("Link: %s\n", post.Url)
	fmt.Println("=====================================")
}
