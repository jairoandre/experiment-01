package liqui

import (
	"experiment-01/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mazznoer/colorgrad"
	"image"
	"sync"
)

const (
	Width  = 640
	Height = 480
)

type CellData struct {
	image.RGBA
	Level [][]float64
}

func (c *CellData) SetLevel(x, y int, level float64) {
	c.Level[y][x] = level
}

func (c *CellData) LevelAt(x, y int) float64 {
	return c.Level[y][x]
}

func NewCellData(rect image.Rectangle) *CellData {
	pt := rect.Max
	level := make([][]float64, 0)
	for y := 0; y < pt.Y; y++ {
		row := make([]float64, pt.X)
		level = append(level, row)
	}
	return &CellData{
		RGBA:  *image.NewRGBA(rect),
		Level: level,
	}
}

type Game struct {
	Canvas *CellData
	Width  int
	Height int
}

var gradient = colorgrad.Inferno()

func NewGame() *Game {
	rect := image.Rect(0, 0, Width, Height)
	canvas := NewCellData(rect)
	//for y := 0; y < Height; y++ {
	//	for x := 0; x < Width; x++ {
	//		level := rand.Float64()
	//		canvas.Set(x, y, gradient.At(level))
	//		canvas.SetLevel(x, y, level)
	//	}
	//}
	return &Game{
		Width:  Width,
		Height: Height,
		Canvas: canvas,
	}
}

func GetNeighbors(x, y int, canvas *CellData) int {
	num := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if j == 0 && i == 0 {
				continue
			}
			nX := x + i
			nY := y + j
			col := canvas.At(nX, nY)
			r, _, _, _ := col.RGBA()
			if r > 100 {
				num++
			}
		}
	}
	return num
}

func (g *Game) Update() error {
	newCanvas := NewCellData(g.Canvas.Bounds())
	wg := sync.WaitGroup{}
	for y := 0; y < Height; y++ {
		wg.Add(1)
		y := y
		go func() {
			defer wg.Done()
			for x := 0; x < Width; x++ {
				newLevel := g.Canvas.LevelAt(x, y) - 0.1
				if newLevel < 0 {
					newLevel = 1.0
				}
				newCanvas.Set(x, y, gradient.At(newLevel))
				newCanvas.SetLevel(x, y, newLevel)
			}
		}()
	}
	wg.Wait()
	g.Canvas = newCanvas
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.Canvas.Pix)
	utils.DebugInfo(screen)
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
