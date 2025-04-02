package main

import (
	"errors"
	"fmt"
	"github.com/lasantos2/pokedex/internal/pokeapi"
	"math/rand"
	"time"
)

var pokedex = make(map[string]pokeapi.Pokemon)

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

	rand.Seed(time.Now().UnixNano())
	pokemonExperience := int32(pokemon.BaseExperience)

	catchFactor := (255*4)/rand.Int31n(pokemonExperience)

	if catchFactor > 25 {
		pokedex[pokemon.Name] = pokemon

		fmt.Printf(" %s was caught!\n", pokemon.Name)
		fmt.Printf(" %s Registered in pokedex!\n", pokemon.Name)
	} else {
		fmt.Printf(" %s wasn't caught!\n", pokemon.Name)
	}

	for r, _ := range pokedex {
		fmt.Println(r)
	}

	return nil

}