package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Shot struct {
	X 			int32
	Y 			int32
	Width 	int32
	Height	int32
	Speed		int32
	Ori 		Orientation
	Color 	rl.Color
}

func (shot *Shot) Draw() {
	if shot.Ori == NORTH || shot.Ori == SOUTH {
		rl.DrawRectangle(shot.X, shot.Y, shot.Width, shot.Height, shot.Color);
	} else {
		rl.DrawRectangle(shot.X, shot.Y, shot.Height, shot.Width, shot.Color);
	}
}