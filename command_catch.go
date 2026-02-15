package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name given")
	} else if len(args) > 1 {
		return fmt.Errorf("too many Pokemon names given")
	}
	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("\nThrowing a Pokeball at %s...\n", name)
	if res > 40 {
		fmt.Printf("%s escaped!\n\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil

}
