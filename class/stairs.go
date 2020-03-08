package class

import (
	// "fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"Ymaets/util"
)

// Stairs body size
var SBS int32 = 40

var S int32 = 8
var T int32 = 2

type Stairs struct {
	X 					int32
	Y 					int32
}

func (stairs *Stairs) Init(x, y int32) {
	stairs.X = x
	stairs.Y = y
}

func (stairs *Stairs) Draw() {
	// Stairs
	rl.DrawRectangle(stairs.X, stairs.Y, SBS+2, SBS+2, rl.NewColor(50, 50, 50, 255))
	rl.DrawRectangle(stairs.X, stairs.Y + 2, S, SBS - 2, rl.LightGray)
	rl.DrawRectangle(stairs.X + 1 * S, stairs.Y + 2, T, SBS - 2, rl.NewColor(110, 110, 110, 255))

	rl.DrawRectangle(stairs.X + 1 * S + 1 * T, stairs.Y + 7, S, SBS - 7, rl.LightGray)
	rl.DrawRectangle(stairs.X + 2 * S + 1 * T, stairs.Y + 7, T, SBS - 7, rl.NewColor(110, 110, 110, 255))

	rl.DrawRectangle(stairs.X + 2 * S + 2 * T, stairs.Y + 12, S, SBS - 12, rl.LightGray)
	rl.DrawRectangle(stairs.X + 3 * S + 2 * T, stairs.Y + 12, T, SBS - 12, rl.NewColor(110, 110, 110, 255))

	rl.DrawRectangle(stairs.X + 3 * S + 3 * T, stairs.Y + 17, S, SBS - 17, rl.LightGray)
	rl.DrawRectangle(stairs.X + 4 * S + 3 * T, stairs.Y + 17, T, SBS - 17, rl.NewColor(110, 110, 110, 255))
}

func (stairs *Stairs) GetHitbox() rl.Rectangle {
	return util.ToRectangle(stairs.X, stairs.Y, SBS+2, SBS+2)
}