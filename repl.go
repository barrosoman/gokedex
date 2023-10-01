package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokedex-go/internal/pokeapi"
	"pokedex-go/internal/pokecache"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func commandHelp(cfg *config) error {
	fmt.Println("Run help command")
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Run exit command")
	return nil
}

func commandMap(cfg *config) error {
	if cfg == nil {
		log.Fatal("The config struct was of value nil, can't continue.")
	}

	var locationList pokeapi.RespLocationsList

	val, ok := cfg.cache.Get(cfg.nextLocationsURL)

	if !ok {
		body := pokeapi.GetLocationBodyFromUrl(cfg.nextLocationsURL)

		cfg.cache.Add(cfg.nextLocationsURL, body)

		locationList = pokeapi.GetLocationsFromBody(body)
	} else {
		log.Println("Getting location list from cache!")
		locationList = pokeapi.GetLocationsFromBody(val)
	}

	cfg.setupLocationURLs(locationList)

	for _, location := range locationList.Locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg == nil {
		log.Fatal("The config struct was of value \"nil\", can't continue.")
	}
	var locationList pokeapi.RespLocationsList

	val, ok := cfg.cache.Get(cfg.previousLocationsURL)

	if !ok {
		body := pokeapi.GetLocationBodyFromUrl(cfg.previousLocationsURL)

		cfg.cache.Add(cfg.previousLocationsURL, body)

		locationList = pokeapi.GetLocationsFromBody(body)
	} else {
		log.Println("Getting location list from cache!")
		locationList = pokeapi.GetLocationsFromBody(val)
	}

	cfg.setupLocationURLs(locationList)

	for _, location := range locationList.Locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandPrintCache(cfg *config) error {
    for key := range cfg.cache.Entries() {
        fmt.Println(key)
    }
	return nil
}

func setupCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 map names, (next 20 if used again)",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map names, (previous 20 if used again)",
			callback:    commandMapb,
		},
		"printcache": {
			name:        "printcache",
			description: "Print all urls cached.",
			callback:    commandPrintCache,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := setupCommands()
	cfg := config{cache: pokecache.NewCache(time.Minute * 5)}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := strings.Split(scanner.Text(), " ")

		commandString := words[0]
		command, ok := commands[commandString]

		if !ok {
			fmt.Printf("Command \"%s\" does not exists.\n", commandString)
			continue
		}

		command.callback(&cfg)
	}
}