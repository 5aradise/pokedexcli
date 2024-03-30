package main

type cliCommand struct {
	description string
	fn          func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			description: "Displays a list of all commands",
			fn:          helpFn,
		},
		"exit": {
			description: "Exit the Pokedex",
			fn:          exitFn,
		},
	}
}
