package main

import (
	"os"
)

func commandExit(config *commandConfig, s *string) error {
	os.Exit(0)
	return nil
}
