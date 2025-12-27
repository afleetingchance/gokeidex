package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registerCommands()

	for {
		fmt.Printf("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := cleanInput(scanner.Text())
		command, ok := Commands[input[0]]
		if ok {
			if len(input) > 1 {
				command.callback(input[1])
			} else {
				command.callback("")
			}
		} else {
			fmt.Println("Unknown command")
		}
		// fmt.Printf("Your command was: %s\n", cleanInput(scanner.Text())[0])

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}
