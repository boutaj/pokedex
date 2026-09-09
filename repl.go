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
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
}

func mainREPL(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if scanner.Err() != nil || len(text) == 0 {
			continue
		}
		command, ok := config.commands[text[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(config)
	}
}

func cleanInput(text string) []string {
	trimmedFields := strings.Fields(text)
	for i, v := range trimmedFields {
		trimmedFields[i] = strings.ToLower(v)
	}
	return trimmedFields
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}