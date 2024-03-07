package pillarhandler

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var spawnPillarChan = make(chan struct{})

func StartPillarSpawner(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			spawnPillarChan <- struct{}{}
		}
	}()
}

func (ph *PillarHandler) SpawnPillar() {
	ph.pillarPairs = append(ph.pillarPairs, RandomPillar())
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
	return pillarHandler
}

// TODO
func (ph *PillarHandler) Update(dt float32) {

	select {
	case <-spawnPillarChan:
		ph.SpawnPillar()
	default:
	}

	for i := len(ph.pillarPairs) - 1; i >= 0; i-- {
		pillarPair := ph.pillarPairs[i]
		pillarPair.posX -= int32(ph.pillarSpeed * dt)

		if pillarPair.posX < -50 {
			ph.pillarPairs = append(ph.pillarPairs[:i], ph.pillarPairs[i+1:]...)
		}
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
