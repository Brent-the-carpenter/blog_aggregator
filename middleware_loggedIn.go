package main

import (
	"context"
	"fmt"

	"github.com/Brent-the-carpenter/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't verify user:%s , error: %w", s.config.CurrentUserName, err)
		}
		return handler(s, cmd, user)
	}

}
