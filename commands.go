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
			description: "Displays the configuration",
			fn:          cfgFn,
		},
		"map": {
			usage:       "map <offset> <limit>",
			description: "Displays the names of {limit} locations, starting from the {offset}, in the Pokemon world",
			fn:          mapFn,
		},
		"mapf": {
			usage:       "mapf <step>",
			description: "Displays the next {step} locations",
			fn:          mapfFn,
		},
		"mapb": {
			usage:       "mapb <step>",
			description: "Displays the previous {step} locations",
			fn:          mapbFn,
		},
		"exit": {
			usage:       "exit",
			description: "Exit the Pokedex",
			fn:          exitFn,
		},
	}
}
