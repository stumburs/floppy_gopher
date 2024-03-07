package gopher

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Gopher struct {
	pos     rl.Vector2
	size    rl.Vector2
	vel     float32
	gravity float32
	// maxSpeed float32
}

func (gopher *Gopher) Draw() {
	rl.DrawRectangleV(gopher.pos, gopher.size, rl.Blue)
}

func (gopher *Gopher) Update(dt float32) {

	gopher.vel += gopher.gravity * dt

	if rl.IsKeyPressed(rl.KeySpace) {
		gopher.vel = -gopher.gravity * 0.3
	}

	gopher.pos.Y += gopher.vel * dt
}

func NewGopher(sizeX, sizeY float32) *Gopher {
	gopher := Gopher{
		pos:     rl.Vector2{X: float32(rl.GetScreenWidth())/4 - sizeX/2, Y: float32(rl.GetScreenHeight())/2 - sizeY/2},
		size:    rl.Vector2{X: sizeX, Y: sizeY},
		vel:     0,
		gravity: 1500,
	}
	return &gopher
}
