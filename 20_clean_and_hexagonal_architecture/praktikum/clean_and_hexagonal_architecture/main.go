package main

import (
	"log"

	"clean_and_hexagonal_architecture/cmd"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
