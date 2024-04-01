package pokeapi

import (
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
