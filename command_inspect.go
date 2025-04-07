package main

import (
	"errors"
	"fmt"
)

/*
Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`


Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`


*/

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

	for _, val := range pokeInfo.Stats {
		fmt.Printf(" -%s: %d\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types:")
	for _, val := range pokeInfo.Types {
		fmt.Printf(" -%s\n", val.Type.Name)
	}



	//fmt.Println("Types: " + pokeInfo.Types)


	return nil

}