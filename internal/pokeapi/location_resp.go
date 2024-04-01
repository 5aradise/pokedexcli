package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const locationURL string = "/location"
const locationChunkSize = 20

type LocationsResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(firstLoc, limit int) ([]struct {
	Name string "json:\"name\""
	URL  string "json:\"url\""
}, error) {
	listLocations := make([]struct {
		Name string "json:\"name\""
		URL  string "json:\"url\""
	}, 0)
	lastLoc := firstLoc + limit - 1
	startInChunk := firstLoc % locationChunkSize
	endInChunk := lastLoc % locationChunkSize
	firstChunk := firstLoc / locationChunkSize
	lastChunk := lastLoc / locationChunkSize
	chunkOfLocs, err := c.getLocationsChunk(firstChunk)
	if err != nil {
		return nil, err
	}
	if firstChunk == lastChunk {
		listLocations = append(listLocations, chunkOfLocs.Results[startInChunk:endInChunk+1]...)
		return listLocations, nil
	}
	listLocations = append(listLocations, chunkOfLocs.Results[startInChunk:]...)
	for i := firstChunk + 1; i < lastChunk; i++ {
		chunkOfLocs, err := c.getLocationsChunk(i)
		if err != nil {
			return nil, err
		}
		listLocations = append(listLocations, chunkOfLocs.Results...)
	}
	chunkOfLocs, err = c.getLocationsChunk(lastChunk)
	if err != nil {
		return nil, err
	}
	listLocations = append(listLocations, chunkOfLocs.Results[:endInChunk]...)
	return listLocations, nil
}

func (c *Client) getLocationsChunk(chunkNum int) (LocationsResp, error) {
	endpoint := fmt.Sprintf("?offset=%v&limit=%v", chunkNum*locationChunkSize, locationChunkSize)
	locsURL := baseURL + locationURL + endpoint

	cacheData, isInCache := c.cache.Get(locsURL)
	if isInCache {
		locationsResp := LocationsResp{}
		err := json.Unmarshal(cacheData, &locationsResp)
		if err != nil {
			return LocationsResp{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", locsURL, nil)
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

	c.cache.Add(locsURL, data)
	return locationsResp, nil
}
