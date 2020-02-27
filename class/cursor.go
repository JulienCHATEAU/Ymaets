package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Cursor struct {
	X 		int32
	Y 		int32
	Size	int32
	Color rl.Color
}

func (cursor *Cursor) Init() {
	cursor.X = 0
	cursor.Y = 0
	cursor.Size = 2
	cursor.Color = rl.Red
}

func (cursor *Cursor) Draw() {
	rl.DrawCircle(cursor.X, cursor.Y, float32(cursor.Size), cursor.Color)
}