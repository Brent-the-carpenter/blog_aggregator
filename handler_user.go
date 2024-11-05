package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Brent-the-carpenter/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: register <username>")
	}
	name := cmd.args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})

	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)

	}

	err = s.config.SetUser(user.Name)
	if err != nil {

		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)

	return nil
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	username := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		os.Exit(1)
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("Username:%s has been logged in.\n", username)
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command does not take arguments. Usage: command")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retreving users:%w", err)
	}

	if len(users) == 0 {
		return fmt.Errorf("No users in the database currently")
	}

	for _, user := range users {

		if user.Name == s.config.CurrentUserName {
			fmt.Printf("- %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("- %s\n", user.Name)
	}

	return nil
}
func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
