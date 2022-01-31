package main

import (
	"experiment-01/line"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := line.NewGame()
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Experiment")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
