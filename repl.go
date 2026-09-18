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
	pokeResponse ApiResponse
}

type ApiResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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