package main

import (
	"errors"
	"fmt"
)

func areasFn(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you didn't enter a location name, write 'help' for more information on using the command")
	}
	locationName := args[0]
	areas, err := cfg.pokeapiClient.GetAreas(locationName)
	if err != nil {
		return err
	}

	if len(areas) == 0 {
		fmt.Println("There are no areas in", locationName)
		return nil
	}
	
	areasMsg := ""
	for _, area := range areas[:len(areas)-1] {
		areasMsg += fmt.Sprintf("- %s\n", area.Name)
	}
	areasMsg += fmt.Sprintf("- %s", areas[len(areas)-1].Name)
	fmt.Println("Areas in", locationName, ":")
	fmt.Println(areasMsg)
	return nil
}
