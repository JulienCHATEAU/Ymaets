package class

import (
	"Ymaets/util"
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

// Monster body size
var MBS int32 = 20
// Monster move speed
var MMS int32 = 3
// Monster max health
var MMH int32 = 75
// Monster max attack
var MMA int32 = 48
// Monster max defense
var MMD int32 = 48
// Monster aggro dist
var MAD float64 = 250.0
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
// Kamikaze shot range (px)
var KSR int32 = 200
// Sniper shot range (px)
var SSR int32 = 350
// Kamikaze fire cooldown
var KFC int32 = 20
// Sniper fire cooldown
var SFC int32 = 40

// Monster timers count
var MTC int32 = 5

type MonsterTimers int32
const (
	MONSTER_TAKE_DAMAGE = iota
	MONSTER_FIRE_COOLDOWN
	MONSTER_LAVA_DAMAGE
	CRIT_DAMAGE
	EXCLAMATION_POINT
)

type MonsterType int32
const (
	KAMIKAZE = iota
	ONE_CANON_KAMIKAZE
	SNIPER
)

var monsters []MonsterType = []MonsterType {KAMIKAZE, ONE_CANON_KAMIKAZE, SNIPER}

func GetMonsters() []MonsterType {
	b := make([]MonsterType, len(monsters))
	copy(b, monsters)
	return b
}

type Monster struct {
	X 						int32
	Y 						int32
	Radius				float32
	MoveSpeed		 	int32
	Stats					util.Stat
	FireCooldown  int32
	HasCanon			bool
	Aggressive		bool
	GivenExp	int32
	Type					MonsterType
	Ori						Orientation
	Settings			map[Setting]bool
	LavaExit			Orientation
	AggroDist			float64
	Animations		Timers
	Color 				rl.Color
}

/* Init */

func (monster *Monster) initKamikaze() {
	monster.HasCanon = false
	monster.Color = rl.NewColor(144, 227, 217, 255)
	monster.FireCooldown = KFC
	monster.Stats.Init(MMS, MMH + 5, MMA + 1, MMD, KSR, 0, 0)
	monster.AggroDist = float64(monster.Stats.Range + 20)
	monster.GivenExp = 5
}

func (monster *Monster) initOneCanonKamikaze() {
	monster.HasCanon = true
	monster.Color = rl.NewColor(255, 112, 0, 255)
	monster.FireCooldown = KFC
	monster.Stats.Init(MMS, MMH + 5, MMA, MMD + 1, KSR, 0, 0)
	monster.AggroDist = float64(monster.Stats.Range + 20)
	monster.GivenExp = 6
}

func (monster *Monster) initSniper() {
	monster.HasCanon = true
	monster.Color = rl.NewColor(16, 57, 120, 255)
	monster.FireCooldown = SFC
	monster.Stats.Init(MMS - 1, MMH, MMA + 3, MMD, SSR, 0, 0)
	monster.AggroDist = float64(monster.Stats.Range + 80)
	monster.GivenExp = 7
}

func (monster *Monster) Init(x, y int32, monsterType MonsterType) {
	monster.X = x
	monster.Y = y
	monster.Radius = float32(MBS / 2)
	monster.Ori = NORTH
	monster.Settings = make(map[Setting]bool)
	monster.LavaExit = NONE
	monster.Aggressive = false
	monster.Animations.Init(MTC)
	monster.Animations.Decrements[CRIT_DAMAGE] = 15
	monster.Animations.Decrements[EXCLAMATION_POINT] = 15
	switch monsterType {
	case KAMIKAZE:
		monster.initKamikaze()
		break
	case ONE_CANON_KAMIKAZE:
		monster.initOneCanonKamikaze()
		break
	case SNIPER:
		monster.initSniper()
		break
	}
	monster.Type = monsterType
}

/* MOVE */

func (monster *Monster) moveKamikaze(_map *Map) {
	var dx int32 = 0
	var dy int32 = 0
	if monster.X < _map.CurrPlayer.X + PBS/2 {
		dx = monster.Stats.MaxSpeed
	} else if monster.X > _map.CurrPlayer.X + PBS/2 {
		dx = -monster.Stats.MaxSpeed
	}
	if monster.Y < _map.CurrPlayer.Y + PBS/2 {
		dy = monster.Stats.MaxSpeed
	} else if monster.Y > _map.CurrPlayer.Y + PBS/2 {
		dy = -monster.Stats.MaxSpeed
	}
	monster.X += dx
	monster.Y += dy
}

func (monster *Monster) moveSniper(_map *Map) {
	var dx int32 = 0
	var dy int32 = 0
	var playerDx int32 = _map.CurrPlayer.X - monster.X
	if playerDx < 0 {
		playerDx = -playerDx
	}
	var playerDy int32 = _map.CurrPlayer.Y - monster.Y
	if playerDy < 0 {
		playerDy = -playerDy
	}
	if playerDx < playerDy {
		if monster.X < _map.CurrPlayer.X + PBS/2 {
			dx = monster.Stats.MaxSpeed
		} else if monster.X > _map.CurrPlayer.X + PBS/2 {
			dx = -monster.Stats.MaxSpeed
		}
		// if monster.Y < _map.CurrPlayer.Y + PBS/2 {
		// 	dy = -monster.Stats.MaxSpeed
		// } else if monster.Y > _map.CurrPlayer.Y + PBS/2 {
		// 	dy = monster.Stats.MaxSpeed
		// }
	} else {
		if monster.Y < _map.CurrPlayer.Y + PBS/2 {
			dy = monster.Stats.MaxSpeed
		} else if monster.Y > _map.CurrPlayer.Y + PBS/2 {
			dy = -monster.Stats.MaxSpeed
		}
		// if monster.X < _map.CurrPlayer.X + PBS/2 {
		// 	dx = monster.Stats.MaxSpeed
		// } else if monster.X > _map.CurrPlayer.X + PBS/2 {
		// 	dx = -monster.Stats.MaxSpeed
		// }
	}
	monster.X += dx
	monster.Y += dy
}

func (monster *Monster) Move(_map *Map) {
	switch monster.Type {
		case KAMIKAZE:
			monster.moveKamikaze(_map)
			break
		case ONE_CANON_KAMIKAZE:
			monster.moveKamikaze(_map)
			break
		case SNIPER:
			monster.moveSniper(_map)
			break
	}
}

func (monster *Monster) HandleLavaExit() {
	if monster.Settings[IS_ON_LAVA] {
		if !monster.Settings[LAVA_EXIT_APPLIED] || monster.Settings[COLLISION_ON_LAST_MOVE] {
			monster.Settings[LAVA_EXIT_APPLIED] = true
			monster.LavaExit = ChooseInOris(oris)
		}
	} else {
		if monster.Settings[LAVA_EXIT_APPLIED] {
			monster.Settings[LAVA_EXIT_APPLIED] = false
			monster.LavaExit = NONE
		}
	}
}

func (monster *Monster) FindSeat(_map *Map) {
	if monster.Settings[IS_ON_LAVA] {
		switch monster.LavaExit {
		case NORTH:
			monster.Y -= monster.Stats.MaxSpeed
			break
	
		case SOUTH:
			monster.Y += monster.Stats.MaxSpeed
			break
	
		case EAST:
			monster.X += monster.Stats.MaxSpeed
			break
	
		case WEST:
			monster.X -= monster.Stats.MaxSpeed
			break
		}
	}
}

/* Fire */

func (monster *Monster) Fire(_map *Map) {
		shot := monster.GetShot()
		monster.Animations.Values[FIRE_COOLDOWN] = monster.FireCooldown
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

/* Collision */

func (monster *Monster) kamikazePlayerCollision(_map *Map) {
	monster.Kill()
	_map.CurrPlayer.TakeDamage(5)
}

func (monster *Monster) sniperPlayerCollision(_map *Map) {
	_map.CurrPlayer.Stats.Speed = 0
}

func (monster *Monster) PlayerCollision(_map *Map) {
	switch monster.Type {
	case KAMIKAZE:
		monster.kamikazePlayerCollision(_map)
		break
	case ONE_CANON_KAMIKAZE:
		monster.kamikazePlayerCollision(_map)
		break
	case SNIPER:
		monster.sniperPlayerCollision(_map)
		break
	}
}

/////

func (monster *Monster) GetExperience() int32 {
	return monster.GivenExp
}

func (monster *Monster) BuffDepOnStage(currStage int32) {
	monster.Stats.MaxHp += 2*currStage
	monster.Stats.Hp += 2*currStage
	monster.Stats.MaxAtt += currStage
	monster.Stats.Att += currStage
	monster.Stats.MaxDef += currStage
	monster.Stats.Def += currStage
	monster.GivenExp += currStage*3/2
}

func (monster *Monster) HandleDamageTaken(damage int32, player Player) {
	fmt.Print("To monster : ")
	fmt.Println(monster.Stats)
	fmt.Println(damage)
	if r1.Int31() % 100 < player.Stats.CritRate {
		damage += damage * 50 / 100
		monster.Animations.Values[CRIT_DAMAGE] = 350
		if player.Settings[MONEY_CRIT_BONUS] {
			player.Money += 5
		}
	}
	monster.Aggressive = true
	monster.TakeDamage(damage)
}

func (monster *Monster) Kill() {
	monster.Stats.Hp = 0
}

func (monster *Monster) IsDead() bool {
	return monster.Stats.Hp == 0
}

func (monster *Monster) GetShot() Shot {
	var shot Shot
	shot.Init(monster.Ori, monster.Color, MSS, MSH, MSW, monster.Stats.Range, 5, MONSTER, monster.GetStats())
	radius := int32(monster.Radius)
	switch monster.Ori {
	case NORTH:
		shot.X = monster.X - MCH/2
		shot.Y = monster.Y - radius - MCH - shot.Height
		break

	case SOUTH:
		shot.X = monster.X - MCH/2
		shot.Y = monster.Y + radius + MCH
		break

	case EAST:
		shot.X = monster.X + radius + MCH
		shot.Y = monster.Y - MCH/2
		break

	case WEST:
		shot.X = monster.X  - radius - MCH - shot.Width
		shot.Y = monster.Y - MCH/2
		break
	}
	return shot
}

func (monster *Monster) GetHitbox() (rl.Vector2, float32) {
	return rl.Vector2 {float32(monster.X), float32(monster.Y)}, monster.Radius
}


func (monster *Monster) TakeDamage(damage int32) {
	if damage <= 0 {
		damage = 1
	}
	monster.Stats.Hp -= damage
	if monster.Stats.Hp - damage < 0 {
		monster.Stats.Hp = 0
	}
	monster.Animations.Values[MONSTER_TAKE_DAMAGE] = 5
}

func (monster *Monster) HandleAnimation(notEnded []int32) {
	for i := 0; i<len(notEnded); i++ {
		switch notEnded[i] {
		case MONSTER_TAKE_DAMAGE:
			var i float32
			for i = 0; i<2; i++ {
				// rl.DrawCircleLines(monster.X, monster.Y, monster.Radius-i, rl.Red)
				rl.DrawCircle(monster.X, monster.Y, monster.Radius, rl.NewColor(255, 0, 0, 100))
			}
			break

		case CRIT_DAMAGE:
			opacity := monster.Animations.Values[CRIT_DAMAGE]
			if opacity > 255 {
				opacity = 255
			}
			rl.DrawText("Crit !", monster.X - 16, monster.Y - int32(monster.Radius) - 27, 15, rl.NewColor(246, 50, 27, uint8(opacity)))
			break

		case EXCLAMATION_POINT:
			opacity := monster.Animations.Values[EXCLAMATION_POINT]
			if opacity > 255 {
				opacity = 255
			}
			rl.DrawText("!", monster.X + 13, monster.Y - int32(monster.Radius) - 16, 18, rl.NewColor(246, 50, 27, uint8(opacity)))
			break

		//ADD ANIMATION HANDLER HERE
		}
	}
}

func (monster *Monster) SpreadLoots(_map *Map) {
	var coins []Coin = make([]Coin, 2)
	coins[0].Init(monster.X + r1.Int31() % 30 - 15, monster.Y + r1.Int31() % 30 - 15)
	coins[1].Init(monster.X + r1.Int31() % 30 - 15, monster.Y + r1.Int31() % 30 - 15)
	_map.CoinsCount += int32(len(coins))
	_map.Coins = append(_map.Coins, coins...)
	if r1.Int() % 200 < 1 {
		var item Item
		items := GetItems()
		length := len(items)
		rand := r1.Int() % length
		typee := items[rand]
		item.Init(monster.X + r1.Int31() % 30 - 15, monster.Y + r1.Int31() % 30 - 15, typee, false)
		_map.AddItem(item)
	}
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

func (monster Monster) GetStats() util.Stat {
	return monster.Stats
}