package main

import (
	"log"
	"os"

	"github.com/adamsma/blogegator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {

	var appState state

	cfg, err := config.Read()
	if err != nil {
		log.Fatal("unable to read configuration file: ", err)
	}

	appState.config = &cfg

	appCommands := commands{
		registered: map[string]func(*state, command) error{},
	}

	// register commands
	appCommands.register("login", handlerLogin)

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		log.Fatal("missing command name")
	}

	cmd := command{
		name: cliArgs[1],
		args: cliArgs[2:],
	}

	err = appCommands.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

}
