package class

import (
	"github.com/gen2brain/raylib-go/raylib"
	"strconv"
	util "Ymaets/util"
)

type ItemName string
const (
	WATER_BOOTS = "Water boots"
	HEART_OF_STEEL = "Heart of steel"
	TURBO_REACTOR = "Turbo reactor"
)

// Item body size
var IBS int32 = 30
// Item Max levl
var IML int32 = 3

type Item struct {
	X 									int32
	Y 									int32
	Size 								int32
	Level								int32
	Name								ItemName
	Description					string
	LevelUpDescription	[]string
	Selected 						bool
}

/* Init */

func (item *Item) initWaterBoots() {
	item.Description = "The water boots allow you to walk on the water."
	item.LevelUpDescription = []string {
		"Water is now walkable",
		"On water, Move speed : +1",
		"On water, Regen : 0.1 Hp/sec",
	} 
}

func (item *Item) initHeartOfSteel() {
	item.Description = "The heart of steel increases your maximum health points."
	item.LevelUpDescription = []string {
		"Max Hp : +25",
		"Max Hp : +25",
		"Max Hp : +50",
	} 
}

func (item *Item) initTurboReactor() {
	item.Description = "The turbo reactor increases your move speed."
	item.LevelUpDescription = []string {
		"Move speed : +1",
		"Move speed : +1",
		"Move speed : +1",
	} 
}

func (item *Item) Init(x, y int32, name ItemName) {
	item.X = x
	item.Y = y
	item.Level = 1
	item.Size = IBS
	switch name {
		case WATER_BOOTS:
			item.initWaterBoots()
			break
		case HEART_OF_STEEL:
			item.initHeartOfSteel()
			break
		case TURBO_REACTOR:
			item.initTurboReactor()
			break
	}
	item.Name = name
	item.Selected = false
}

func (item *Item) GetLevelUpDescription(i int32) string {
	space := ""
	if i == 0 {
		space += " "
	}
	return "Lvl " + space + strconv.Itoa(int(i+1)) + ")   " + item.LevelUpDescription[i]
}

/* Effect */

func (item *Item) setWaterWalkable(_map *Map, value bool) {
	_map.CurrPlayer.Settings[CAN_WALK_ON_WATER] = value
}

func (item *Item) addHealthPoints(_map *Map, value int32) {
	var hpPercentage float32 = float32(_map.CurrPlayer.Stats.Hp) / float32(_map.CurrPlayer.Stats.MaxHp)
	_map.CurrPlayer.Stats.MaxHp += value
	_map.CurrPlayer.Stats.Hp = int32(float32(_map.CurrPlayer.Stats.MaxHp) * hpPercentage)
}

func (item *Item) addSpeed(_map *Map, value int32) {
	_map.CurrPlayer.Stats.MaxSpeed += value
}

func (item *Item) applyEffectWaterBoots(_map *Map, prod int32) {
	item.setWaterWalkable(_map, prod == 1)
	if item.Level > 1 {//lvl2

	}
	if item.Level > 2 {//lvl3

	}
}

func (item *Item) applyEffectHeartOfSteel(_map *Map, prod int32) {
	var healthPoints int32 = 25
	if item.Level > 1 {//lvl2
		healthPoints += 25
	}
	if item.Level > 2 {//lvl3
		healthPoints += 50
	}
	item.addHealthPoints(_map, prod * healthPoints)
}

func (item *Item) applyEffectTurboReactor(_map *Map, prod int32) {
	var speed int32 = 1
	if item.Level > 1 {//lvl2
		speed++
	}
	if item.Level > 2 {//lvl3
		speed++
	}
	item.addSpeed(_map, prod * speed)
}

func (item *Item) ApplyEffect(_map *Map) {
	switch item.Name {
		case WATER_BOOTS:
			item.applyEffectWaterBoots(_map, 1)
			break
		case HEART_OF_STEEL:
			item.applyEffectHeartOfSteel(_map, 1)
			break
		case TURBO_REACTOR:
			item.applyEffectTurboReactor(_map, 1)
			break
	}
}

func (item *Item) removeEffectWaterBoots(_map *Map, prod int32) {
	item.applyEffectWaterBoots(_map, prod)
}

func (item *Item) removeEffectHeartOfSteel(_map *Map, prod int32) {
	item.applyEffectHeartOfSteel(_map, prod)
}

func (item *Item) removeEffectTurboReactor(_map *Map, prod int32) {
	item.applyEffectTurboReactor(_map, prod)
}

func (item *Item) RemoveEffect(_map *Map) {
	switch item.Name {
		case WATER_BOOTS:
			item.removeEffectWaterBoots(_map, -1)
			break
		case HEART_OF_STEEL:
			item.removeEffectHeartOfSteel(_map, -1)
			break;
		case TURBO_REACTOR:
			item.removeEffectTurboReactor(_map, -1)
			break
	}
}

func (item *Item) LevelUp(_map *Map) bool {
	possible := item.CanLevelUp()
	if possible {
		item.RemoveEffect(_map)
		item.Level++
		item.ApplyEffect(_map)
	}
	return possible
}

func (item *Item) CanLevelUp() bool {
	if item.Level < IML {
		return true
	}
	return false
}

/* Draw */

func (item *Item) drawWaterBoots() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.DarkBlue)
	
}

func (item *Item) drawHeartOfSteel() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.Pink)
}

func (item *Item) drawTurboReactor() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.Green)
}


func (item *Item) Draw() {
	if item.Selected {
		var margin int32 = 3
		rl.DrawRectangle(item.X - margin, item.Y - margin, item.Size + margin*2, item.Size + margin*2, rl.RayWhite)
		rl.DrawRectangleLinesEx(util.ToRectangle(item.X - margin, item.Y - margin, item.Size + margin*2, item.Size + margin*2), 3, rl.Black)
	} else {
		rl.DrawRectangle(item.X, item.Y, item.Size, item.Size, rl.RayWhite)
		rl.DrawRectangleLinesEx(util.ToRectangle(item.X, item.Y, item.Size, item.Size), 2, rl.Black)
	}
	switch item.Name {
		case WATER_BOOTS:
			item.drawWaterBoots()
			break
		case HEART_OF_STEEL:
			item.drawHeartOfSteel()
			break;
		case TURBO_REACTOR:
			item.drawTurboReactor()
			break
	}
}

//////

func (item *Item) GetHitbox() rl.Rectangle {
	return rl.Rectangle{float32(item.X), float32(item.Y), float32(item.Size), float32(item.Size)}
}