package main

import (
	"errors"
	"fmt"
	"math/rand"
)

//catch pokemon on location
func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf(" Throwing a Pokeball at %s...\n", pokemon.Name)
	
	pokemonExperience := int32(pokemon.BaseExperience)

	catchFactor := (255*4)/rand.Int31n(pokemonExperience)

	if catchFactor > 25 {
		fmt.Printf(" %s was caught!\n", pokemon.Name)
		fmt.Printf(" %s Registered in pokedex!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf(" %s wasn't caught!\n", pokemon.Name)
	}

	return nil

}