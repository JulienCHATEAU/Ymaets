package class

import (
	"github.com/gen2brain/raylib-go/raylib"
	"strconv"
	"strings"
	util "Ymaets/util"
)

var items []ItemName = []ItemName {WATER_BOOTS, HEART_OF_STEEL, TURBO_REACTOR, FIRE_HELMET, INVISIBLE_CAPE, ABUNDANT_PURSE, TRIFORCE_LOCKET, GOLDEN_CLOVER}
// ADD ITEMNAME ABOVE

func GetItems() []ItemName {
	b := make([]ItemName, len(items))
	copy(b, items)
	return b
}

type ItemName string
const (
	WATER_BOOTS = "Water boots"
	HEART_OF_STEEL = "Heart of steel"
	TURBO_REACTOR = "Turbo reactor"
	FIRE_HELMET = "Fire helmet"
	INVISIBLE_CAPE = "Invisible cape"
	ABUNDANT_PURSE = "Abundant purse"
	TRIFORCE_LOCKET = "Triforce locket"
	GOLDEN_CLOVER = "Golden clover"
	HEALTH_POTION = "Health potion"
	BAG_POCKET = "Bag pocket"
	// ADD ITEMNAME ABOVE
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
	Price 							int32
	OnSale							bool
	IsConsumable				bool
	Discount 						int32
	Name								ItemName
	Description					string
	LevelUpDescription	[]string
	Selected 						bool
}

/* Init */

func (item *Item) initWaterBoots() {
	item.Description = "The water boots allow you to walk on water."
	item.LevelUpDescription = []string {
		"Water is walkable",
		"On water, Move speed : +1",
		"On water, Regen : 0.5 Hp/sec",
	}
	item.Price = 2500
}

func (item *Item) initHeartOfSteel() {
	item.Description = "The heart of steel increases your maximum health points."
	item.LevelUpDescription = []string {
		"Max Hp : +20",
		"Max Hp : +30",
		"Max Hp : +50",
	} 
	item.Price = 2000
}

func (item *Item) initTurboReactor() {
	item.Description = "The turbo reactor increases your move speed."
	item.LevelUpDescription = []string {
		"Move speed : +1",
		"Move speed : +1",
		"Move speed : +1",
	} 
	item.Price = 2000
}

func (item *Item) initFireHelmet() {
	item.Description = "The fire helmet reduces the damage taken walking on lava."
	item.LevelUpDescription = []string {
		"Lava deals half damage",
		"Lava no longer deals damage",
		"On lava, Range : +50",
	} 
	item.Price = 2300
}

func (item *Item) initInvisibleCape() {
	item.Description = "The invisible cape increases your furtivity."
	item.LevelUpDescription = []string {
		"Furtivity : +40",
		"If furtive, Move speed : +1",
		"If furtive, Range : +30",
	} 
	item.Price = 2700
}

func (item *Item) initAbundantPurse() {
	item.Description = "The abundant purse increases the amount of gold monsters drop."
	item.LevelUpDescription = []string {
		"On monsters, Money : +30%",
		"Prices in shop : -20%",
		"Regen : 1 Hp/100 golds picked",
	} 
	item.Price = 3000
}

func (item *Item) initTriforceLocket() {
	item.Description = "The triforce locket increases your 3 base stats."
	item.LevelUpDescription = []string {
		"Max Hp, Att, Def : +5, +2, +2",
		"Max Hp, Att, Def : +15, +6, +6",
		"Max Hp, Att, Def : +50, +20, +20",
	} 
	item.Price = 2200
}

func (item *Item) initGoldenClover() {
	item.Description = "The golden clover increases your critical rates."
	item.LevelUpDescription = []string {
		"Critical rate : +5%",
		"Critical rate : +10%",
		"On critical, Money : +5",
	} 
	item.Price = 2600
}

func (item *Item) initHealthPotion() {
	item.Description = "The health potion heals you depending on your missing health."
	item.LevelUpDescription = []string {
		"Hp : +30% of missing Hp",
		"",
		"",
	} 
	item.Price = 1500
	item.IsConsumable = true
}

func (item *Item) initBagPocket() {
	item.Description = "The bag pocket adds a new item slot in your bag."
	item.LevelUpDescription = []string {
		"Max bag size : +1",
		"",
		"",
	} 
	item.Price = 2000
	item.IsConsumable = true
}

//ADD INIT FUNCTIONS ABOVE

func (item *Item) Init(x, y int32, name ItemName, onSale bool) {
	item.X = x
	item.Y = y
	item.Level = 1
	item.Size = IBS
	item.OnSale = onSale
	item.Discount = 0
	item.IsConsumable = false
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
		case FIRE_HELMET:
			item.initFireHelmet()
			break
		case INVISIBLE_CAPE:
			item.initInvisibleCape()
			break
		case ABUNDANT_PURSE:
			item.initAbundantPurse()
			break
		case TRIFORCE_LOCKET:
			item.initTriforceLocket()
			break
		case GOLDEN_CLOVER:
			item.initGoldenClover()
			break
		case HEALTH_POTION:
			item.initHealthPotion()
			break
		case BAG_POCKET:
			item.initBagPocket()
			break
	}
	item.Name = name
	item.Selected = false
}

