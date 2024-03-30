package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()
		parsed := parseInput(text)
		if len(parsed) == 0 {
			continue
		}

		commandHead := parsed[0]
		commandArgs := parsed[1:]
		command, isCommand := commands[commandHead]
		if !isCommand {
			fmt.Println("Wrong command!")
			fmt.Println("To check the list of all commands, enter 'help'")
			continue
		}

		err := command.fn(cfg, commandArgs...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func parseInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
