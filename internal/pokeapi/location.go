package pokeapi

import (
	"encoding/json"
	"log"
)

type RespLocationInfo struct {
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
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
    Name string
    URL string
}

func GetLocationInfoFromBody(body []byte) (RespLocationInfo, error) {
    var locationInfo RespLocationInfo

	if err := json.Unmarshal(body, &locationInfo); err != nil {
		log.Println("Couldn't unmarshall json of locations list.")
        return RespLocationInfo{}, err
	}

	return locationInfo, nil
}

func GetPokemonsFromLocation(locationInfo RespLocationInfo) ([]Pokemon, error) {
    pokemonsInLocation := make([]Pokemon, 0)

    for _, v := range locationInfo.PokemonEncounters {
        pokemon := Pokemon{Name: v.Pokemon.Name, URL: v.Pokemon.URL}

        pokemonsInLocation = append(pokemonsInLocation, pokemon)
    }

    return pokemonsInLocation, nil
}
