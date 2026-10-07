package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"math/rand/v2"

	"github.com/boutaj/pokedex/internal/pokecache"
)

func commandExit(config *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func (api *MapApiResponse) Get(url string, cache *pokecache.Cache) error {
	if body, ok := cache.Get(url); ok {
		return json.Unmarshal(body, api)
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("fetching locations: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, api)
	if err != nil {
		return err
	}
	cache.Add(url, body)

	return nil
}

func commandMap(config *config, args []string) error {

	api := &config.pokeResponse

	if api.Next == nil {
		return nil
	}
	if err := api.Get(*api.Next, config.cache); err != nil {
		return err
	}

	for _, value := range api.Results {
		fmt.Println(value.Name)
	}

	return nil
}

func commandMapb(config *config, args []string) error {

	api := &config.pokeResponse

	if api.Previous == nil {
		return nil
	}

	if err := api.Get(*api.Previous, config.cache); err != nil {
		return err
	}

	for _, value := range api.Results {
		fmt.Println(value.Name)
	}

	return nil
}

type ExploreApiResponse struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []interface{} `json:"condition_values"`
				PokemonDetails  interface{}   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (api *ExploreApiResponse) Get(url string, cache *pokecache.Cache) error {
	if body, ok := cache.Get(url); ok {
		return json.Unmarshal(body, api)
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Exploring locations: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, api)
	if err != nil {
		return err
	}
	cache.Add(url, body)

	return nil
}

func commandExplore(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("usage: explore <area_name>")
		return nil
	}
	location := args[0]
	
	api := &ExploreApiResponse{}
	if err := api.Get(config.apiUrl + "/" + location, config.cache); err != nil {
		return err
	}

	fmt.Println("Found Pokemon: ")
	for _, names := range api.PokemonEncounters {
		fmt.Println(" - " + names.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("usage: catch <pokemon_name>")
		return nil
	}
	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	randomInt := rand.IntN(100)
	if randomInt < 10 || randomInt > 59 {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)
	}
	return nil
}