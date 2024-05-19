package main

import (
	"os"
)

func commandExit(config *commandConfig) error {
	os.Exit(0)
	return nil
}
