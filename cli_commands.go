package main

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
	"catch": {
		name:        "catch",
		description: "Attempt to catch a given pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "See Pokedex info about a given pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "See list of caught pokemon",
		callback:    commandPokedex,
	},
}

var mapConfig = config{}

func registerCommands() error {
	Commands = commandList
	return nil
}
