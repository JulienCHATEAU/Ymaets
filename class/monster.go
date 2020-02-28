package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// Monster body size
var MBS int32 = 18
// Monster move speed
var MMS int32 = 3

type Monster struct {
	X 						int32
	Y 						int32
	Radius				float32
	MoveSpeed		 	int32
	Color 				rl.Color
}

func (monster *Monster) Init(x, y int32) {
	monster.X = x
	monster.Y = y
	monster.Radius = float32(MBS / 2)
	monster.MoveSpeed = MMS
	monster.Color = rl.Magenta
}

func (monster *Monster) GetHitbox() (rl.Vector2, float32) {
	return rl.Vector2 {float32(monster.X), float32(monster.Y)}, monster.Radius
}

func (monster *Monster) Draw() {
	rl.DrawCircle(monster.X, monster.Y, monster.Radius, monster.Color)
}