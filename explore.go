package main

import (
	"fmt"
	"goidex/internal/pokeapi"
)

func commandExplore(location string) error {
	pokemon, err := pokeapi.GetPokemonFromLocation(location)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Exploring %s...\n", location)

	if len(pokemon) > 0 {
		fmt.Println("Found Pokemon:")
		for _, poke := range pokemon {
			fmt.Printf(" - %s\n", poke)
		}
	} else {
		fmt.Println("There are no Pokemon here")
	}

	return nil
}
