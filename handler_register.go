package main

import (
	"context"
	"fmt"
	"github.com/Brent-the-carpenter/blog_aggregator/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: register <username>")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		os.Exit(1)

	}

	err = s.config.SetUser(user.Name)
	if err != nil {

		return err
	}

	fmt.Printf("User:%s was successfully created\n", user.Name)
	fmt.Println("UserId:", user.ID)
	fmt.Println("Created at:", user.CreatedAt)
	fmt.Println("Updated at:", user.UpdatedAt)

	return nil
}
