package liquid

import (
	"experiment-01/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"sync"
)

const (
	Width  = 640
	Height = 480
)

type Game struct {
	Cells       []*Cell
	TotalCells  int
	StepRoutine int
	Canvas      *image.RGBA
	Width       int
	Height      int
}

func NewGame() *Game {
	cells := make([]*Cell, 0)
	canvas := image.NewRGBA(image.Rect(0, 0, 640, 480))
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			cells = append(cells, NewCell(x, y))
		}
	}
	return &Game{
		Canvas:      canvas,
		Cells:       cells,
		TotalCells:  len(cells),
		StepRoutine: len(cells) / 1000,
		Width:       Width,
		Height:      Height,
	}
}

func (g *Game) Update() error {
	g.Canvas = image.NewRGBA(g.Canvas.Bounds())
	wg := sync.WaitGroup{}
	for step := 0; step < g.TotalCells; step += g.StepRoutine {
		wg.Add(1)
		start := step
		end := step + g.StepRoutine
		if end > g.TotalCells {
			end = g.TotalCells
		}
		go func() {
			defer wg.Done()
			for _, cell := range g.Cells[start:end] {
				cell.Draw(g.Canvas)
			}
		}()
	}
	wg.Wait()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.Canvas.Pix)
	utils.DebugInfo(screen)
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return 640, 480
}
