package selfportrait

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mazznoer/colorgrad"
	"image"
	"image/jpeg"
	"os"
)

const (
	Width    = 640
	Height   = 480
	GradNum  = 256
	ColorDiv = float32(0xff << 8)
)

type Game struct {
	Width    int
	Height   int
	Time     int
	Image    *ebiten.Image
	Shader   *ebiten.Shader
	Gradient [GradNum * 3]float32
}

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func GradientToFloatArray(gradient colorgrad.Gradient) [GradNum * 3]float32 {
	var result [GradNum * 3]float32
	for idx, color := range gradient.Colors(GradNum) {
		r, g, b, _ := color.RGBA()
		offset := idx * 3
		result[offset] = float32(r) / ColorDiv
		result[offset+1] = float32(g) / ColorDiv
		result[offset+2] = float32(b) / ColorDiv
	}
	return result
}

func NewGame() *Game {
	file, err := os.Open("selfportrait/me.jpg")
	if err != nil {
		panic("error reading image file")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	shader, err := ebiten.NewShader(shaderCode)
	if err != nil {
		panic("error loading shader")
	}
	g := &Game{
		Width:    bounds.Dx(),
		Height:   bounds.Dy(),
		Image:    ebiten.NewImageFromImage(img),
		Shader:   shader,
		Gradient: GradientToFloatArray(colorgrad.Plasma()),
	}
	return g
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
		"Gradient":   g.Gradient,
	}
	op.Images[0] = g.Image
	s.DrawRectShader(g.Width, g.Height, g.Shader, op)
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
