package main

import "fmt"

func pokedexFn(cfg *config, args ...string) error {
	pokedexMsg := "Your Pokedex:\n"
	if len(cfg.pokedex) == 0 {
		pokedexMsg += "Here is empty"
	} else {
		pokemonCount := 0
		for pokemon := range cfg.pokedex {
			if pokemonCount == len(cfg.pokedex)-1 {
				pokedexMsg += fmt.Sprintf(" - %s", pokemon)
				continue
			}
			pokedexMsg += fmt.Sprintf(" - %s\n", pokemon)
			pokemonCount++
		}
	}
	fmt.Println(pokedexMsg)
	return nil
}
