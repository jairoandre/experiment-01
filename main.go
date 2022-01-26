package main

import (
	"experiment-01/liquid"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := liquid.NewGame()

	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Experiment")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
