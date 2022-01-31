package main

import (
	"experiment-01/fluid"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := fluid.NewGame()
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Experiment")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
