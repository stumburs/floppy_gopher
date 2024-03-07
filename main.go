package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(640, 960, "floppy gopher")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}
