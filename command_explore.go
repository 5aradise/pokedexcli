package main

import (
	"errors"
	"fmt"
)

func exploreFn(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you didn't enter an area name, write 'help' for more information on using the command")
	}
	areaName := args[0]
	pokemons, err := cfg.pokeapiClient.GetPokemons(areaName)
	if err != nil {
		return err
	}
	pokemonsMsg := ""
	for _, pokemon := range pokemons[:len(pokemons)-1] {
		pokemonsMsg += fmt.Sprintf("- %s\n", pokemon.Name)
	}
	pokemonsMsg += fmt.Sprintf("- %s", pokemons[len(pokemons)-1].Name)
	fmt.Println("Pokemones in", areaName, ":")
	fmt.Println(pokemonsMsg)
	return nil
}
