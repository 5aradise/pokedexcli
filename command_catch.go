package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func catchFn(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you didn't enter a pokemon name, write 'help' for more information on using the command")
	}

	pokemonName := args[0]
	if _, isHavePokemon := cfg.pokedex[pokemonName]; isHavePokemon {
		return errors.New("you already caught this pokemon")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemonResp(pokemonName)
	if err != nil {
		return err
	}

	pokemonExp := pokemon.BaseExperience
	fmt.Println("Your level:", cfg.playerLevel)
	fmt.Println(strings.Title(pokemonName), "level:", pokemonExp)

	fmt.Println("Throwing a Pokeball at", pokemonName, "...")
	time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
	poekmonAttack := rand.Intn(pokemonExp)
	if cfg.playerLevel < poekmonAttack {
		fmt.Println("You missed!")
		return nil
	}

	cfg.pokedex[pokemonName] = pokemon
	fmt.Println("You caught a", pokemon.Name)

	gainedLvl := int(math.Round(float64(pokemonExp) / 10))
	cfg.playerLevel += gainedLvl
	msglvl := "levels"
	if gainedLvl == 1 {
		msglvl = "level"
	}
	fmt.Println("You`ve gained", gainedLvl, msglvl)
	return nil
}
