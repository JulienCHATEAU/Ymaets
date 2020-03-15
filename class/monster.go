package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// Monster body size
var MBS int32 = 20
// Monster move speed
var MMS int32 = 3
// Monster max health
var MMH int32 = 50
// Monster aggro dist
var MAD float64 = 250.0
// Monster timers count
var MTC int32 = 2
// Monster canon width
var MCW int32 = 6
// Monster canon height
var MCH int32 = 8
// Monster shot width
var MSW int32 = 10
// Monster shot height
var MSH int32 = 4
// Monster shot speed (px/frame)
var MSS int32 = 5
// Monster shot range (px)
var MSR int32 = 200
// Monster fire cooldown
var MFC int32 = 20

type MonsterTimers int32
const (
	MONSTER_TAKE_DAMAGE = iota
	MONSTER_FIRE_COOLDOWN
)

type MonsterType int32
const (
	KAMIKAZE = iota
	ONE_CANON_KAMIKAZE
)

type Monster struct {
	X 						int32
	Y 						int32
	Radius				float32
	MoveSpeed		 	int32
	Hp					 	int32
	MaxHp				 	int32
	HasCanon			bool
	Type					MonsterType
	Ori						Orientation
	AggroDist			float64
	Animations		Timers
	Color 				rl.Color
}

/* Init */

func (monster *Monster) initKamikaze() {
	monster.HasCanon = false
	monster.Color = rl.NewColor(144, 227, 217, 255)
}

func (monster *Monster) initOneCanonKamikaze() {
	monster.HasCanon = true
		monster.Color = rl.NewColor(255, 112, 0, 255)
}

func (monster *Monster) Init(x, y int32, monsterType MonsterType) {
	monster.X = x
	monster.Y = y
	monster.Radius = float32(MBS / 2)
	monster.MoveSpeed = MMS
	monster.Ori = NORTH
	monster.Hp = MMH
	monster.MaxHp = MMH
	monster.AggroDist = MAD
	monster.Animations.Init(MTC)
	switch monsterType {
	case KAMIKAZE:
		monster.initKamikaze()
		break

	case ONE_CANON_KAMIKAZE:
		monster.initOneCanonKamikaze()
		break
	}
	monster.Type = monsterType
}

/* MOVE */

func (monster *Monster) moveKamikaze(_map *Map) {
	var dx int32 = 0
	var dy int32 = 0
	if monster.X < _map.CurrPlayer.X {
		dx = monster.MoveSpeed
	} else {
		dx = -monster.MoveSpeed
	}
	if monster.Y < _map.CurrPlayer.Y {
		dy = monster.MoveSpeed
	} else {
		dy = -monster.MoveSpeed
	}
	monster.X += dx
	monster.Y += dy
}

func (monster *Monster) Move(_map *Map) {
	switch monster.Type {
		case KAMIKAZE:
			monster.moveKamikaze(_map)
		case ONE_CANON_KAMIKAZE:
			monster.moveKamikaze(_map)
			break;
	}
}

/* FIRE */

func (monster *Monster) Fire(_map *Map) {
		shot := monster.GetShot()
		monster.Animations.Values[FIRE_COOLDOWN] = MFC
		if int32(len(_map.Shots)) > _map.ShotsCount {
			_map.Shots[_map.ShotsCount] = shot
		} else {
			_map.Shots = append(_map.Shots, shot)
		}
		_map.ShotsCount++
}

/* Orient */

func (monster *Monster) Orient(_map *Map) {
	if _map.CurrPlayer.X - _map.CurrPlayer.Y > monster.X - monster.Y && _map.CurrPlayer.X + _map.CurrPlayer.Y < monster.X + monster.Y {
		monster.Ori = NORTH
	} else if _map.CurrPlayer.X - _map.CurrPlayer.Y < monster.X - monster.Y && _map.CurrPlayer.X + _map.CurrPlayer.Y > monster.X + monster.Y {
		monster.Ori = SOUTH
	} else if _map.CurrPlayer.X - _map.CurrPlayer.Y > monster.X - monster.Y && _map.CurrPlayer.X + _map.CurrPlayer.Y > monster.X + monster.Y {
		monster.Ori = EAST
	} else if _map.CurrPlayer.X - _map.CurrPlayer.Y < monster.X - monster.Y && _map.CurrPlayer.X + _map.CurrPlayer.Y < monster.X + monster.Y {
		monster.Ori = WEST
	}
}

/////

func (monster *Monster) GetShot() Shot {
	var shot Shot
	shot.Init(monster.Ori, monster.Color, MSS, MSH, MSW, MSR, MONSTER)
	radius := int32(monster.Radius)
	switch monster.Ori {
	case NORTH:
		shot.X = monster.X - MCW/2
		shot.Y = monster.Y - radius - MCH - shot.Height
		break

	case SOUTH:
		shot.X = monster.X - MCW/2
		shot.Y = monster.Y + radius
		break

	case EAST:
		shot.X = monster.X + radius
		shot.Y = monster.Y - MCW/2
		break

	case WEST:
		shot.X = monster.X  - radius - MCH - shot.Width
		shot.Y = monster.Y - MCW/2
		break
	}
	return shot
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
	rl.DrawCircle(monster.X, monster.Y, monster.Radius, rl.Black)
	rl.DrawCircle(monster.X, monster.Y, monster.Radius-1, monster.Color)
	// Canon
	radius := int32(monster.Radius)
	if monster.HasCanon {
		// Canon
		switch monster.Ori {
		case NORTH:
			rl.DrawRectangle(monster.X - MCW/2, monster.Y - radius - MCH, MCW, MCH, rl.Black);
			break

		case SOUTH:
			rl.DrawRectangle(monster.X - MCW/2, monster.Y + radius, MCW, MCH, rl.Black);
			break

		case EAST:
			rl.DrawRectangle(monster.X + radius, monster.Y - MCW/2, MCH, MCW, rl.Black);
			break

		case WEST:
			rl.DrawRectangle(monster.X  - radius - MCH, monster.Y - MCW/2, MCH, MCW, rl.Black);
			break
		}
	}
	// Animations
	notEnded, _ := monster.Animations.Decrement()
	monster.HandleAnimation(notEnded)
}