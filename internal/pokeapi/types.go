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
	Pokemon Pokemon
}

type Pokemon struct {
	Name            string
	Base_experience int
	Height          int
	Weight          int
	Hp              int
	Attack          int
	Defense         int
	SpecialAttack   int
	SpecialDefense  int
	Speed           int
	Types           []string
}

type pokemonResponse struct {
	Name            string
	Base_experience int
	Height          int
	Weight          int
	Stats           []struct {
		Base_stat int
		Stat      struct {
			Name string
		}
	}
	Types []struct {
		Type struct {
			Name string
		}
	}
}
