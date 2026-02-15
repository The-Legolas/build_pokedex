package main

import (
	"fmt"
)

func commandHelp(_ *config, _ []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, value := range getCommands() {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println()
	return nil
}
