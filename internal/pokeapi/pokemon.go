package pokeapi

import (
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon string) (bool, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		mon := Pokemon{}
		err := json.Unmarshal(val, &mon)
		if err != nil {
			return false, err
		}
		return attemptCatch(mon), nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return false, err
	}

	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()

	var mon Pokemon

	err = decoder.Decode(&mon)

	if err != nil {
		return false, err
	}

	return attemptCatch(mon), nil

}

func attemptCatch(mon Pokemon) bool {
	chance := math.Max(0.02, (1.0 - float64(mon.BaseExperience/635)))

	hitNum := rand.Float64()

	if chance < hitNum {
		return false
	}

	return true
}
