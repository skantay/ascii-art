package main

import (
	"fmt"

	asciiart "github.com/skantay/ascii-art/internal/asciiArt"
)

func main() {
	if err := asciiart.Run(); err != nil {
		fmt.Println(err)
	}
}
