package class

import (
	"github.com/gen2brain/raylib-go/raylib"
	util "Ymaets/util"
)

type ItemName string
const (
	WATER_BOOTS = "Water boots"
	HEART_OF_STEEL = "Heart of steel"
)

// Item body size
var IBS int32 = 30

type Item struct {
	X 					int32
	Y 					int32
	Level				int32
	Name				ItemName
	Description	string
}

/* Init */

func (item *Item) initWaterBoots() {
	item.Description = "The water boots allows you to walk on the water"
}

func (item *Item) initHeartOfSteel() {
	item.Description = "The heart of steel increases your maximum health points"
}

func (item *Item) Init(x, y int32, name ItemName) {
	item.X = x
	item.Y = y
	item.Level = 1
	switch name {
		case WATER_BOOTS:
			item.initWaterBoots()
		case HEART_OF_STEEL:
			item.initHeartOfSteel()
			break;
	}
	item.Name = name
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

func (item *Item) ApplyEffect(_map *Map) {
	switch item.Name {
		case WATER_BOOTS:
			item.applyEffectWaterBoots(_map, true)
		case HEART_OF_STEEL:
			item.applyEffectHeartOfSteel(_map, 50)
			break;
	}
}

func (item *Item) removeEffectWaterBoots(_map *Map) {
	item.applyEffectWaterBoots(_map, false)
}

func (item *Item) removeEffectHeartOfSteel(_map *Map) {
	item.applyEffectHeartOfSteel(_map, -50)
}

func (item *Item) RemoveEffect(_map *Map) {
	switch item.Name {
		case WATER_BOOTS:
			item.removeEffectWaterBoots(_map)
		case HEART_OF_STEEL:
			item.removeEffectHeartOfSteel(_map)
			break;
	}
}

/* Draw */

func (item *Item) drawWaterBoots() {
	rl.DrawRectangle(item.X+5, item.Y+5, IBS-10, IBS-10, rl.DarkBlue)
	
}

func (item *Item) drawHeartOfSteel() {
	rl.DrawRectangle(item.X+5, item.Y+5, IBS-10, IBS-10, rl.Pink)
}

func (item *Item) Draw() {
	rl.DrawRectangleLinesEx(util.ToRectangle(item.X, item.Y, IBS, IBS), 2, rl.Black)
	switch item.Name {
		case WATER_BOOTS:
			item.drawWaterBoots()
		case HEART_OF_STEEL:
			item.drawHeartOfSteel()
			break;
	}
}

//////

func (item *Item) GetHitbox() rl.Rectangle {
	return rl.Rectangle{float32(item.X), float32(item.Y), float32(IBS), float32(IBS)}
}