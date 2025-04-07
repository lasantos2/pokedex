package main

import (
	"errors"
	"fmt"
)

//fetch information from pokedex and display some information
func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokeInfo, ok := cfg.caughtPokemon[name]

	if !ok {
		return errors.New(" That pokemon has not been caught!")
	}

	fmt.Println("Name: " + pokeInfo.Name)
	fmt.Println("Height:", pokeInfo.Height)
	fmt.Println("Weight:", pokeInfo.Weight)
	fmt.Println("Stats:")

	for _, val := range pokeInfo.Stats {
		fmt.Printf(" -%s: %d\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types:")
	for _, val := range pokeInfo.Types {
		fmt.Printf(" -%s\n", val.Type.Name)
	}

	return nil

}