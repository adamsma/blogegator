package main

import (
	"fmt"
	"log"

	"github.com/adamsma/blogegator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal("unable to read configuration file: ", err)
	}

	err = config.SetUser("Marcus", &cfg)
	if err != nil {
		log.Fatal("unable to set username: ", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("unable to read updated configuration file: ", err)
	}

	fmt.Printf("%+v\n", cfg)

}
