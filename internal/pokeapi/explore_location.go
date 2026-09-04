package pokeapi

//Explore Location
func (c *Client) ExploreLocation(location string) []string {
	url := baseURL + "/location-area/" + location
	return []string{}
}
