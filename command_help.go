package main

import "fmt"

func helpFn(cfg *config, args ...string) error {
	helpMsg := ""
	for _, comand := range getCommands() {
		helpMsg += fmt.Sprintf("%s - %s\n", comand.usage, comand.description)
	}
	fmt.Print(helpMsg)
	return nil
}
