package main

import (
	"os"
)

func commandExit(config *commandConfig, args ...string) error {
	os.Exit(0)
	return nil
}
