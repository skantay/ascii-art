package main

import (
	"log"

	asciiart "github.com/skantay/ascii-art/internal/asciiArt"
)

func main() {
	if err := asciiart.Run(); err != nil {
		log.Fatal(err)
	}
}
