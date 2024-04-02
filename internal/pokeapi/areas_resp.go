package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Client) getLocationResp(location string) (LocationResp, error) {
	locationResp := LocationResp{}

	endpoint := "/" + location
	locURL := baseURL + locationURL + endpoint

	cacheData, isInCache := c.cache.Get(locURL)
	if isInCache {
		err := json.Unmarshal(cacheData, &locationResp)
		if err != nil {
			return LocationResp{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", locURL, nil)
	if err != nil {
		return LocationResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return LocationResp{}, fmt.Errorf("wrong location name")
	}

	if resp.StatusCode > 399 {
		return LocationResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResp{}, err
	}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationResp{}, err
	}

	c.cache.Add(locURL, data)
	return locationResp, nil
}
