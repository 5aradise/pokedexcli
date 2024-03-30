package main

import "os"

func exitFn() error {
	os.Exit(0)
	return nil
}
