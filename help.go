package main

import (
	"fmt"
)

func commandHelp(_ string) error {
	fmt.Printf("Welcome to the Pokedex!\n Usage:\n\n")

	for _, command := range Commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
