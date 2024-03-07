package pillarhandler

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type pillarPair struct {
	gapSize  float32
	posX     int32
	midpoint int32
}

type PillarHandler struct {
	pillarPairs []*pillarPair
	pillarSpeed float32
}

func NewPillarHandler() *PillarHandler {
	pillarHandler := &PillarHandler{
		pillarPairs: make([]*pillarPair, 0),
		pillarSpeed: 200,
	}

	// Add one random pillar
	pillarHandler.pillarPairs = append(pillarHandler.pillarPairs, RandomPillar())

	return pillarHandler
}

// TODO
func (ph *PillarHandler) Update(dt float32) {
	//ph.pillarPairs = append(ph.pillarPairs, RandomPillar())

	for _, pillarPair := range ph.pillarPairs {
		pillarPair.posX -= int32(ph.pillarSpeed * dt)
	}
}

func (ph *PillarHandler) Draw() {
	for _, pillarPair := range ph.pillarPairs {
		// Top
		rl.DrawRectangle(int32(pillarPair.posX), 0, 50, pillarPair.midpoint-int32(pillarPair.gapSize), rl.Red)
		// Bottom
		rl.DrawRectangle(int32(pillarPair.posX), pillarPair.midpoint+int32(pillarPair.gapSize), 50, int32(rl.GetScreenHeight())-pillarPair.midpoint-int32(pillarPair.gapSize), rl.Red)
	}
}

func (ph *PillarHandler) PillarCount() int {
	return len(ph.pillarPairs)
}

func RandomPillar() *pillarPair {
	return &pillarPair{
		gapSize:  100,
		posX:     int32(rl.GetScreenWidth()) + 50,
		midpoint: int32(float32(rl.GetScreenHeight()) / 2),
	}
}
