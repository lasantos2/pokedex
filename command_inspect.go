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

	pokeInfo := cfg.caughtPokemon[name]

	// if pokeInfo == nil {
	// 	fmt.Println("%s has not been caught", name)
	// 	return nil, errors.NewError("%s has not been caught", name)
	// }

	fmt.Println("Name: " + pokeInfo.Name)
	fmt.Println("Height:", pokeInfo.Height)
	fmt.Println("Weight:", pokeInfo.Weight)
	fmt.Println("Stats:")



	//fmt.Println("Types: " + pokeInfo.Types)


	return nil

}