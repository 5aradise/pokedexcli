package pokeapi

import (
	"encoding/json"
	"fmt"
)

const areaURL string = "/location-area"

func (c *Client) GetPokemons(area string) ([]struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}, error) {
	areaResp, err := c.getAreaResp(area)
	if err != nil {
		return nil, err
	}
	listPokemons := make([]struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}, 0)
	for _, pokemon := range areaResp.PokemonEncounters {
		listPokemons = append(listPokemons, pokemon.Pokemon)
	}
	return listPokemons, nil
}

func (c *Client) getAreaResp(area string) (resp AreaResp, err error) {
	endpoint := "/" + area
	areaURL := baseURL + areaURL + endpoint

	data, err := c.getResp(areaURL)
	if err != nil {
		if err.Error() == "bad status code: 404" {
			err = fmt.Errorf("wrong area name")
			return
		}
		return
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		resp = AreaResp{}
		return
	}

	return
}