/* Effect */

func (item *Item) oneSettingPerLevel(sett1, sett2, sett3 Setting, _map *Map, prod int32) {
	add := prod == 1
	_map.CurrPlayer.Settings[sett1] = add
	if item.Level > 1 {//lvl2
		_map.CurrPlayer.Settings[sett2] = add
	}
	if item.Level > 2 {//lvl3
		_map.CurrPlayer.Settings[sett3] = add
	}
}

func (item *Item) applyEffectWaterBoots(_map *Map, prod int32) {
	item.oneSettingPerLevel(CAN_WALK_ON_WATER, SPEED_ON_WATER, REGEN_ON_WATER, _map, prod)
}

func (item *Item) applyEffectHeartOfSteel(_map *Map, prod int32) {
	var healthPoints int32 = 20
	if item.Level > 1 {//lvl2
		healthPoints += 30
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

func (item *Item) applyEffectFireHelmet(_map *Map, prod int32) {
	item.oneSettingPerLevel(LAVA_DEALS_HALF, LAVA_DEALS_NOTHING, RANGE_ON_LAVA, _map, prod)
}

func (item *Item) applyEffectInvisibleCape(_map *Map, prod int32) {
	add := prod == 1
	item.addFurtivity(_map, prod * 40)
	if item.Level > 1 {//lvl2
		_map.CurrPlayer.Settings[SPEED_IF_FURTIVE] = add
	}
	if item.Level > 2 {//lvl3
		_map.CurrPlayer.Settings[RANGE_IF_FURTIVE] = add
	}
}

func (item *Item) applyEffectAbundantPurse(_map *Map, prod int32) {
	item.oneSettingPerLevel(MONEY_DROP_BONUS, SHOP_DISCOUNT, REGEN_ON_MONEY, _map, prod)
}

func (item *Item) applyEffectTriforceLocket(_map *Map, prod int32) {
	item.addHealthPoints(_map, prod * 5)
	item.addAttack(_map, prod * 2)
	item.addDefense(_map, prod * 2)
	if item.Level > 1 {//lvl2
		item.addHealthPoints(_map, prod * 15)
		item.addAttack(_map, prod * 6)
		item.addDefense(_map, prod * 6)
	}
	if item.Level > 2 {//lvl3
		item.addHealthPoints(_map, prod * 50)
		item.addAttack(_map, prod * 20)
		item.addDefense(_map, prod * 20)
	}
}

func (item *Item) applyEffectGoldenClover(_map *Map, prod int32) {
	item.addCritRate(_map, prod * 5)
	if item.Level > 1 {//lvl2
		item.addCritRate(_map, prod * 10)
	}
	if item.Level > 2 {//lvl3
		_map.CurrPlayer.Settings[MONEY_CRIT_BONUS] = prod == 1
	}
}

func (item *Item) applyEffectHealthPotion(_map *Map, prod int32) {
	_map.CurrPlayer.HealMissingHp(30)
}

func (item *Item) applyEffectBagPocket(_map *Map, prod int32) {
	_map.CurrPlayer.ExtendBag()
}

//ADD APPLY FUNCTIONS ABOVE

func (item *Item) apply(_map *Map, value int32) {
	switch item.Name {
		case WATER_BOOTS:
			item.applyEffectWaterBoots(_map, value)
			break
		case HEART_OF_STEEL:
			item.applyEffectHeartOfSteel(_map, value)
			break
		case TURBO_REACTOR:
			item.applyEffectTurboReactor(_map, value)
			break
		case FIRE_HELMET:
			item.applyEffectFireHelmet(_map, value)
			break
		case INVISIBLE_CAPE:
			item.applyEffectInvisibleCape(_map, value)
			break
		case ABUNDANT_PURSE:
			item.applyEffectAbundantPurse(_map, value)
			break
		case TRIFORCE_LOCKET:
			item.applyEffectTriforceLocket(_map, value)
			break
		case GOLDEN_CLOVER:
			item.applyEffectGoldenClover(_map, value)
			break
		case HEALTH_POTION:
			item.applyEffectHealthPotion(_map, value)
			break
		case BAG_POCKET:
			item.applyEffectBagPocket(_map, value)
			break
	}
	if item.IsConsumable {
		_map.CurrPlayer.RemoveFromBag(*item)
	}
}

func (item *Item) ApplyEffect(_map *Map) {
	item.apply(_map, 1)
}

func (item *Item) RemoveEffect(_map *Map) {
	item.apply(_map, -1)
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

func (item *Item) drawFireHelmet() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.Orange)
}

func (item *Item) drawInvisibleCape() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.NewColor(25, 2, 93, 200))
}

func (item *Item) drawAbundantPurse() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.NewColor(133, 120, 35, 255))
}

func (item *Item) drawTriforceLocket() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.NewColor(100, 100, 84, 255))
}

