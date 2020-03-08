package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// Monster body size
var MBS int32 = 18
// Monster move speed
var MMS int32 = 3
// Monster max health
var MMH int32 = 50
// Monster aggro dist
var MAD float64 = 250.0
// Monster timers count
var MTC int32 = 1

type MonsterTimers int32
const (
	MONSTER_TAKE_DAMAGE = iota
)

type Monster struct {
	X 						int32
	Y 						int32
	Radius				float32
	MoveSpeed		 	int32
	Hp					 	int32
	MaxHp				 	int32
	AggroDist			float64
	Animations		Timers
	Color 				rl.Color
}

func (monster *Monster) Init(x, y int32) {
	monster.X = x
	monster.Y = y
	monster.Radius = float32(MBS / 2)
	monster.MoveSpeed = MMS
	monster.Hp = MMH
	monster.MaxHp = MMH
	monster.AggroDist = MAD
	monster.Animations.Init(MTC)
	monster.Color = rl.NewColor(57, 57, 57, 255)
}

func (monster *Monster) GetHitbox() (rl.Vector2, float32) {
	return rl.Vector2 {float32(monster.X), float32(monster.Y)}, monster.Radius
}


func (monster *Monster) TakeDamage(damage int32) {
	if damage > 0 {
		monster.Hp -= damage
		if monster.Hp - damage < 0 {
			monster.Hp = 0
		}
		monster.Animations.Values[MONSTER_TAKE_DAMAGE] = 5
	}
}

func (monster *Monster) HandleAnimation(notEnded []int32) {
	for i := 0; i<len(notEnded); i++ {
		switch notEnded[i] {
		case MONSTER_TAKE_DAMAGE:
			var i float32
			for i = 0; i<2; i++ {
				rl.DrawCircleLines(monster.X, monster.Y, monster.Radius-i, rl.Red)
			}
			break

		//ADD ANIMATION HANDLER HERE
		}
	}
}

func (monster *Monster) SpreadCoins() []Coin {
	var coins []Coin = make([]Coin, 2)
	coins[0].Init(monster.X + r1.Int31() % 30 - 15, monster.Y + r1.Int31() % 30 - 15)
	coins[1].Init(monster.X + r1.Int31() % 30 - 15, monster.Y + r1.Int31() % 30 - 15)
	return coins
}

func (monster *Monster) Draw() {
	// Monster
	rl.DrawCircle(monster.X, monster.Y, monster.Radius, monster.Color)
	// Animations
	notEnded, _ := monster.Animations.Decrement()
	monster.HandleAnimation(notEnded)
}