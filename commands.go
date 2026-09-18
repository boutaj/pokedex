package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

func (api *ApiResponse) Get(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(body, &api)
	if err != nil {
		return err
	}
	
	return nil
}

func commandMap(config *config) error {

	api := &config.pokeResponse
	
	api.Get(*api.Next)

	for _, value := range api.Results {
		fmt.Println(value.Name)
	}

	return nil
}

func commandMapb(config *config) error {

	api := &config.pokeResponse

	if api.Previous == nil {
		return nil
	}
	
	api.Get(*api.Previous)

	for _, value := range api.Results {
		fmt.Println(value.Name)
	}

	return nil
}