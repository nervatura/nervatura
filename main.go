package main

import (
	"log"
	"os"

	"github.com/nervatura/nervatura/v6/pkg/app"
)

var (
	version = "dev"
)

func main() {
	if _, err := app.New(version, nil); err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
