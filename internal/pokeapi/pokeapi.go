package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/5aradise/pokedexcli/internal/pokecache"
)

const baseURL string = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(cacheTTL time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheTTL, 20),
	}
}

func (c *Client) getResp(url string) (data []byte, err error) {
	cacheData, isInCache := c.cache.Get(url)
	if isInCache {
		data = cacheData
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		err = fmt.Errorf("bad status code: %v", resp.StatusCode)
		return
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		data = nil
		return
	}

	c.cache.Add(url, data)
	return
}