func (item *Item) drawGoldenClover() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.NewColor(163, 222, 18, 255))
}

func (item *Item) drawHealthPotion() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.NewColor(255, 5, 5, 255))
}

func (item *Item) drawBagPocket() {
	rl.DrawRectangle(item.X+5, item.Y+5, item.Size-10, item.Size-10, rl.NewColor(220, 191, 32, 255))
}

//ADD DRAW FUNCTIONS ABOVE

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
		case FIRE_HELMET:
			item.drawFireHelmet()
			break
		case INVISIBLE_CAPE:
			item.drawInvisibleCape()
			break
		case ABUNDANT_PURSE:
			item.drawAbundantPurse()
			break
		case TRIFORCE_LOCKET:
			item.drawTriforceLocket()
			break
		case GOLDEN_CLOVER:
			item.drawGoldenClover()
			break
		case HEALTH_POTION:
			item.drawHealthPotion()
			break
		case BAG_POCKET:
			item.drawBagPocket()
			break
	}
}

//////

func (item *Item) GetLevelUpDescription(i int32) string {
	space := ""
	if i == 0 {
		space += " "
	}
	return "Lvl " + space + strconv.Itoa(int(i+1)) + ")   " + item.LevelUpDescription[i]
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

func (item *Item) addHealthPoints(_map *Map, value int32) {
	var hpPercentage float32 = float32(_map.CurrPlayer.Stats.Hp) / float32(_map.CurrPlayer.Stats.MaxHp)
	_map.CurrPlayer.Stats.MaxHp += value
	_map.CurrPlayer.Stats.Hp = int32(float32(_map.CurrPlayer.Stats.MaxHp) * hpPercentage)
}

func (item *Item) addAttack(_map *Map, value int32) {
	var attPercentage float32 = float32(_map.CurrPlayer.Stats.Att) / float32(_map.CurrPlayer.Stats.MaxAtt)
	_map.CurrPlayer.Stats.MaxAtt += value
	_map.CurrPlayer.Stats.Att = int32(float32(_map.CurrPlayer.Stats.MaxAtt) * attPercentage)
}

func (item *Item) addDefense(_map *Map, value int32) {
	var defPercentage float32 = float32(_map.CurrPlayer.Stats.Def) / float32(_map.CurrPlayer.Stats.MaxDef)
	_map.CurrPlayer.Stats.MaxDef += value
	_map.CurrPlayer.Stats.Def = int32(float32(_map.CurrPlayer.Stats.MaxDef) * defPercentage)
}

func (item *Item) addCritRate(_map *Map, value int32) {
	_map.CurrPlayer.Stats.CritRate += value
}

func (item *Item) addSpeed(_map *Map, value int32) {
	_map.CurrPlayer.Stats.MaxSpeed += value
}

func (item *Item) addRange(_map *Map, value int32) {
	_map.CurrPlayer.Stats.Range += value
}

func (item *Item) addFurtivity(_map *Map, value int32) {
	_map.CurrPlayer.Stats.Furtivity += value
}

func (item *Item) GetHitbox() rl.Rectangle {
	return rl.Rectangle{float32(item.X), float32(item.Y), float32(item.Size), float32(item.Size)}
}

func (item *Item) GetPrice() int32 {
	price := item.Price
	price -= (price * item.Discount / 100)
	return price
}

func (item *Item) DrawItemName(currX, currY int32) {
	name := string(item.Name)
	if item.OnSale {
		name += " - Price : " + strconv.Itoa(int(item.GetPrice())) + " gold "
		if item.Discount > 0 {
			name += "(-" + strconv.Itoa(int(item.Discount)) + "%)"
		}
	}
	rl.DrawText(name, currX, currY + 50, 23, rl.NewColor(144, 12, 63, 255))
	rl.DrawRectangle(currX, currY + 85, 140, 2, rl.DarkGray)
}

func (item *Item) DrawItemDescription(currX, currY int32) {
	var lineCount int32 = 0
		var maxLineChar int32 = 30
		var currLine string = ""
		var currWordLength int32
		var currLineLength int32
		for _, word := range strings.Split(item.Description, " ") {
			currWordLength = int32(len(word))
			currLineLength = int32(len(currLine))
			if currLineLength + currWordLength + 1 > maxLineChar {
				rl.DrawText(currLine, currX, currY + lineCount * 30, 20, rl.DarkGray)
				lineCount++
				currLine = word
			} else {
				if currLine != "" {
					currLine += " "
				}
				currLine += word
			}
		}
		rl.DrawText(currLine, currX, currY + lineCount * 30, 20, rl.DarkGray)
}

func (item *Item) DrawItemUpgrades(currX, currY int32) {
	var i int32
		var color rl.Color = rl.DarkGray
		for i = 0; i<IML; i++ {
			if i+1 <= item.Level {
				color = rl.NewColor(200, 32, 16, 255)
			} else {
				color = rl.DarkGray
			}
			rl.DrawText(item.GetLevelUpDescription(i), currX, currY + i * 30, 18, color)
		}
}