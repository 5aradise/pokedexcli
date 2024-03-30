package main

import "fmt"

func cfgFn(cfg *config, args ...string) error {
	fmt.Printf("Offset: %v\n", cfg.locationOffset)
	fmt.Printf("Limit: %v\n", cfg.locationLimit)
	return nil
}
