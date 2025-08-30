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
	log.Printf("Version: %s\n", version)
	if _, err := app.New(version, nil); err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
