package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/stumburs/floppy_gopher/gopher"
	pillarhandler "github.com/stumburs/floppy_gopher/pillar_handler"
)

func main() {

	rl.InitWindow(640, 960, "floppy gopher")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// vars
	gopher := gopher.NewGopher(50, 50)
	pillarHandler := pillarhandler.NewPillarHandler()

	for !rl.WindowShouldClose() {
		// Update
		dt := rl.GetFrameTime()
		pillarHandler.Update(dt)
		gopher.Update(dt)

		// fmt.Printf("Pillars: %v\n", pillarHandler.PillarCount())

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(201, 233, 246, 255))
		pillarHandler.Draw()
		gopher.Draw()
		rl.DrawFPS(20, 20)
		rl.EndDrawing()
	}
}
