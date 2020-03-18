package class

import (
	"github.com/gen2brain/raylib-go/raylib"
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

type Item struct {
	X 					int32
	Y 					int32
	Size 				int32
	Level				int32
	Name				ItemName
	Description	string
	Selected 		bool
}

/* Init */

func (item *Item) initWaterBoots() {
	item.Description = "The water boots allow you to walk on the water."
}

func (item *Item) initHeartOfSteel() {
	item.Description = "The heart of steel increases your maximum health points."
}

func (item *Item) initTurboReactor() {
	item.Description = "The turbo reactor increases your maximum speed."
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

/* Effect */

func (item *Item) applyEffectWaterBoots(_map *Map, value bool) {
	_map.CurrPlayer.Settings[CAN_WALK_ON_WATER] = value
}

func (item *Item) applyEffectHeartOfSteel(_map *Map, value int32) {
	var hpPercentage float32 = float32(_map.CurrPlayer.Hp) / float32(_map.CurrPlayer.MaxHp)
	_map.CurrPlayer.MaxHp += value
	_map.CurrPlayer.Hp = int32(float32(_map.CurrPlayer.MaxHp) * hpPercentage)
}

func (item *Item) applyEffectTurboReactor(_map *Map, value int32) {
	_map.CurrPlayer.MaxSpeed += value
}

func (item *Item) ApplyEffect(_map *Map) {
	switch item.Name {
		case WATER_BOOTS:
			item.applyEffectWaterBoots(_map, true)
			break
		case HEART_OF_STEEL:
			item.applyEffectHeartOfSteel(_map, 50)
			break
		case TURBO_REACTOR:
			item.applyEffectTurboReactor(_map, 1)
			break
	}
}

func (item *Item) removeEffectWaterBoots(_map *Map) {
	item.applyEffectWaterBoots(_map, false)
}

func (item *Item) removeEffectHeartOfSteel(_map *Map) {
	item.applyEffectHeartOfSteel(_map, -50)
}

func (item *Item) removeEffectTurboReactor(_map *Map) {
	item.applyEffectTurboReactor(_map, -1)
}

func (item *Item) RemoveEffect(_map *Map) {
	switch item.Name {
		case WATER_BOOTS:
			item.removeEffectWaterBoots(_map)
			break
		case HEART_OF_STEEL:
			item.removeEffectHeartOfSteel(_map)
			break;
		case TURBO_REACTOR:
			item.removeEffectTurboReactor(_map)
			break
	}
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