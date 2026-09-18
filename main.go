package main

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			name: "map",
			description: "Displays the names of next 20 location areas in the Pokemon",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon",
			callback: commandMapb,
		},
	}
}

func main() {
	url := "https://pokeapi.co/api/v2/location-area"
	config := &config{
		commands: getCommands(),
		pokeResponse: ApiResponse{Next: &url},
	}
	mainREPL(config)
}