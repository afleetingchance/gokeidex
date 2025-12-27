package main

import (
	"fmt"
	"goidex/internal/pokeapi"
	"math/rand"
)

func commandCatch(pokemonName string) error {
	pokemon, err := pokeapi.GetPokemon(pokemonName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchChance := rand.Intn(10) + (3 * (1 - pokemon.Base_experience/306))
	if catchChance >= 5 {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
