package main

import (
	"fmt"
	"goidex/internal/pokeapi"
)

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
