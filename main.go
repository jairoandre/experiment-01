package main

import (
	"experiment-01/utils"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	vec "github.com/jairoandre/vector-go"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	Canvas         *image.RGBA
	Particles      []*Particle
	Tick           int
	LastUpdateTime int64
	LastDrawTime   int64
	Paused         bool
}

func (g *Game) FadeEffect() {
	uni := image.NewUniform(color.Black)
	alpha := image.NewUniform(color.Alpha{A: 0xa3})
	draw.DrawMask(g.Canvas, g.Canvas.Bounds(), uni, image.Point{}, alpha, image.Point{}, draw.Over)
}

func updateParticles(canvas *image.RGBA, mouseVec2d vec.Vector2d, particles []*Particle) {
	for _, particle := range particles {
		force := mouseVec2d.Sub(particle.Pos).Normalize().Mul(0.1)
		particle.ApplyForce(force)
		particle.Update()
		particle.Draw(canvas)
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		g.Paused = !g.Paused
	}
	if g.Tick > 0 && g.Paused {
		return nil
	}
	start := time.Now()
	//g.FadeEffect()
	g.Canvas = image.NewRGBA(image.Rect(0, 0, Width, Height))
	mx, my := ebiten.CursorPosition()
	mouseVec2d := vec.NewVec2dFromInt(mx, my)
	wg := sync.WaitGroup{}
	for i := 0; i < len(g.Particles)-sliceRoutineSize; i += sliceRoutineSize {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			updateParticles(g.Canvas, mouseVec2d, g.Particles[i:i+sliceRoutineSize])
		}()
	}
	wg.Wait()
	end := time.Now()
	delta := end.Sub(start)
	if g.Tick%100 == 0 {
		g.LastUpdateTime = delta.Milliseconds()
	}
	g.Tick += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	start := time.Now()
	screen.ReplacePixels(g.Canvas.Pix)
	end := time.Now()
	delta := end.Sub(start)
	if g.Tick%100 == 0 {
		g.LastDrawTime = delta.Milliseconds()
	}
	utils.DebugInfo(screen)
	utils.DebugInfoMessage(screen, fmt.Sprintf("\nParticles: %d", NumParticles))
	utils.DebugInfoMessage(screen, fmt.Sprintf("\n\nUpdate time: %dms", g.LastUpdateTime))
	utils.DebugInfoMessage(screen, fmt.Sprintf("\n\nDraw time: %dms", g.LastDrawTime))

}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

const (
	Width            = 1280 / 2
	Height           = 720 / 2
	NumParticles     = 100000
	sliceRoutineSize = 1000
)

func main() {
	particles := make([]*Particle, 0)
	for i := 0; i < NumParticles; i++ {
		x := int(rand.Float64() * Width)
		y := int(rand.Float64() * Height)
		particle := NewParticle(x, y, 1)
		particles = append(particles, &particle)
	}
	g := Game{
		Canvas:    image.NewRGBA(image.Rect(0, 0, Width, Height)),
		Particles: particles,
		Paused:    true,
	}

	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Experiment")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
