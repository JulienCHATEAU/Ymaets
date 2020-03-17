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

func DrawHealthBar(hp, maxHp, x, y, objectSize, barHeight int32) {
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
	rl.DrawRectangle(x - 2, y - 5 - barHeight, int32(healthBarMaxWidth * hpPercentage), barHeight, healthBarColor);
}

func ShowEnterKey(x, y int32) {
	rl.DrawRectangle(x-2, y, 15, 13, rl.NewColor(100, 100, 100, 255))
	rl.DrawRectangle(x + 1, y + 5, 14, 19, rl.NewColor(100, 100, 100, 255))
	rl.DrawRectangle(x, y, 15, 12, rl.NewColor(155, 155, 155, 255))
	rl.DrawRectangle(x + 3, y + 5, 12, 18, rl.NewColor(155, 155, 155, 255))
	rl.DrawText("<-", x+3, y+2, 1, rl.Black)
}

func ShowBackspaceKey(x, y int32) {
	rl.DrawRectangle(x-3, y+1, 29, 19, rl.NewColor(100, 100, 100, 255))
	rl.DrawRectangle(x, y+1, 26, 17, rl.NewColor(155, 155, 155, 255))
	rl.DrawText("<-", x+4, y+5, 6, rl.Black)
}

func ShowClassicKey(x, y int32, key string) {
	rl.DrawRectangle(x-3, y, 23, 22, rl.NewColor(100, 100, 100, 255))
	rl.DrawRectangle(x, y, 20, 20, rl.NewColor(155, 155, 155, 255))
	rl.DrawText(key, x+6, y+6, 1, rl.Black)
}