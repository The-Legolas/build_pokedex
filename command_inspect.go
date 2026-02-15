package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name given")
	} else if len(args) > 1 {
		return fmt.Errorf("too many pokemon names given")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Printf("%s have not been caught yet...\n", pokemon.Name)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println()
	fmt.Printf("Types:\n")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  -%s\n", typeInfo.Type.Name)
	}
	fmt.Println()

	return nil
}
