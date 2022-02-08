package line

import (
	"experiment-01/utils"
	"github.com/hajimehoshi/ebiten/v2"
	vec "github.com/jairoandre/vector-go"
	"image"
	"image/color"
	"math"
	"math/rand"
	"sync"
)

const (
	Width       = 640
	Height      = 480
	RoutineStep = 1000
)

type Game struct {
	Width     int
	Height    int
	Canvas    *image.RGBA
	Particles []*vec.Vector2d
	Tick      float64
}

func NewGame() *Game {
	var particles [50000]*vec.Vector2d
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

func Hash11(f float64) float64 {
	a := math.Sin(f*3325) * 33253
	b := math.Floor(a)
	return (a - b) - 0.5
}

func (g *Game) Update() error {
	g.Tick += 0.01
	g.Canvas = image.NewRGBA(g.Canvas.Bounds())
	//ctx := gg.NewContextForRGBA(g.Canvas)
	wg := sync.WaitGroup{}
	for i := 0; i < len(g.Particles); i += RoutineStep {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx, p := range g.Particles[i : i+RoutineStep] {
				iF := float64(idx)
				//ctx.DrawCircle(p.X, p.Y, 4)
				//ctx.SetRGBA(1, 1, 1, 0.5)
				//ctx.Fill()
				x, y := p.IntCoords()
				g.Canvas.Set(x, y, color.White)
				p.X += Hash11(iF * g.Tick * 0.02)
				p.Y += Hash11(iF * g.Tick * 0.32)
			}
		}()
	}
	wg.Wait()
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	s.ReplacePixels(g.Canvas.Pix)
	utils.DebugInfo(s)
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
