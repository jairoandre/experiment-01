package main

import (
	"experiment-01/liqui"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := liqui.NewGame()
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Experiment")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
