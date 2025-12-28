package pokeapi

import (
	"encoding/json"
	"errors"
	"goidex/internal/pokecache"
	"net/http"
	"time"
)

var baseUrl string = "https://pokeapi.co/api/v2/"
var cache *pokecache.Cache = pokecache.NewCache(time.Hour)

func GetLocations(url string) (locations []string, next, previous string, err error) {
	if url == "" {
		url = baseUrl + "location-area?offset=0&limit=20"
	}

	resBytes, _, err := get(url)
	if err != nil {
		return nil, "", "", err
	}

	var res locationsResponse
	if err := json.Unmarshal(resBytes, &res); err != nil {
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

	resBytes, statusCode, err := get(url)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New("not a valid location")
	}

	var res locationResponse
	if err := json.Unmarshal(resBytes, &res); err != nil {
		return nil, err
	}

	pokemon = []string{}
	for _, encounter := range res.PokemonEncounters {
		pokemon = append(pokemon, encounter.Pokemon.Name)
	}

	return pokemon, nil
}

func GetPokemon(pokemonName string) (pokemonResult Pokemon, err error) {
	url := baseUrl + "pokemon/" + pokemonName

	resBytes, statusCode, err := get(url)
	if err != nil {
		return pokemonResult, err
	}

	if statusCode != http.StatusOK {
		return pokemonResult, errors.New(pokemonName + " does not exist")
	}

	var res pokemonResponse
	if err := json.Unmarshal(resBytes, &res); err != nil {
		return pokemonResult, err
	}

	return formatPokemonResponse(res), nil
}
