package util

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func DrawHealthBar(hp, maxHp, x, y, objectSize int32) {
	var healthBarColor rl.Color = rl.Gray
	hpPercentage := float32(hp) / float32(maxHp)
	if hpPercentage > 0.5 {
		healthBarColor = rl.Green
	} else if hpPercentage > 0.15 {
		healthBarColor = rl.Orange
	} else if hpPercentage > 0 {
		healthBarColor = rl.Red
	}
	healthBarMaxWidth := float32(objectSize + 4)
	rl.DrawRectangle(x - 2, y - 7, int32(healthBarMaxWidth * hpPercentage), 2, healthBarColor);
}