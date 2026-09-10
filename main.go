package main

import (
	"fmt"
	"log"

	"github.com/Ravinder2102/gator/internal/config"
)

func main() {
	gatorConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %v\n", gatorConfig)

	err = gatorConfig.SetUser("Ravi")
	if err != nil {
		log.Fatalf("error Setting user: %v", err)
	}

	gatorConfig, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %v\n", gatorConfig)
}
