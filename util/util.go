package util

import (
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

func ToRectangle(x, y, width, height int32) rl.Rectangle {
	return rl.Rectangle {float32(x), float32(y), float32(width), float32(height)}
}

func PointsDistance(x1, y1, x2, y2 int32) (float64) {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func DrawHealthBar(hp, maxHp, x, y, objectSize int32) {
	var healthBarColor rl.Color = rl.Gray
	hpPercentage := float32(hp) / float32(maxHp)
	if hpPercentage > 0.5 {
		healthBarColor = rl.Green
	} else if hpPercentage > 0.20 {
		healthBarColor = rl.Orange
	} else if hpPercentage > 0 {
		healthBarColor = rl.Red
	}
	healthBarMaxWidth := float32(objectSize + 4)
	var barHeight int32 = 3
	rl.DrawRectangle(x - 2, y - 5 - barHeight, int32(healthBarMaxWidth * hpPercentage), barHeight, healthBarColor);
}