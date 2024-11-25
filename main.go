package main

import (
	"fmt"
	"log"

	"github.com/adamsma/blogegator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Unable to read configuration file: ", err)
	}

	// fmt.Println("Initial:")
	// fmt.Printf("%+v\n", cfg)
	// fmt.Println("----------")

	config.SetUser("Marcus", &cfg)

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Unable to read updated configuration file: ", err)
	}

	// fmt.Println("Updated:")
	fmt.Printf("%+v\n", cfg)
	// fmt.Println("----------")

}
