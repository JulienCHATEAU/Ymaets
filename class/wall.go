package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Wall struct {
	X 			int32
	Y 			int32
	Width		int32
	Height	int32
	Color 	rl.Color
}

func (wall *Wall) Init(x, y, width, height int32, color rl.Color) {
	wall.X = x
	wall.Y = y
	wall.Width = width
	wall.Height = height
	wall.Color = color
}

func (wall *Wall) InitBorder(x, y, width, height int32) {
	wall.Init(x, y, width, height, rl.Brown)
}

func (wall *Wall) Draw() {
	rl.DrawRectangle(wall.X, wall.Y, wall.Width, wall.Height, wall.Color)
}

func (wall *Wall) GetHitbox() rl.Rectangle {
	return rl.Rectangle{float32(wall.X), float32(wall.Y), float32(wall.Width), float32(wall.Height)}
}