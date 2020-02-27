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
	cursor.Size = 4
	cursor.Color = rl.NewColor(155, 19, 19, 255)
}

func (cursor *Cursor) Draw() {
	rl.DrawRectangle(cursor.X-1, cursor.Y-7, 2, 14, cursor.Color)
	rl.DrawRectangle(cursor.X-7, cursor.Y-1, 14, 2, cursor.Color)
	rl.DrawCircleLines(cursor.X, cursor.Y, float32(cursor.Size), cursor.Color)
}