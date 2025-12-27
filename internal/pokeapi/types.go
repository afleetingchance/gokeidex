package pokeapi

type locationsResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []location
}

type location struct {
	Name string
	Url  string
}

type locationResponse struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon pokemon
}

type pokemon struct {
	Name string
	Url  string
}
