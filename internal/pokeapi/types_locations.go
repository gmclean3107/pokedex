package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	Id                     int    `json: "id"`
	Name                   string `json: "name"`
	Game_index             int    `json: "game_index"`
	Encounter_method_rates []struct {
		Name string `json: "name"`
		Url  string `json: "url"`
	} `json: "encounter_method_rates`
	Location struct {
		Name string `json: "name"`
		Url  string `json: "url"`
	} `json: "location"`
	Names []struct {
		Name     string `json: "name"`
		Language struct {
			Name string `json: "name"`
			Url  string `json: "url"`
		} `json: "language"`
	} `json: "names"`
	Pokemon_encounters []struct {
		Pokemon struct {
			Name            string `json: "name"`
			Url             string `json: "url"`
			Version_details []struct {
				Version struct {
					Name string `json: "name"`
					Url  string `json: "url"`
				} `json: "version"`
				Max_chance        int `json: "max_chance"`
				Encounter_details []struct {
					Min_level int `json: "min_level"`
					Max_level int `json: "max_level"`
					Chance    int `json: "chance"`
					Method    struct {
						Name string `json: "name"`
						Url  string `json: "url"`
					} `json: "method"`
				} `json: "encounter_details"`
			} `json: "version_details"`
		} `json: "pokemon"`
	} `json: "pokemon_encounters`
}
