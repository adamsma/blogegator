package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adamsma/blogegator/internal/config"
)

type state struct {
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

	// register commands
	appCommands.register("login", handlerLogin)

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
