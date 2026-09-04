package pokeapi

import (
	"encoding/json"
	"net/http"
)

// Explore Location
func (c *Client) ExploreLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Location{}, err
	}

	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()

	locationResp := Location{}

	err = decoder.Decode(&locationResp)

	if err != nil {
		return Location{}, err
	}

	return locationResp, nil
}
