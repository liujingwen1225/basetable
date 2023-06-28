package main

import (
	"basetable.com/internal/basetable"
	_ "go.uber.org/automaxprocs"
	"os"
)

func main() {
	command := basetable.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
