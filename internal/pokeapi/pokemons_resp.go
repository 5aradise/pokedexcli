package pokeapi

import (
	"encoding/json"
	"fmt"
)

const areaURL string = "/location-area"

type AreaResp struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
