package particles

import (
	vec "github.com/jairoandre/vector-go"
	"github.com/mazznoer/colorgrad"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

var Gradient = colorgrad.Magma()

type Particle struct {
	Dim    int
	Pos    vec.Vector2d
	Weight float64
	Vel    vec.Vector2d
	Acc    vec.Vector2d
	Alpha  *image.Uniform
	Color  color.Color
}

func NewParticle(x, y, dim int) Particle {
	a := uint8(50 + (125 * rand.Float64()))
	col := Gradient.At(rand.Float64())
	alpha := image.NewUniform(color.Alpha{A: a})
	//alpha := image.NewUniform(color.Alpha{A: 0xff})
	return Particle{
		Dim:    dim,
		Weight: 2.0 + rand.Float64()*float64(dim),
		Pos:    vec.NewVec2dFromInt(x, y),
		Alpha:  alpha,
		Color:  col,
	}
}

var zeroPoint = image.Point{}

func (p *Particle) Update() {
	p.Vel = p.Vel.Add(p.Acc)
	if p.Vel.Len() > 4 {
		p.Vel = p.Vel.Normalize().Mul(4.0)
	}
	p.Pos = p.Pos.Add(p.Vel)
	p.Acc = vec.Vector2d{}
}

func (p *Particle) ApplyForce(force vec.Vector2d) {
	p.Acc = p.Acc.Add(force.Div(p.Weight))
}

func (p *Particle) Draw(canvas *image.RGBA) {
	uni := image.NewUniform(p.Color)
	x := int(p.Pos.X)
	y := int(p.Pos.Y)
	offset := p.Dim
	rect := image.Rect(x-offset, y-offset, x+p.Dim, y+p.Dim)
	draw.DrawMask(canvas, rect, uni, zeroPoint, nil, zeroPoint, draw.Over)
}
