package main

import (
	"os"

	"github.com/jwengr/saw/cmd"
)

func main() {
	if err := cmd.SawCommand.Execute(); err != nil {
		os.Exit(0)
	}
}
