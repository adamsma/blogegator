package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adamsma/blogegator/internal/config"
	"github.com/adamsma/blogegator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read configuration file: %v", err)
	}

	appState := &state{
		config: &cfg,
	}

	appCommands := commands{
		registered: map[string]func(*state, command) error{},
	}

	// establish database connection
	db, err := sql.Open("postgres", appState.config.DBURL)
	if err != nil {
		log.Fatalf("unable to establish connection to database: %v", err)
	}

	appState.db = database.New(db)

	// register commands
	appCommands.register("login", handlerLogin)
	appCommands.register("register", handlerRegister)
	appCommands.register("reset", handlerReset)
	appCommands.register("users", handlerListUsers)

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		fmt.Println("Usage: blogegator <command> [args...]")
		return
	}

	cmd := command{
		name: cliArgs[1],
		args: cliArgs[2:],
	}

	err = appCommands.run(appState, cmd)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

}
