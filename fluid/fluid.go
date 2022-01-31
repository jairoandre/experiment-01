package fluid

import (
	"experiment-01/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mazznoer/colorgrad"
	"image"
	"math"
	"sync"
)

const (
	N                     = 400
	NFloat                = float64(N)
	TotalPixels           = N * N
	LinearSolveIterations = 4
)

type Game struct {
	Width  int
	Height int
	Fluid  *Fluid
}

type Fluid struct {
	Dt      float64
	Diff    float64
	Visc    float64
	S       [TotalPixels]float64
	Density [TotalPixels]float64
	Vx      [TotalPixels]float64
	Vy      [TotalPixels]float64
	Vx0     [TotalPixels]float64
	Vy0     [TotalPixels]float64
}

func IX(x, y int) int {
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	if x > N-1 {
		x = N - 1
	}
	if y > N-1 {
		y = N - 1
	}
	return x + y*N
}

func NewFluid(dt, diff, visc float64) *Fluid {
	return &Fluid{
		Dt:   dt,
		Diff: diff,
		Visc: visc,
	}
}

var gradient = colorgrad.Inferno()

func (f *Fluid) Draw() *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, N, N))
	wg := sync.WaitGroup{}
	for x := 0; x < N; x++ {
		wg.Add(1)
		x := x
		go func() {
			defer wg.Done()
			for y := 0; y < N; y++ {
				idx := IX(x, y)
				col := gradient.At(f.Density[idx])
				canvas.Set(x, y, col)
			}
		}()
	}
	wg.Wait()
	return canvas
}

func (f *Fluid) Step() {
	Diffuse(1, f.Vx0[:], f.Vx[:], f.Diff, f.Dt)
	Diffuse(2, f.Vy0[:], f.Vy[:], f.Diff, f.Dt)
	Project(f.Vx0[:], f.Vy0[:], f.Vx[:], f.Vy[:])
	Advect(1, f.Vx[:], f.Vx0[:], f.Vx0[:], f.Vy0[:], f.Dt)
	Advect(2, f.Vy[:], f.Vy0[:], f.Vx0[:], f.Vy0[:], f.Dt)
	Project(f.Vx[:], f.Vy[:], f.Vx0[:], f.Vy0[:])
	Diffuse(0, f.S[:], f.Density[:], f.Diff, f.Dt)
	Advect(0, f.Density[:], f.S[:], f.Vx[:], f.Vy[:], f.Dt)
}

func (f *Fluid) AddDensity(x, y int, density float64) {
	idx := IX(x, y)
	f.Density[idx] += density
}

func (f *Fluid) AddVelocity(x, y int, vx, vy float64) {
	idx := IX(x, y)
	f.Vx[idx] += vx
	f.Vy[idx] += vy
}

func SetBnd(b int, x []float64) {
	for i := 0; i < N-1; i++ {
		i := i
		idx1 := IX(i, 0)
		idx2 := IX(i, N-1)
		x[idx1] = x[IX(i, 1)]
		x[idx2] = x[IX(i, N-2)]
		if b == 2 {
			x[idx1] = -x[idx1]
			x[idx2] = -x[idx2]
		}
	}
	for j := 0; j < N-1; j++ {
		j := j
		idx1 := IX(0, j)
		idx2 := IX(N-1, j)
		x[idx1] = x[IX(1, j)]
		x[idx2] = x[IX(N-2, j)]
		if b == 1 {
			x[idx1] = -x[idx1]
			x[idx2] = -x[idx2]
		}
	}
	x[IX(0, 0)] = 0.5 * (x[IX(1, 0)] + x[IX(0, 1)])
	x[IX(0, N-1)] = 0.5 * (x[IX(1, N-1)] + x[IX(0, N-2)])
	x[IX(N-1, 0)] = 0.5 * (x[IX(N-2, 0)] + x[IX(N-1, 1)])
	x[IX(N-1, N-1)] = 0.5 * (x[IX(N-2, N-1)] + x[IX(N-1, N-2)])

}

