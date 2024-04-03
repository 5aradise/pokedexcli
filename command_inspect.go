package main

import (
	"errors"
	"fmt"
)

func inspectFn(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you didn't enter a pokemon name, write 'help' for more information on using the command")
	}

	pokemonName := args[0]
	pokemon, isInPokedex := cfg.pokedex[pokemonName]
	if !isInPokedex {
		return errors.New("you don't have this pokemon, write 'pokedex' to see your pokemons")
	}

	pokemonInfoMsg := ""
	pokemonInfoMsg += fmt.Sprintf("Name: %s\n", pokemon.Name)
	pokemonInfoMsg += fmt.Sprintf("Height: %d\n", pokemon.Height)
	pokemonInfoMsg += fmt.Sprintf("Weight: %d\n", pokemon.Weight)

	pokemonInfoMsg += "Stats:\n"
	for _, stat := range pokemon.Stats {
		pokemonInfoMsg += fmt.Sprintf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	pokemonInfoMsg += "Types:\n"
	for _, typ := range pokemon.Types[:len(pokemon.Types)-1] {
		pokemonInfoMsg += fmt.Sprintf(" - %s\n", typ.Type.Name)
	}
	pokemonInfoMsg += fmt.Sprintf(" - %s", pokemon.Types[len(pokemon.Types)-1].Type.Name)

	fmt.Println(pokemonInfoMsg)
	return nil
}
