package main

import (
	"os"

	"github.com/philsabourin/go-toy-robot/game"
)

func main() {
	game.StartAndReadFrom(os.Stdin)
}
