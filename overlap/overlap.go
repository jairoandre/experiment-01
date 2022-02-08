package overlap

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	vec "github.com/jairoandre/vector-go"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mazznoer/colorgrad"
	"image"
	"math/rand"
)

const (
	Width  = 640
	Height = 480
)

var grad = colorgrad.Plasma()

type Game struct {
	Width   int
	Height  int
	Circles [100]Circle
	Canvas  *image.RGBA
}

type Circle struct {
	Pos    vec.Vector2d
	Radius float64
	Color  colorful.Color
}

func (c *Circle) Draw(canvas *image.RGBA) {
	ctx := gg.NewContextForRGBA(canvas)
	ctx.Push()
	r, g, b := c.Color.LinearRgb()
	ctx.SetRGBA(r, g, b, 0.6)
	ctx.DrawCircle(c.Pos.X, c.Pos.Y, c.Radius)
	ctx.Fill()
	ctx.Pop()
}

func NewGame() *Game {
	g := &Game{
		Width:  Width,
		Height: Height,
		Canvas: image.NewRGBA(image.Rect(0, 0, Width, Height)),
	}
	for i := 0; i < len(g.Circles); i++ {
		g.Circles[i].Pos.X = rand.Float64() * Width
		g.Circles[i].Pos.Y = rand.Float64() * Height
		g.Circles[i].Radius = rand.Float64() * 50
		g.Circles[i].Color = grad.At(rand.Float64())
	}
	return g
}

func (g *Game) Update() error {
	for i := 0; i < len(g.Circles); i++ {
		g.Circles[i].Pos.X += 0.01
		g.Circles[i].Pos.Y += 0.1
	}
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	for _, c := range g.Circles {
		c.Draw(g.Canvas)
	}
	s.ReplacePixels(g.Canvas.Pix)
	g.Canvas = image.NewRGBA(g.Canvas.Bounds())
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
