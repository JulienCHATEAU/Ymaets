package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Wall struct {
	X 				int32
	Y 				int32
	Width			int32
	Height		int32
	Crossable bool
	Walkable 	bool
	Color 		rl.Color
}

func (wall *Wall) init(x, y, width, height int32, crossable, walkable bool, color rl.Color) {
	wall.X = x
	wall.Y = y
	wall.Width = width
	wall.Height = height
	wall.Crossable = crossable
	wall.Walkable = walkable
	wall.Color = color
}

func (wall *Wall) InitLava(x, y, width, height int32) {
	wall.init(x, y, width, height, true, true, rl.NewColor(245, 90, 0, 255))
}

func (wall *Wall) InitWater(x, y, width, height int32) {
	wall.init(x, y, width, height, true, false, rl.NewColor(104, 215, 250, 255))
}

func (wall *Wall) InitWall(x, y, width, height int32, color rl.Color) {
	wall.init(x, y, width, height, false, false, color)
}

func (wall *Wall) InitBorder(x, y, width, height int32) {
	wall.InitWall(x, y, width, height, rl.Brown)
}

func (wall *Wall) Draw() {
	rl.DrawRectangle(wall.X, wall.Y, wall.Width, wall.Height, wall.Color)
}

func (wall *Wall) GetHitbox() rl.Rectangle {
	return rl.Rectangle{float32(wall.X), float32(wall.Y), float32(wall.Width), float32(wall.Height)}
}