package main

import (
	"os"

	"github.com/and-rad/ocis-books/pkg/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
