package pillarhandler

import rl "github.com/gen2brain/raylib-go/raylib"

type pillarPair struct {
	gapSize  float32
	posX     int32
	midpoint int32
}

func RandomPillar() *pillarPair {
	return &pillarPair{
		gapSize:  100,
		posX:     int32(rl.GetScreenWidth()) + 50,
		midpoint: int32(float32(rl.GetScreenHeight()) / 2),
	}
}
