package main

import (
	"experiment-01/fluid"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := fluid.NewGame()
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Experiment")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
