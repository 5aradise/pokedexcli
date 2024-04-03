package main

type cliCommand struct {
	usage       string
	description string
	fn          func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			usage:       "help",
			description: "Displays a list of all commands",
			fn:          helpFn,
		},
		"config": {
			usage:       "config",
			description: "Displays the default configuration",
			fn:          cfgFn,
		},
		"map": {
			usage:       "map [offset] [limit]",
			description: "Displays the names of [limit] locations, starting from the [offset], in the Pokemon world",
			fn:          mapFn,
		},
		"mapf": {
			usage:       "mapf [step]",
			description: "Displays the next [step] locations",
			fn:          mapfFn,
		},
		"mapb": {
			usage:       "mapb [step]",
			description: "Displays the previous [step] locations",
			fn:          mapbFn,
		},
		"areas": {
			usage:       "areas {location_name}",
			description: "Displays areas in a {location_name} location",
			fn:          areasFn,
		},
		"explore": {
			usage:       "explore {area_name}",
			description: "Displays pokemons in a {area_name} area",
			fn:          exploreFn,
		},
		"catch": {
			usage:       "catch {pokemon_name}",
			description: "Trys to catches a {pokemon_name} pokemon",
			fn:          catchFn,
		},
		"inspect": {
			usage:       "inspect {pokemon_name}",
			description: "Displays information about a {pokemon_name} pokemon",
			fn:          inspectFn,
		},
		"pokedex": {
			usage:       "pokedex",
			description: "Displays your pokemons",
			fn:          pokedexFn,
		},
		"exit": {
			usage:       "exit",
			description: "Exit the Pokedex",
			fn:          exitFn,
		},
	}
}
