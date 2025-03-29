package main

import (
	"errors"
	"fmt"
)

//catch pokemon on location
func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	// location, err := cfg.pokeapiClient.GetLocation(name)
	// if err != nil {
	// 	return err
	// }

	// fmt.Printf("Exploring %s...\n", location.Name)
	// fmt.Printf("Found Pokemon: \n")
	// for _, enc := range location.PokemonEncounters {
	fmt.Printf(" Throwing a Pokeball at %s\n", name)
	// }

	return nil

}