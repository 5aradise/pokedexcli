package main

import "os"

func exitFn(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
