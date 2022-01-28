package fluid

import (
	"experiment-01/utils"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"sync"
)

const (
	Width  = 640 * 2
	Height = 480 * 2
)

type Game struct {
	Width  int
	Height int
	Canvas *image.RGBA
	Data1  *image.RGBA
	Tick   int
}

func NewGame() *Game {
	rect := image.Rect(0, 0, Width, Height)
	return &Game{
		Width:  Width,
		Height: Height,
		Canvas: image.NewRGBA(rect),
		Data1:  image.NewRGBA(rect),
	}
}

func (g *Game) Update() error {
	g.Tick++
	//if g.Tick%1 != 0 {
	//	return nil
	//}
	newCanvas := image.NewRGBA(g.Canvas.Bounds())
	//newData1 := image.NewRGBA(g.Canvas.Bounds())
	wg := &sync.WaitGroup{}
	for x := 0; x < Width; x++ {
		x := x
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := 0; y < Height; y++ {
				thisCol := g.Canvas.At(x, y)
				if y == 0 && (x%10 == 0 || g.Tick%10 == 0) {
					newCanvas.Set(x, y, color.White)
					//newData1.Set(x, y, color.RGBA{R: uint8(255 * rand.Float64())})
				}
				if thisCol == color.Black {
					continue
				}
				newY := y + 1
				newCanvas.Set(x, newY, thisCol)
			}
		}()
	}
	wg.Wait()
	g.Canvas = newCanvas
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	s.ReplacePixels(g.Canvas.Pix)
	utils.DebugInfo(s)
	utils.DebugInfoMessage(s, fmt.Sprintf("\nTotal pixels: %d", Width*Height))
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
