package main

import "fmt"

func commandPokedex(_ string) error {
	if len(pokedex) == 0 {
		fmt.Println("You have not caught any Pokemon yet!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range pokedex {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}

	return nil
}
