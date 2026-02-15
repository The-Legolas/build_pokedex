package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many arguments given")
	}
	fmt.Println("Your Pkedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
