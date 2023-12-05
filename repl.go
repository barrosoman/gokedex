package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"pokedex-go/internal/pokeapi"
	"strconv"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, pars ...string) error
}

func commandExplore(cfg *config, pars ...string) error {
	if cfg == nil {
		log.Fatal("The config struct was of value nil, can't continue.")
	}

    if len(pars) < 1 {
        log.Println("Usage of command is 'explore \"location name\"'.")
        return errors.New("Usage of command is 'explore \"location name\"'.")
    }

    areaName := pars[0]

    fmt.Println("Exploring", areaName, "...")

    locationInfo, err := cfg.client.GetLocation(areaName)

    if err != nil {
        return err
    }

    if len(locationInfo.PokemonEncounters) == 0 {
        log.Println("No pokemons found in this area.")
        return errors.New("No pokemons found in this area.")
    }

    fmt.Println("Found Pokemon:")
    for _, v := range locationInfo.PokemonEncounters {
        fmt.Println(" - ", v.Pokemon.Name)
    }

    return nil
}

func commandPokedex(cfg *config, pars ...string) error {
    fmt.Println("Your pokedex:")

    for name, _ := range cfg.pokemons {
        fmt.Printf(" -%s\n", name)
    }


    return nil
}

func commandCatch(cfg *config, pars ...string) error {
	if cfg == nil {
		log.Fatal("The config struct was of value nil, can't continue.")
	}

    if len(pars) < 1 {
        log.Println("Usage of command is 'explore \"location name\"'.")
        return errors.New("Usage of command is 'explore \"location name\"'.")
    }

    pokemonName := pars[0]

    if _, ok := cfg.pokemons[pokemonName]; ok {
        fmt.Println(pokemonName + " already catched!")
        return nil
    }

    fmt.Println("Catching", pokemonName, "...")

    pokemonInfo, err := cfg.client.GetPokemon(pokemonName)

    if err != nil {
        return errors.New("Couldn't find pokemon " + pokemonName)
    }

    randomValue := rand.Float64()


    if randomValue > (1.0 - (float64(pokemonInfo.BaseExperience) / 700.0)) {
        fmt.Println("Catched " + pokemonInfo.Name + "!")

        cfg.pokemons[pokemonName] = pokemonInfo
    } else {
        fmt.Println(pokemonInfo.Name + " escaped!")
    }

    return nil
}

func printPokemonInfo(pokemon pokeapi.Pokemon) {
    fmt.Println("Name: " + pokemon.Name)
    fmt.Println("Height: " + strconv.Itoa(pokemon.Height))
    fmt.Println("Weight: " + strconv.Itoa(pokemon.Weight))
    fmt.Println("Stats: ")
    for _, v := range pokemon.Stats {
        fmt.Println("  -" + v.Stat.Name + ": " + strconv.Itoa(v.BaseStat))
    }
    fmt.Println("Types: ")
    for _, v := range pokemon.Types {
        fmt.Println("  -" + v.Type.Name)
    }
}

func commandInspect(cfg *config, pars ...string) error {
	if cfg == nil {
		log.Fatal("The config struct was of value nil, can't continue.")
	}

    if len(pars) < 1 {
        log.Println("Usage of command is 'explore \"location name\"'.")
        return errors.New("Usage of command is 'explore \"location name\"'.")
    }

    pokemonName := pars[0]

    if v, ok := cfg.pokemons[pokemonName]; !ok {
        fmt.Println("you have not caught that pokemon.")
    } else {
        printPokemonInfo(v)
    }

	return nil
}

func commandHelp(cfg *config, pars ...string) error {
	fmt.Println("Run help command")
	return nil
}

func commandExit(cfg *config, pars ...string) error {
	fmt.Println("Run exit command")
	return nil
}

func commandMap(cfg *config, pars ...string) error {
	if cfg == nil {
		log.Fatal("The config struct was of value nil, can't continue.")
	}

	var locationList pokeapi.RespLocationsList

	val, ok := cfg.client.Cache.Get(cfg.nextLocationsURL)

	if !ok {
		body, err := cfg.client.GetLocationBodyFromUrl(cfg.nextLocationsURL)

        if err != nil {
            return errors.New("Couldn't get further locations.")
        }

		cfg.client.Cache.Add(cfg.nextLocationsURL, body)

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

func commandMapb(cfg *config, pars ...string) error {
	if cfg == nil {
		log.Fatal("The config struct was of value \"nil\", can't continue.")
	}
	var locationList pokeapi.RespLocationsList

	val, ok := cfg.client.Cache.Get(cfg.previousLocationsURL)

	if !ok {
		body, err := cfg.client.GetLocationBodyFromUrl(cfg.previousLocationsURL)

        if err != nil {
            return errors.New("Couldn't get previous locations.")
        }

		cfg.client.Cache.Add(cfg.previousLocationsURL, body)

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

func commandPrintCache(cfg *config, pars ...string) error {
    for key := range cfg.client.Cache.Entries() {
        fmt.Println(key)
    }
	return nil
}

func setupCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Displays a help message.",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemons found in an area.",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows information about a catched pokemon.",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon.",
			callback:    commandCatch,
		},
		"map": {
			name:        "map",
			description: "Displays 20 map names, (next 20 if used again).",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map names, (previous 20 if used again).",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print all pokemons on your pokedex.",
			callback:    commandPokedex,
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
    cfg := config{client: pokeapi.NewClient(time.Minute * 5, time.Minute * 5), pokemons: make(map[string]pokeapi.Pokemon, 0)}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := strings.Fields(scanner.Text())

        if (len(words) < 1) {
            continue
        }

		commandString := words[0]
		command, ok := commands[commandString]

		if !ok {
			fmt.Printf("Command \"%s\" does not exists.\n", commandString)
			continue
		}

        err := command.callback(&cfg, words[1:]...)

        if err != nil {
            fmt.Println(err)
        }
	}
}
