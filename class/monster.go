package class

import (
	"github.com/gen2brain/raylib-go/raylib"
	util "Ymaets/util"
)

// Monster body size
var MBS int32 = 18
// Monster move speed
var MMS int32 = 3
// Monster max health
var MMH int32 = 50

type Monster struct {
	X 						int32
	Y 						int32
	Radius				float32
	MoveSpeed		 	int32
	Hp					 	int32
	MaxHp				 	int32
	Color 				rl.Color
}

func (monster *Monster) Init(x, y int32) {
	monster.X = x
	monster.Y = y
	monster.Radius = float32(MBS / 2)
	monster.MoveSpeed = MMS
	monster.Hp = MMH
	monster.MaxHp = MMH
	monster.Color = rl.Magenta
}

func (monster *Monster) GetHitbox() (rl.Vector2, float32) {
	return rl.Vector2 {float32(monster.X), float32(monster.Y)}, monster.Radius
}


func (monster *Monster) TakeDamage(damage int32) {
	monster.Hp -= damage
	if monster.Hp - damage < 0 {
		monster.Hp = 0
	}
}

func (monster *Monster) Draw() {
	util.DrawHealthBar(monster.Hp, monster.MaxHp, monster.X - int32(monster.Radius), monster.Y - int32(monster.Radius), int32(monster.Radius) * 2)
	rl.DrawCircle(monster.X, monster.Y, monster.Radius, monster.Color)
}