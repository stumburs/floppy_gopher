package game

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/stumburs/floppy_gopher/gopher"
	pillarhandler "github.com/stumburs/floppy_gopher/pillar_handler"
)

func RunGame() {
	rl.InitWindow(640, 960, "floppy gopher")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Create gopher and pillar handler instances
	g := gopher.NewGopher(50, 50)
	ph := pillarhandler.NewPillarHandler()
	pillarhandler.StartPillarSpawner(time.Second * 2)

	for !rl.WindowShouldClose() {
		// Update
		dt := rl.GetFrameTime()
		ph.Update(dt)
		g.Update(dt)

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(201, 233, 246, 255))
		ph.Draw()
		g.Draw()
		rl.DrawFPS(20, 20)
		rl.DrawText(fmt.Sprintf("Pillars: %d", ph.PillarCount()), 20, 40, 20, rl.Lime)
		rl.EndDrawing()
	}
}
