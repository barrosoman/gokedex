package main

import (
	"bufio"
	"errors"
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

func commandExplore(cfg *config) error {
	if cfg == nil {
		log.Fatal("The config struct was of value nil, can't continue.")
	}


    if len(cfg.parametersStrings) < 1 {
        log.Println("Usage of command is 'explore \"location name\"'.")
        return errors.New("Usage of command is 'explore \"location name\"'.")
    }

    areaName := cfg.parametersStrings[0]
    areaURL := pokeapi.BaseURL + "location-area/" + areaName

    fmt.Println("Exploring", areaName, "...")

    var pokemons []pokeapi.Pokemon
    var err error

	val, ok := cfg.cache.Get(areaURL)


	if !ok {
		log.Println("Getting location info from internet!")
		body, err := pokeapi.GetBodyFromUrl(areaURL)

        if err != nil {
            return errors.New("Couldn't explore " + areaName)
        }

		cfg.cache.Add(areaURL, body)

        locationInfo, err := pokeapi.GetLocationInfoFromBody(body)

        pokemons, err = pokeapi.GetPokemonsFromLocation(locationInfo)
	} else {
		log.Println("Getting location info from cache!")

        locationInfo, err := pokeapi.GetLocationInfoFromBody(val)


        if err != nil {
            return errors.New("Couldn't explore " + areaName)
        }

		pokemons, err = pokeapi.GetPokemonsFromLocation(locationInfo)
	}

    if err != nil {
        return errors.New("Couldn't explore " + areaName)
    }

    fmt.Println("Found Pokemon:")

    for _, v := range pokemons {
        fmt.Println(" - ", v.Name)
    }

    return nil
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
		body, err := pokeapi.GetLocationBodyFromUrl(cfg.nextLocationsURL)

        if err != nil {
            return errors.New("Couldn't get further locations.")
        }

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
		body, err := pokeapi.GetLocationBodyFromUrl(cfg.previousLocationsURL)

        if err != nil {
            return errors.New("Couldn't get previous locations.")
        }

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
			name:        "exit",
			description: "Displays a help message",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays ",
			callback:    commandExplore,
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

        cfg.parametersStrings = words[1:len(words)]

		if !ok {
			fmt.Printf("Command \"%s\" does not exists.\n", commandString)
			continue
		}

        err := command.callback(&cfg)

        if err != nil {
            fmt.Println(err)
        }
	}
}
