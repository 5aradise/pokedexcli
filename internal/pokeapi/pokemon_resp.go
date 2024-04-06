package pokeapi

import (
	"encoding/json"
	"fmt"
)

const pokeURL string = "/pokemon"

func (c *Client) GetPokemonResp(pokemon string) (pokemonResp PokemonResp, err error) {
	endpoint := "/" + pokemon
	pokemonURL := baseURL + pokeURL + endpoint
	data, err := c.getResp(pokemonURL)
	if err != nil {
		if err.Error() == "bad status code: 404" {
			err = fmt.Errorf("wrong area name")
			return
		}
		return
	}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		pokemonResp = PokemonResp{}
		return
	}
	return
}
