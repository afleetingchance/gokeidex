package main

import (
	"fmt"
	"goidex/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string) error
	config      *config
}

type config struct {
	Next     string
	Previous string
}

var Commands map[string]cliCommand

var commandList = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Lists next page of 20 location areas",
		callback:    commandMap,
		config:      &mapConfig,
	},
	"mapb": {
		name:        "mapb",
		description: "Lists previous page of 20 location areas",
		callback:    commandMapBack,
		config:      &mapConfig,
	},
	"explore": {
		name:        "explore",
		description: "Lists pokemon in a given location",
		callback:    commandExplore,
	},
}

var mapConfig = config{}

func registerCommands() error {
	Commands = commandList
	return nil
}

func commandExit(_ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ string) error {
	fmt.Printf("Welcome to the Pokedex!\n Usage:\n\n")

	for _, command := range Commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(_ string) error {
	if mapConfig.Previous != "" && mapConfig.Next == "" {
		fmt.Println("you're on the last page")
	} else {
		mapHelper(mapConfig.Next)
	}

	return nil
}

func commandMapBack(_ string) error {
	if mapConfig.Next != "" && mapConfig.Previous == "" {
		fmt.Println("you're on the first page")
	} else {
		mapHelper(mapConfig.Previous)
	}

	return nil
}

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

func mapHelper(url string) error {
	locations, next, prev, err := pokeapi.GetLocations(url)
	if err != nil {
		return err
	}

	mapConfig.Next = next
	mapConfig.Previous = prev

	for _, location := range locations {
		fmt.Println(location)
	}

	return nil
}
