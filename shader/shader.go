package shader

import (
	"experiment-01/utils"
	"github.com/hajimehoshi/ebiten/v2"
	vec "github.com/jairoandre/vector-go"
	"github.com/mazznoer/colorgrad"
	"image"
)

const (
	Width  = 1024
	Height = 768
)

type Game struct {
	Canvas *ebiten.Image
	Shader *ebiten.Shader
	Width  int
	Height int
	Time   int
}

var gradient = colorgrad.BrBG()

func NewGame() *Game {
	shader, err := ebiten.NewShader(fireWorks)
	if err != nil {
		panic("error loading shader")
	}
	center := vec.NewVec2dFromInt(Width/2, Height/2)
	centerLen := center.Len()
	canvas := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			thisPos := vec.NewVec2dFromInt(x, y)
			canvas.Set(x, y, gradient.At(center.Sub(thisPos).Len()/centerLen))
		}
	}
	return &Game{
		Shader: shader,
		Canvas: ebiten.NewImageFromImage(canvas),
		Width:  1024,
		Height: 768,
	}

}

func (g *Game) Update() error {
	g.Time++
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	cx, cy := ebiten.CursorPosition()
	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]interface{}{
		"Time":       float32(g.Time) / 60,
		"Cursor":     []float32{float32(cx), float32(cy)},
		"ScreenSize": []float32{float32(g.Width), float32(g.Height)},
	}
	//op.Images[0] = g.Canvas
	s.DrawRectShader(g.Width, g.Height, g.Shader, op)

	utils.DebugInfo(s)
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
