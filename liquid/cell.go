package liquid

import (
	vec "github.com/jairoandre/vector-go"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

type Cell struct {
	Level uint8
	Pos   vec.Vector2d
	Vel   vec.Vector2d
	Up    *Cell
	Down  *Cell
	Left  *Cell
	Right *Cell
}

func NewCell(x, y int) *Cell {
	pos := vec.NewVec2dFromInt(x, y)
	var level uint8
	if rand.Float64() < 0.01 {
		level = 255
	}
	return &Cell{
		Level: level,
		Pos:   pos,
	}
}

func (c *Cell) Update() {

}

func (c *Cell) Draw(canvas *image.RGBA) {
	uniform := image.NewUniform(color.RGBA{R: c.Level, G: c.Level, B: c.Level, A: 0xff})
	x, y := c.Pos.IntCoords()
	draw.Draw(canvas, image.Rect(x, y, x+1, y+1), uniform, image.Point{}, draw.Over)
}
