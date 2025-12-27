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

func GetPokemon(pokemonName string) (pokemonResult Pokemon, err error) {
	url := baseUrl + "pokemon/" + pokemonName

	rawRes, err := http.Get(url)
	if err != nil {
		return pokemonResult, err
	}
	defer rawRes.Body.Close()

	if rawRes.StatusCode != http.StatusOK {
		return pokemonResult, errors.New(pokemonName + " does not exist")
	}

	var res pokemonResponse
	decoder := json.NewDecoder(rawRes.Body)
	if err := decoder.Decode(&res); err != nil {
		return pokemonResult, err
	}

	return formatPokemonResponse(res), nil
}

func formatPokemonResponse(res pokemonResponse) (pokemon Pokemon) {
	pokemon.Name = res.Name
	pokemon.Base_experience = res.Base_experience
	pokemon.Height = res.Height
	pokemon.Weight = res.Weight
	for _, stat := range res.Stats {
		switch stat.Stat.Name {
		case "hp":
			pokemon.Hp = stat.Base_stat
		case "attack":
			pokemon.Attack = stat.Base_stat
		case "defense":
			pokemon.Defense = stat.Base_stat
		case "special-attack":
			pokemon.SpecialAttack = stat.Base_stat
		case "special-defense":
			pokemon.SpecialDefense = stat.Base_stat
		case "speed":
			pokemon.Speed = stat.Base_stat
		}
	}

	for _, pokeType := range res.Types {
		pokemon.Types = append(pokemon.Types, pokeType.Type.Name)
	}

	return pokemon
}
