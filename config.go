package main

import "github.com/5aradise/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient  pokeapi.Client
	locationOffset int
	locationLimit  int
}

func NewConfig() config {
	return config{
		pokeapiClient:  pokeapi.NewClient(),
		locationOffset: 0,
		locationLimit:  20,
	}
}

func (cfg *config) setLocationOffset(offset int) {
	locationCount := 1036
	if offset < 0 {
		cfg.locationOffset = 0
		return
	}
	if offset >= locationCount {
		return
	}
	cfg.locationOffset = offset
}

func (cfg *config) setLocationLimit(limit int) {
	if limit < 0 {
		cfg.locationLimit = 1
		return
	}
	if limit == 0 {
		return
	}
	cfg.locationLimit = limit
}
