package class

import (
	// "fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

type WallTimers int32
const (
	DAMAGE_DEALT = iota
)

type WallType int32
const (
	Solid = iota
	Water
	Lava
)

// // Wall timers count
// var WTC int32 = 1
// Lava walk damage
var LWD int32 = 2
// Lava damage tick
var LDT int32 = 80

type Wall struct {
	X 					int32
	Y 					int32
	Width				int32
	Height			int32
	Crossable 	bool
	Type 				WallType
	Walkable 		bool
	WalkDamage 	int32
	Animations	Timers
	Color 			rl.Color
}

func RemoveWall(index *int, walls []Wall) []Wall {
	for i := *index; i<len(walls)-1; i++ {
		walls[i] = walls[i+1]
	}
	*(index)--
	return walls[:len(walls)-1]
}

func (wall *Wall) init(x, y, width, height int32, crossable, walkable bool, color rl.Color) {
	wall.X = x
	wall.Y = y
	wall.Width = width
	wall.Height = height
	wall.Crossable = crossable
	wall.Walkable = walkable
	wall.WalkDamage = LWD
	// wall.Animations.Init(WTC)
	wall.Color = color
}

func (wall *Wall) InitLava(x, y, width, height int32) {
	wall.init(x, y, width, height, true, true, rl.NewColor(245, 90, 0, 200))
	wall.Type = Lava
}

func (wall *Wall) InitWater(x, y, width, height int32) {
	wall.init(x, y, width, height, true, false, rl.NewColor(104, 215, 250, 255))
	wall.Type = Water
}

func (wall *Wall) InitWall(x, y, width, height int32, color rl.Color) {
	wall.init(x, y, width, height, false, false, color)
	wall.Type = Solid
}

func (wall *Wall) InitBorder(x, y, width, height int32) {
	wall.InitWall(x, y, width, height, rl.Brown)
	wall.Type = Solid
}

// func (wall *Wall) HandleAnimation(notEnded []int32, justEnded []int32) {
// 	for i := 0; i<len(notEnded); i++ {
// 		switch notEnded[i] {
// 		case DAMAGE_DEALT:
// 			wall.WalkDamage = 0
// 			break

// 		//ADD ANIMATION HANDLER HERE
// 		}
// 	}

// 	for i := 0; i<len(justEnded); i++ {
// 		switch justEnded[i] {
// 		case DAMAGE_DEALT:
// 			wall.WalkDamage = LWD
// 			break

// 		//ADD ANIMATION HANDLER HERE
// 		}
// 	}
// }

func (wall *Wall) Draw() {
	// Wall
	rl.DrawRectangle(wall.X, wall.Y, wall.Width, wall.Height, wall.Color)
	//Animations
	// wall.Animations.Decrement()
}

func (wall *Wall) GetHitbox() rl.Rectangle {
	return rl.Rectangle{float32(wall.X), float32(wall.Y), float32(wall.Width), float32(wall.Height)}
}