func Diffuse(b int, x, x0 []float64, diff, dt float64) {
	a := dt * diff * (N - 2) * (N - 2)
	LinearSolve(b, x, x0, a, 1+6*a)
}

func LinearSolve(b int, x, x0 []float64, a, c float64) {
	invC := 1 / c
	wg := sync.WaitGroup{}
	for k := 0; k < LinearSolveIterations; k++ {
		for j := 1; j < N-1; j++ {
			j := j
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 1; i < N-1; i++ {
					idx := IX(i, j)
					x[idx] = (x0[idx] + a*(x[IX(i+1, j)]+x[IX(i-1, j)]+x[IX(i, j+1)]+x[IX(i, j-1)])) * invC
				}
			}()
		}
	}
	wg.Wait()
	SetBnd(b, x)
}

func Project(vx, vy, p, div []float64) {
	wg := sync.WaitGroup{}
	for j := 1; j < N-1; j++ {
		j := j
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i < N-1; i++ {
				idx := IX(i, j)
				div[idx] = -0.5 * (vx[IX(i+1, j)] - vx[IX(i-1, j)] + vy[IX(i, j+1)] - vy[IX(i, j-1)]) / N
				p[IX(i, j)] = 0
			}
		}()
	}
	wg.Wait()
	SetBnd(0, div)
	SetBnd(0, p)
	LinearSolve(0, p, div, 1, 6)
	for j := 1; j < N-1; j++ {
		j := j
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i < N-1; i++ {
				vx[IX(i, j)] -= 0.5 * (p[IX(i+1, j)] - p[IX(i-1, j)]) * N
				vy[IX(i, j)] -= 0.5 * (p[IX(i, j+1)] - p[IX(i, j-1)]) * N
			}
		}()
	}
	wg.Wait()
	SetBnd(1, vx)
	SetBnd(2, vy)
}

func Advect(b int, d, d0, vx, vy []float64, dt float64) {
	dtx := dt * (N - 2)
	dty := dtx

	wg := sync.WaitGroup{}
	for j := 1; j < N-1; j++ {
		wg.Add(1)
		j := j
		jFloat := float64(j)
		go func() {
			defer wg.Done()
			for i := 1; i < N-1; i++ {
				iFloat := float64(i)
				idx := IX(i, j)
				tmp1 := dtx * vx[idx]
				tmp2 := dty * vy[idx]
				x := iFloat - tmp1
				y := jFloat - tmp2
				if x < 0.5 {
					x = 0.5
				}
				if x > NFloat+0.5 {
					x = NFloat + 0.5
				}
				i0 := math.Floor(x)
				i1 := i0 + 1
				if y < 0.5 {
					y = 0.5
				}
				if y > NFloat+0.5 {
					y = NFloat + 0.5
				}
				j0 := math.Floor(y)
				j1 := j0 + 1

				s1 := x - i0
				s0 := 1 - s1
				t1 := y - j0
				t0 := 1 - t1

				i0i := int(i0)
				i1i := int(i1)
				j0i := int(j0)
				j1i := int(j1)

				d[idx] = s0*(t0*d0[IX(i0i, j0i)]+t1*d0[IX(i0i, j1i)]) +
					s1*(t0*d0[IX(i1i, j0i)]+t1*d0[IX(i1i, j1i)])
			}
		}()

	}
	wg.Wait()
	SetBnd(b, d)
}

func NewGame() *Game {
	fluid := NewFluid(0.5, 0, 0)
	g := &Game{
		Width:  N,
		Height: N,
		Fluid:  fluid,
	}
	return g
}

func (g *Game) Update() error {
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			g.Fluid.AddDensity(N/2+i, N/2+j, 0.1)
			g.Fluid.AddVelocity(N/2+i, N/2+j, float64(i), float64(j))
		}
	}
	g.Fluid.Step()
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	canvas := g.Fluid.Draw()
	s.ReplacePixels(canvas.Pix)
	utils.DebugInfo(s)
}

func (g *Game) Layout(oW, oH int) (int, int) {
	return g.Width, g.Height
}
