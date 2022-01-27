package gameoflife

import (
	"experiment-01/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"math/rand"
	"sync"
)

const (
	Width  = 1024
	Height = 768
)

type Game struct {
	Canvas *image.RGBA
	Width  int
	Height int
}

func NewGame() *Game {
	rect := image.Rect(0, 0, Width, Height)
	canvas := image.NewRGBA(rect)
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if rand.Float64() > .5 {
				canvas.Set(x, y, color.White)
			}
		}
	}
	return &Game{
		Width:  Width,
		Height: Height,
		Canvas: canvas,
	}
}

func GetNeighbors(x, y int, canvas *image.RGBA) int {
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

const (
	brushRadius = 5
)

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		for j := -brushRadius; j <= brushRadius; j++ {
			for i := -brushRadius; i <= brushRadius; i++ {
				g.Canvas.Set(mx+i, my+j, color.White)
			}
		}

		return nil
	}
	newCanvas := image.NewRGBA(g.Canvas.Bounds())
	wg := sync.WaitGroup{}
	for y := 0; y < Height; y++ {
		wg.Add(1)
		y := y
		go func() {
			defer wg.Done()
			for x := 0; x < Width; x++ {
				col := g.Canvas.At(x, y)
				r, _, _, _ := col.RGBA()
				alive := r > 100
				num := GetNeighbors(x, y, g.Canvas)
				if alive && (num == 2 || num == 3) {
					newCanvas.Set(x, y, color.White)
				} else if !alive && num == 3 {
					newCanvas.Set(x, y, color.White)
				}
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
