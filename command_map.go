package main

import (
	"fmt"
	"strconv"
)

func mapFn(cfg *config, args ...string) error {
	if len(args) != 0 {
		offsetStr := args[0]
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return err
		}
		cfg.locationOffset = offset
		if len(args) != 1 {
			limitStr := args[1]
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				return err
			}
			cfg.locationLimit = limit
		}
	}
	return printCfgLocations(cfg)
}

func mapfFn(cfg *config, args ...string) error {
	if len(args) != 0 {
		stepStr := args[0]
		step, err := strconv.Atoi(stepStr)
		if err != nil {
			return err
		}
		cfg.locationLimit = step
	}
	cfg.locationOffset += cfg.locationLimit
	return printCfgLocations(cfg)
}

func mapbFn(cfg *config, args ...string) error {
	if len(args) != 0 {
		stepStr := args[0]
		step, err := strconv.Atoi(stepStr)
		if err != nil {
			return err
		}
		cfg.locationLimit = step
	}
	cfg.locationOffset -= cfg.locationLimit
	return printCfgLocations(cfg)
}

func printCfgLocations(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.locationOffset, cfg.locationLimit)
	if err != nil {
		return err
	}
	locations := locationResp.Results
	printLocations(locations)
	return err
}

func printLocations(locations []struct {
	Name string "json:\"name\""
	URL  string "json:\"url\""
}) {
	locationsMsg := ""
	for _, location := range locations[:len(locations)-1] {
		locationsMsg += fmt.Sprintf("- %s\n", location.Name)
	}
	locationsMsg += fmt.Sprintf("- %s", locations[len(locations)-1].Name)
	fmt.Println(locationsMsg)
}
