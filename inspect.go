package main

import (
	"fmt"
)

func commandInspect(pokemonName string) error {
	pokemon, ok := pokedex[pokemonName]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n", pokemon.Hp)
	fmt.Printf("  -attack: %d\n", pokemon.Attack)
	fmt.Printf("  -defense: %d\n", pokemon.Defense)
	fmt.Printf("  -special-attack: %d\n", pokemon.SpecialAttack)
	fmt.Printf("  -special-defense: %d\n", pokemon.SpecialDefense)
	fmt.Printf("  -speed: %d\n", pokemon.Speed)

	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  -%s\n", pokeType)
	}

	return nil
}
