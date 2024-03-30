package main

import (
	"fmt"
	"strconv"
)

func mapFn(cfg *config, args ...string) error {
	if len(args) != 0 {
		offset, limit, err := parseMapArgs(args)
		if err != nil {
			return err
		}
		cfg.setLocationOffset(offset)
		cfg.setLocationLimit(limit)
	}
	return printCfgLocations(cfg)
}

func mapfFn(cfg *config, args ...string) error {
	if len(args) != 0 {
		limit, _, err := parseMapArgs(args)
		if err != nil {
			return err
		}
		cfg.setLocationLimit(limit)
	}
	cfg.setLocationOffset(cfg.locationOffset + cfg.locationLimit)
	return printCfgLocations(cfg)
}

func mapbFn(cfg *config, args ...string) error {
	if len(args) != 0 {
		limit, _, err := parseMapArgs(args)
		if err != nil {
			return err
		}
		cfg.setLocationLimit(limit)
	}
	cfg.setLocationOffset(cfg.locationOffset - cfg.locationLimit)
	return printCfgLocations(cfg)
}

func parseMapArgs(args []string) (arg1, arg2 int, err error) {
	arg1Str := args[0]
	arg1, err = strconv.Atoi(arg1Str)
	if err != nil {
		return
	}
	if len(args) > 1 {
		arg2Str := args[1]
		arg2, err = strconv.Atoi(arg2Str)
		if err != nil {
			return
		}
	}
	return
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
	fmt.Println("Locations:")
	fmt.Println(locationsMsg)
}
