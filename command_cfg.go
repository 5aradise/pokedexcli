package main

import "fmt"

func cfgFn(cfg *config, args ...string) error {
	fmt.Printf("Level: %v\n", cfg.playerLevel)
	fmt.Printf("Location offset: %v\n", cfg.locationOffset)
	fmt.Printf("Location limit: %v\n", cfg.locationLimit)
	return nil
}
