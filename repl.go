package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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
	}
}

func mainREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if scanner.Err() != nil || len(text) == 0 {
			continue
		}
		command, ok := getCommands()[text[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}
}

func cleanInput(text string) []string {
	trimmedFields := strings.Fields(text)
	for i, v := range trimmedFields {
		trimmedFields[i] = strings.ToLower(v)
	}
	return trimmedFields
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}