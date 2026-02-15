package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no map name given")
	} else if len(args) > 1 {
		return fmt.Errorf("too many map name given")
	}

	locationName := args[0]

	entries, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s...", locationName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range entries.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
