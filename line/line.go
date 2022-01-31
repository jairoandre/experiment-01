package line

import (
	"experiment-01/utils"
	"github.com/hajimehoshi/ebiten/v2"
	vec "github.com/jairoandre/vector-go"
	"image"
	"image/color"
	"math/rand"
)

const (
	Width  = 640
	Height = 480
)

type Game struct {
	Width     int
	Height    int
	Canvas    *image.RGBA
	Particles []*vec.Vector2d
}

func NewGame() *Game {
	var particles [200]*vec.Vector2d
	for idx, _ := range particles {
		x := rand.Float64() * Width
		y := rand.Float64() * Height
		v := vec.NewVec2d(x, y)
		particles[idx] = &v
	}
	game := &Game{
		Width:     Width,
		Height:    Height,
		Particles: particles[:],
		Canvas:    image.NewRGBA(image.Rect(0, 0, Width, Height)),
	}
	return game
}

func (g *Game) Update() error {
	for _, p := range g.Particles {
		x, y := p.IntCoords()
		g.Canvas.Set(x, y, color.White)
		p.X += rand.Float64() - 0.5
		p.Y += rand.Float64() - 0.5
	}
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	s.ReplacePixels(g.Canvas.Pix)
	utils.DebugInfo(s)
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
