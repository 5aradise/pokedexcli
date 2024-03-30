package main

import "fmt"

func helpFn() error {
	helpMsg := ""
	for name, comand := range getCommands() {
		helpMsg += fmt.Sprintf("%s: %s\n", name, comand.description)
	}
	fmt.Print(helpMsg)
	return nil
}
