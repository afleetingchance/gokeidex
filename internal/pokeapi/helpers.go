package pokeapi

import (
	"io"
	"net/http"
)

func get(url string) (res []byte, status int, err error) {
	resBytes, ok := cache.Get(url)

	if ok {
		return resBytes, http.StatusOK, nil
	}

	rawRes, err := http.Get(url)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rawRes.Body.Close()

	resBytes, err = io.ReadAll(rawRes.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	cache.Add(url, resBytes)

	return resBytes, http.StatusOK, nil
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
