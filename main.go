package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Brent-the-carpenter/blog_aggregator/internal/config"
	"github.com/Brent-the-carpenter/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {

	userConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error occurred while reading config: %v", err)
	}

	db, err := sql.Open("postgres", userConfig.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	programState := &state{
		db:     database.New(db),
		config: &userConfig,
	}

	Commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	Commands.register("login", handlerLogin)
	Commands.register("register", handleRegister)
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	err = Commands.run(programState, command{name: commandName, args: commandArgs})
	if err != nil {
		log.Fatal(err)
	}
}
