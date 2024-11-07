package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Brent-the-carpenter/gator/internal/config"
	"github.com/Brent-the-carpenter/gator/internal/database"
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
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:     dbQueries,
		config: &userConfig,
	}

	Commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	Commands.register("login", handlerLogin)
	Commands.register("register", handlerRegister)
	Commands.register("users", handlerGetUsers)
	Commands.register("agg", handlerAgg)
	Commands.register("reset", handlerReset)
	Commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	Commands.register("feeds", handlerListsFeeds)
	Commands.register("follow", middlewareLoggedIn(handlerFollow))
	Commands.register("following", middlewareLoggedIn(handlerListFeedFollows))
	Commands.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	Commands.register("browse", middlewareLoggedIn(handlerBrowse))
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
