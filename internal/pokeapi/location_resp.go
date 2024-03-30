package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(offset, limit int) (LocationsResp, error) {
	endpoint := fmt.Sprintf("/location?offset=%v&limit=%v", offset, limit)
	listLocsURL := baseURL + endpoint

	req, err := http.NewRequest("GET", listLocsURL, nil)
	if err != nil {
		return LocationsResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationsResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResp{}, err
	}

	locationsResp := LocationsResp{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationsResp{}, err
	}

	return locationsResp, nil
}
