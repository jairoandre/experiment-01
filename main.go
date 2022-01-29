package main

import (
	"experiment-01/particles"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := particles.NewGame()
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Experiment")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
