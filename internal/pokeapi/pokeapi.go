package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

var baseUrl string = "https://pokeapi.co/api/v2/"

func GetLocations(url string) (locations []string, next, previous string, err error) {
	if url == "" {
		url = baseUrl + "location-area"
	}

	rawRes, err := http.Get(url)
	if err != nil {
		return nil, "", "", err
	}
	defer rawRes.Body.Close()

	var res locationsResponse
	decoder := json.NewDecoder(rawRes.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, "", "", err
	}

	locations = []string{}
	for _, result := range res.Results {
		locations = append(locations, result.Name)
	}

	return locations, res.Next, res.Previous, nil
}

func GetPokemonFromLocation(location string) (pokemon []string, err error) {
	url := baseUrl + "location-area/" + location

	rawRes, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rawRes.Body.Close()

	if rawRes.StatusCode != http.StatusOK {
		return nil, errors.New("not a valid location")
	}

	var res locationResponse
	decoder := json.NewDecoder(rawRes.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}

	pokemon = []string{}
	for _, encounter := range res.PokemonEncounters {
		pokemon = append(pokemon, encounter.Pokemon.Name)
	}

	return pokemon, nil
}
