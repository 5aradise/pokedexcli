package pokeapi

import (
	"encoding/json"
	"fmt"
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
	}, 0, limit)
	lastLoc := firstLoc + limit - 1
	startInChunk := firstLoc % locationChunkSize
	endInChunk := lastLoc % locationChunkSize
	firstChunk := firstLoc / locationChunkSize
	lastChunk := lastLoc / locationChunkSize
	chunkOfLocs, err := c.getLocationsChunkResp(firstChunk)
	if err != nil {
		return nil, err
	}
	if firstChunk == lastChunk {
		listLocations = append(listLocations, chunkOfLocs.Results[startInChunk:endInChunk+1]...)
		return listLocations, nil
	}
	listLocations = append(listLocations, chunkOfLocs.Results[startInChunk:]...)
	for i := firstChunk + 1; i < lastChunk; i++ {
		chunkOfLocs, err := c.getLocationsChunkResp(i)
		if err != nil {
			return nil, err
		}
		listLocations = append(listLocations, chunkOfLocs.Results...)
	}
	chunkOfLocs, err = c.getLocationsChunkResp(lastChunk)
	if err != nil {
		return nil, err
	}
	listLocations = append(listLocations, chunkOfLocs.Results[:endInChunk]...)
	return listLocations, nil
}

func (c *Client) getLocationsChunkResp(chunkNum int) (resp LocationsResp, err error) {
	endpoint := fmt.Sprintf("?offset=%v&limit=%v", chunkNum*locationChunkSize, locationChunkSize)
	locsURL := baseURL + locationURL + endpoint

	data, err := c.getResp(locsURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		resp = LocationsResp{}
		return
	}

	return
}
