package gopher

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Gopher struct {
	Pos     rl.Vector2
	Size    rl.Vector2
	Vel     rl.Vector2
	Gravity float32
}

func (gopher Gopher) Draw() {
	rl.DrawRectangleV(gopher.Pos, gopher.Size, rl.Blue)
}

func (gopher Gopher) Update(dt float32) {
	gopher.Pos.Y -= 100 * dt
	// TODO
}

func NewGopher(sizeX, sizeY float32) *Gopher {
	gopher := Gopher{
		Pos:     rl.Vector2{X: float32(rl.GetScreenWidth())/2 - sizeX/2, Y: float32(rl.GetScreenHeight())/2 - sizeY/2},
		Size:    rl.Vector2{X: sizeX, Y: sizeY},
		Vel:     rl.Vector2Zero(),
		Gravity: 10,
	}
	return &gopher
}
