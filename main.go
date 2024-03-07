package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/stumburs/floppy_gopher/gopher"
)

func main() {

	rl.InitWindow(640, 960, "floppy gopher")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// vars
	gopher := gopher.NewGopher(50, 50)

	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()
		// Update
		gopher.Update(dt)

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		gopher.Draw()
		rl.DrawFPS(20, 20)
		rl.EndDrawing()
	}
}
