package pokeapi

import (
	"encoding/json"
	"fmt"
)

type LocationResp struct {
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
}

func (c *Client) GetAreas(location string) ([]struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}, error) {
	locationResp, err := c.getLocationResp(location)
	if err != nil {
		return nil, err
	}
	return locationResp.Areas, nil
}

func (c *Client) getLocationResp(location string) (resp LocationResp, err error) {
	endpoint := "/" + location
	locURL := baseURL + locationURL + endpoint

	data, err := c.getResp(locURL)
	if err != nil {
		if err.Error() == "bad status code: 404" {
			err = fmt.Errorf("wrong location name")
			return
		}
		return
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		resp = LocationResp{}
		return
	}

	return
}
