package main

import (
	"fmt"
)

// show all pokemon caught and registered to pokedex
func commandPokedex(cfg *config, args ...string) error {
	
	fmt.Println("Your Pokedex:")

	for pokemon := range cfg.caughtPokemon {
		fmt.Println("-",pokemon)
	}

	return nil
}