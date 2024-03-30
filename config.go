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
