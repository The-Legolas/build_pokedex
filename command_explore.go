package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no map name given")
	} else if len(args) > 1 {
		return fmt.Errorf("too many map names given")
	}

	locationName := args[0]
	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s...", location.Name)
	fmt.Println("Found Pokemon:")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
