package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/boutaj/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	apiUrl       string
	commands     map[string]cliCommand
	pokeResponse MapApiResponse
	cache        *pokecache.Cache
}

type MapApiResponse struct {
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
		if err := command.callback(config, text[1:]); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	trimmedFields := strings.Fields(text)
	for i, v := range trimmedFields {
		trimmedFields[i] = strings.ToLower(v)
	}
	return trimmedFields
}