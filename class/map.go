package class

import (
	"fmt"
	"Ymaets/util"
	"strconv"
	"time"
	"math/rand"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/nickdavies/go-astar/astar"
)

//Map opening size
var MOS int32 = 80
//Map opening hitbox edge
var MOHE int32 = 50
//Map time counters count
var MTCC int32 = 1

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

type MapTimeCounters int
const (
	MONSTERS_AGRESSIVITY = iota
)

type Map struct {
	Visited				bool
	Coords				Coord
	Width 				int32
	Height 				int32
	BorderSize 		int32
	CurrPlayer 		Player
	Curs			 		Cursor
	ShotsCount 		int32
	Opening 			[]Orientation
	Shots 				[]Shot
	Walls 				[]Wall
	Teleporters		map[TeleporterType]*Teleporter
	CoinsCount		int32
	Coins 				[]Coin
	MonstersCount	int32
	Monsters 			[]Monster
	ItemsCount 		int32
	Items 				[]Item
	TimeCounters	[]TimeCounter
} 

func (_map *Map) GetOpeningHitboxes() []rl.Rectangle {
	var hitboxes []rl.Rectangle = make([]rl.Rectangle, len(_map.Opening))
	for index, opening := range _map.Opening {
		switch opening {
		case NORTH:
			hitboxes[index] = rl.Rectangle {float32(_map.Width / 2 - MOS / 2), 0, float32(MOS), float32(MOHE)}
			break

		case WEST:
			hitboxes[index] = rl.Rectangle {0, float32(_map.Height / 2 - MOS / 2), float32(MOHE), float32(MOS)}
			break

		case SOUTH:
			hitboxes[index] = rl.Rectangle {float32(_map.Width / 2 - MOS / 2), float32(_map.Height - MOHE), float32(MOS), float32(MOHE)}
			break

		case EAST:
			hitboxes[index] = rl.Rectangle {float32(_map.Width - MOHE), float32(_map.Height / 2 - MOS / 2), float32(MOHE), float32(MOS)}
			break
		}
	}
	return hitboxes
}

func (_map *Map) InitBorders() {
	//Borders
	contentWallCount := 3
	borderCount := 4 + len(_map.Opening)
	_map.Walls = make([]Wall, borderCount + contentWallCount)
	_map.Walls[0].InitBorder(0, 0, _map.Width, _map.BorderSize)
	_map.Walls[1].InitBorder(0, 0, _map.BorderSize, _map.Height)
	_map.Walls[2].InitBorder(0, _map.Height - _map.BorderSize, _map.Width, _map.BorderSize)
	_map.Walls[3].InitBorder(_map.Width - _map.BorderSize, 0, _map.BorderSize, _map.Height)
	for i := 0; i<len(_map.Opening); i++ {
		switch _map.Opening[i] {
		case NORTH:
			_map.Walls[0].Width = _map.Walls[0].Width/2 - MOS/2
			_map.Walls[4+i].InitBorder(_map.Walls[0].X + _map.Walls[0].Width + MOS, 0, _map.Walls[0].Width, _map.BorderSize)
		break

		case WEST:
			_map.Walls[1].Height = _map.Walls[1].Height/2 - MOS/2
			_map.Walls[4+i].InitBorder(0, _map.Walls[1].Y + _map.Walls[1].Height + MOS, _map.BorderSize, _map.Walls[1].Height)
			break

		case SOUTH:
			_map.Walls[2].Width = _map.Walls[2].Width/2 - MOS/2
			_map.Walls[4+i].InitBorder(_map.Walls[2].X + _map.Walls[2].Width + MOS, _map.Walls[2].Y, _map.Walls[2].Width, _map.BorderSize)
			break

		case EAST:
			_map.Walls[3].Height = _map.Walls[3].Height/2 - MOS/2
			_map.Walls[4+i].InitBorder(_map.Walls[3].X, _map.Walls[3].Y + _map.Walls[3].Height + MOS, _map.BorderSize, _map.Walls[3].Height)
			break
		}
	}
}

func (_map *Map) InitShop(windowSize, borderSize int32) {
	_map.Visited = true
	_map.BorderSize = borderSize
	_map.Width = windowSize
	_map.Height = windowSize
	_map.Opening = []Orientation{}
	_map.Coins = make([]Coin, 0)
	_map.Curs.Init()
	_map.Monsters = make([]Monster, 50)
	_map.Teleporters = make(map[TeleporterType]*Teleporter)
	_map.Teleporters[STAIRS] = &Teleporter{}
	_map.Teleporters[STAIRS].Init(TELEPORTER_NOT_OK, TELEPORTER_NOT_OK, STAIRS)
	_map.Teleporters[SHOP] = &Teleporter{}
	_map.Teleporters[SHOP].Init(TELEPORTER_NOT_OK, TELEPORTER_NOT_OK, SHOP)
	_map.Teleporters[RETURN_STAGE] = &Teleporter{}
	_map.Teleporters[RETURN_STAGE].Init(windowSize / 2 - SBS / 2 + _map.BorderSize / 2, windowSize / 2 + 300, RETURN_STAGE)
	_map.Shots = make([]Shot, 50)
	_map.ItemsCount = 4
	_map.Items = make([]Item, 50)
	items := GetItems()
	rand, length := 0, 0
	var baseMargin int32 = 120
	var margin int32 = baseMargin
	size := IBS * 4 + margin * 3
	fmt.Println(size)
	fmt.Println(_map.Width)
	var index int32
	var typee ItemName
	for index = 0; index < _map.ItemsCount; index++ {
		if index == _map.ItemsCount-1 {
			if r1.Int() % 100 < 50 {
				typee = BAG_POCKET
			} else {
				typee = HEALTH_POTION
			}
		} else {
			length = len(items)
			rand = r1.Int() % length
			typee = items[rand]
			items[rand] = items[length-1]
			items = items[:length-1]
			fmt.Println(items)
		}
		fmt.Println(typee)
		_map.Items[index].Init(8 + (_map.Width / 2 - size / 2) + margin * index, 500, typee, true)
		if r1.Int31() % 100 < 5 {
			_map.Items[index].Discount += 15
		}
		margin = baseMargin + IBS
	}
	_map.TimeCounters = make([]TimeCounter, MTCC)
	_map.InitBorders()
	var shopBck Wall
	shopBck.InitWall(250, 40, 317, 140, rl.LightGray)
	_map.Walls = append(_map.Walls, shopBck)
	shopBck.InitWall(247, 37, 323, 146, rl.Gray)
	_map.Walls = append(_map.Walls, shopBck)
	shopBck.InitWall(40, 210, 720, 200, rl.RayWhite)
	_map.Walls = append(_map.Walls, shopBck)
	shopBck.InitWall(37, 207, 726, 206, rl.Gray)
	_map.Walls = append(_map.Walls, shopBck)
}

func (_map *Map) Init(coord Coord, windowSize, borderSize int32, opening []Orientation) {
	_map.Visited = false
	_map.Coords = coord
	_map.BorderSize = borderSize
	_map.Width = windowSize
	_map.Height = windowSize
	_map.Opening = opening
	_map.Curs.Init()
	_map.CoinsCount = 0
	_map.Coins = make([]Coin, 0)
	_map.MonstersCount = 0
	_map.ShotsCount = 0
	_map.Teleporters = make(map[TeleporterType]*Teleporter)
	_map.Teleporters[STAIRS] = &Teleporter{}
	_map.Teleporters[STAIRS].Init(TELEPORTER_NOT_OK, TELEPORTER_NOT_OK, STAIRS)
	_map.Teleporters[SHOP] = &Teleporter{}
	_map.Teleporters[SHOP].Init(TELEPORTER_NOT_OK, TELEPORTER_NOT_OK, SHOP)
	_map.Teleporters[RETURN_STAGE] = &Teleporter{}
	_map.Teleporters[RETURN_STAGE].Init(TELEPORTER_NOT_OK, TELEPORTER_NOT_OK, RETURN_STAGE)
	_map.Shots = make([]Shot, 50)
	_map.ItemsCount = 0
	_map.Items = make([]Item, 50)
	_map.TimeCounters = make([]TimeCounter, MTCC)
	_map.TimeCounters[MONSTERS_AGRESSIVITY].Init(false, 60)
}

func (_map *Map) ResetTimeCounters() {
	for index, _ := range _map.Monsters {
		_map.Monsters[index].Aggressive = false
	}
	for index, _ := range _map.TimeCounters {
		_map.TimeCounters[index].Reset()
		_map.TimeCounters[index].Off()
	}
}

func (_map *Map) HandleTimeCountersEnd(index int) {
	switch index {
	case MONSTERS_AGRESSIVITY:
		for i, _ := range _map.Monsters {
			_map.Monsters[i].Aggressive = true
		}
		break
	}
}

func (_map *Map) IncrementTimeCounters() {
	ended := false
	for index, _ := range _map.TimeCounters {
		ended = _map.TimeCounters[index].Increment()
		if ended {
			_map.HandleTimeCountersEnd(index)
		}
	}
}

func (_map *Map) Update(ori Orientation, windowSize int32) {
	_map.ShotsCount = 0
	switch ori {
	case NORTH:
		_map.CurrPlayer.Y = windowSize - PBS
		break

	case EAST:
		_map.CurrPlayer.X = 0
		break

	case SOUTH:
		_map.CurrPlayer.Y = 0
		break

	case WEST:
		_map.CurrPlayer.X = windowSize - PBS
		break
	}
}

func (_map *Map) GetFreeSurface() int32 {
	var freeSurface int32 = _map.Width * _map.Height
	freeSurface -= ((_map.Width - _map.BorderSize) * _map.BorderSize) * 4
	return freeSurface
}

func (_map *Map) BuffMonstersDepOnStage(currentStage int32) {
	for i, _ := range _map.Monsters {
		_map.Monsters[i].BuffDepOnStage(currentStage)
	}
}

func (_map *Map) DrawMenu(size, borderSize, currentStage int32) {
	var textStarting int32 = 50 
	var textCount int32 = 0
	rl.DrawRectangle(_map.Width, 0, size, _map.Height, rl.NewColor(65, 87, 106, 255))
	rl.DrawRectangle(_map.Width + borderSize, borderSize, size - borderSize*2, _map.Height - borderSize*2, rl.RayWhite)
	// rl.DrawFPS(_map.Width + size - 100, 20)
	// rl.DrawRectangle(_map.Width + 44, textStarting + 50 * textCount - 8, 150, 28, rl.Gray)
	// rl.DrawRectangle(_map.Width + 46, textStarting + 50 * textCount - 6, 150, 28, rl.LightGray)
	rl.DrawText("Stage n° " + strconv.Itoa(int(currentStage)), _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount++
	rl.DrawText("Lvl " + strconv.Itoa(int(_map.CurrPlayer.Level)) + "     Hp : " + strconv.Itoa(int(_map.CurrPlayer.Stats.Hp)) + " / " + strconv.Itoa(int(_map.CurrPlayer.Stats.MaxHp)), _map.Width + 30, 24 + textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount++
	rl.DrawRectangle(_map.Width + 26, textStarting + 50 * textCount, 248, 24, rl.Gray)
	util.DrawHealthBar(_map.CurrPlayer.Stats.Hp, _map.CurrPlayer.Stats.MaxHp, _map.Width + 30, 27 + textStarting + 50 * textCount, 240, 20)
	expStage := _map.CurrPlayer.GetCurrentExperienceStage()
	_map.CurrPlayer.Level--
	lowExpStage := _map.CurrPlayer.GetCurrentExperienceStage()
	_map.CurrPlayer.Level++
	diff := expStage - lowExpStage
	currDiff := _map.CurrPlayer.Experience - lowExpStage
	rl.DrawRectangle(_map.Width + 28, 35 + textStarting + 50 * textCount, 244, 5, rl.LightGray)
	util.DrawExperienceBar(currDiff, diff, _map.Width + 30, 35 + textStarting + 50 * textCount, 240, 5)
	textCount++
	textCount++
	// var coin Coin
	// coin.InitWithRadius(_map.Width + 138, textStarting + 50 * textCount + 8, 6)
	// coin.Draw()
	rl.DrawText("Mini map : ", _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount += 6
	rl.DrawText("Money : " + strconv.Itoa(int(_map.CurrPlayer.Money)) + " gold", _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount++
	rl.DrawText("Speed : " + strconv.Itoa(int(_map.CurrPlayer.Stats.Speed)) + " / " + strconv.Itoa(int(_map.CurrPlayer.Stats.MaxSpeed)), _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount++
	rl.DrawText("Bag : ", _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	util.ShowClassicKey(_map.Width + 95, textStarting + 50 * textCount, "E")
}

func (_map *Map) MonsterMove(index int32) bool {
	var isPlayerFurtive bool 
	if _map.Monsters[index].Aggressive && util.PointsDistance(_map.Monsters[index].X, _map.Monsters[index].Y, _map.CurrPlayer.X, _map.CurrPlayer.Y) <= _map.Monsters[index].AggroDist - float64(_map.CurrPlayer.Stats.Furtivity) {
		if !_map.Monsters[index].Settings[PLAYER_NEAR] {
			_map.Monsters[index].Settings[PLAYER_NEAR] = true
			_map.Monsters[index].Animations.Values[EXCLAMATION_POINT] = 350
		}
		_map.Monsters[index].Move(_map)
		_map.Monsters[index].Orient(_map)
		if _map.Monsters[index].HasCanon && _map.Monsters[index].Animations.Values[FIRE_COOLDOWN] == 0 {
			_map.Monsters[index].Fire(_map)
		}
		isPlayerFurtive = false
	} else {
		_map.Monsters[index].FindSeat(_map)
		_map.Monsters[index].Settings[PLAYER_NEAR] = false
		isPlayerFurtive = true
	}
	return isPlayerFurtive
}

func (_map *Map) MonsterCheckMoveCollision(index *int32, savedX, savedY int32) {
	center, radius := _map.Monsters[*index].GetHitbox()
	_map.Monsters[*index].Settings[IS_ON_LAVA] = false
	_map.Monsters[*index].Settings[COLLISION_ON_LAST_MOVE] = false
	for _, wall := range _map.Walls {
		if rl.CheckCollisionCircleRec(center, radius, wall.GetHitbox()) {
			if !wall.Walkable {
				_map.Monsters[*index].X = savedX
				_map.Monsters[*index].Y = savedY
				_map.Monsters[*index].Settings[COLLISION_ON_LAST_MOVE] = true
				continue
			} else {
				if wall.Type == Lava {
					_map.Monsters[*index].Settings[IS_ON_LAVA] = true
					if _map.Monsters[*index].Animations.Values[MONSTER_LAVA_DAMAGE] == 0 {
						_map.Monsters[*index].TakeDamage(wall.WalkDamage)
						if _map.Monsters[*index].IsDead() {
							_map.removeMonster(index)
						}
						_map.Monsters[*index].Animations.Values[MONSTER_LAVA_DAMAGE] = LDT
					}
				}
			}
		}
	}

	_map.Monsters[*index].HandleLavaExit()

	var monsterCenter rl.Vector2
	var monsterRadius float32
	var i int32
	for i = 0; i < _map.MonstersCount; i++ {
		if i != *index {
			monsterCenter, monsterRadius = _map.Monsters[i].GetHitbox()
			if rl.CheckCollisionCircles(center, radius, monsterCenter, monsterRadius) {
				_map.Monsters[*index].X = savedX
				_map.Monsters[*index].Y = savedY
				return
			}
		}
	}
	
	playerHitbox := _map.CurrPlayer.GetHitbox()
	if rl.CheckCollisionCircleRec(center, radius, playerHitbox) {
		_map.Monsters[*index].X = savedX
		_map.Monsters[*index].Y = savedY
		_map.Monsters[*index].PlayerCollision(_map)
		if _map.Monsters[*index].IsDead() {
			_map.removeMonster(index)
		}
		return
	}
}

func (_map *Map) CursorMove(mouseX, mouseY int32) {
	_map.Curs.X = mouseX
	_map.Curs.Y = mouseY
}

func (_map *Map) CursorDraw() {
	_map.Curs.Draw()
}

var lastKeyPressedIndex int
var ops [4]int32 = [4]int32 {1, -1, -1, 1}

func (_map *Map) getChangeMapOri() Orientation {
	var ori Orientation
	if _map.CurrPlayer.X < -PBS {
		ori = WEST
	} else if _map.CurrPlayer.X > _map.Width {
		ori = EAST
	} else if _map.CurrPlayer.Y < -PBS {
		ori = NORTH
	} else if _map.CurrPlayer.Y > _map.Height {
		ori = SOUTH
	} else {
		ori = NONE
	}
	return ori
}

func (_map *Map) PlayerMove() Orientation {
	var oneKeyPressed bool = false
	var dests [4]*int32 = [4]*int32 {&(_map.CurrPlayer.X), &(_map.CurrPlayer.X), &(_map.CurrPlayer.Y), &(_map.CurrPlayer.Y)}
	for index, key := range _map.CurrPlayer.Move_keys {
		if rl.IsKeyDown(key) {
			lastKeyPressedIndex = index
			if !oneKeyPressed && _map.CurrPlayer.Stats.Speed < _map.CurrPlayer.Stats.MaxSpeed {
				_map.CurrPlayer.Stats.Speed += 1
			}
			oneKeyPressed = true
			*(dests[index]) += ops[index] * _map.CurrPlayer.Stats.Speed;
		}
	}
	if !oneKeyPressed {
		if _map.CurrPlayer.Stats.Speed > 0 {
			*(dests[lastKeyPressedIndex]) += ops[lastKeyPressedIndex] * _map.CurrPlayer.Stats.Speed;
			_map.CurrPlayer.Stats.Speed -= 1
		}
	}
	return _map.getChangeMapOri()
}

func (_map *Map) PlayerOri(mouseX, mouseY int32) {
	// if mouseX > 0 && mouseX < _map.Width && mouseY > 0 && mouseY < _map.Height {
	// 	_map.CurrPlayer.SetOriFromMouse(mouseX, mouseY)
	// }
	if rl.IsKeyDown(_map.CurrPlayer.Ori_keys[0]) {
		_map.CurrPlayer.Ori = EAST;
	}
	if rl.IsKeyDown(_map.CurrPlayer.Ori_keys[1]) {
		_map.CurrPlayer.Ori = WEST;
	}
	if rl.IsKeyDown(_map.CurrPlayer.Ori_keys[2]) {
		_map.CurrPlayer.Ori = NORTH;
	}
	if rl.IsKeyDown(_map.CurrPlayer.Ori_keys[3]) {
		_map.CurrPlayer.Ori = SOUTH;
	}
}

func (_map *Map) PlayerFire() {
	// if (rl.IsMouseButtonDown(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeySpace)) && _map.CurrPlayer.Animations.Values[FIRE_COOLDOWN] == 0 {
	if (rl.IsKeyDown(rl.KeySpace)) && _map.CurrPlayer.Animations.Values[FIRE_COOLDOWN] == 0 {
		shot := _map.CurrPlayer.GetShot()
		_map.CurrPlayer.Animations.Values[FIRE_COOLDOWN] = PFC
		if int32(len(_map.Shots)) > _map.ShotsCount {
			_map.Shots[_map.ShotsCount] = shot
		} else {
			_map.Shots = append(_map.Shots, shot)
		}
		_map.ShotsCount++
	}
}

func (_map *Map) PlayerCheckOriCollision(savedOri Orientation) {
	hitbox := _map.CurrPlayer.GetHitbox()
	for _, wall := range _map.Walls {
		if !wall.Walkable && !_map.CurrPlayer.Settings[CAN_WALK_ON_WATER] {
			if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
				_map.CurrPlayer.Ori = savedOri
				return
			}
		}
	}
	var index int32
	var center rl.Vector2
	var radius float32
	var playerHitbox rl.Rectangle
	for index = 0; index < _map.MonstersCount; index++ {
		center, radius = _map.Monsters[index].GetHitbox()
		playerHitbox = _map.CurrPlayer.GetHitbox()
		if rl.CheckCollisionCircleRec(center, radius, playerHitbox) {
			_map.CurrPlayer.Ori = savedOri
			return
		}
	}
}

func (_map *Map) IsPlayerOnTeleporter(telep TeleporterType) bool {
	return rl.CheckCollisionRecs(_map.CurrPlayer.GetHitbox(), _map.Teleporters[telep].GetHitbox())
}

func (_map *Map) PlayerCheckMoveCollision(savedX, savedY int32) {
	hitbox := _map.CurrPlayer.GetHitbox()
	var index int32
	var center rl.Vector2
	var radius float32
	_map.CurrPlayer.Settings[IS_ON_WATER] = false
	_map.CurrPlayer.Settings[IS_ON_LAVA] = false
	for index, _ := range _map.Walls {
		if rl.CheckCollisionRecs(hitbox, _map.Walls[index].GetHitbox()) {
			if !_map.Walls[index].Walkable {
				if _map.Walls[index].Type == Water && _map.CurrPlayer.Settings[CAN_WALK_ON_WATER] {
					_map.CurrPlayer.Settings[IS_ON_WATER] = true
					break
				}
				_map.CurrPlayer.X = savedX
				_map.CurrPlayer.Y = savedY
				return
			} else {
				if _map.Walls[index].Type == Lava {
					_map.CurrPlayer.Settings[IS_ON_LAVA] = true
					if _map.CurrPlayer.Animations.Values[LAVA_DAMAGE] == 0 && !_map.CurrPlayer.Settings[LAVA_DEALS_NOTHING] { // Lava ticks damage
						if _map.CurrPlayer.Settings[LAVA_DEALS_HALF] {
							_map.CurrPlayer.TakeDamage(_map.Walls[index].WalkDamage / 2)
						} else {
							_map.CurrPlayer.TakeDamage(_map.Walls[index].WalkDamage)
						}
						_map.CurrPlayer.Animations.Values[LAVA_DAMAGE] = LDT
					}
				}
			}
		}
	}

	for index = 0; index < _map.CoinsCount; index++ {
		center, radius = _map.Coins[index].GetHitbox()
		if rl.CheckCollisionCircleRec(center, radius, hitbox) {
			currMoney := _map.CurrPlayer.Money
			_map.CurrPlayer.Money += _map.Coins[index].Value
			if _map.CurrPlayer.Settings[MONEY_DROP_BONUS] {
				_map.CurrPlayer.Money += _map.Coins[index].Value * 30 / 100
			}
			if _map.CurrPlayer.Settings[REGEN_ON_MONEY] {
				_map.CurrPlayer.Heal(_map.CurrPlayer.Money / 100 - currMoney / 100)
			}
			_map.removeCoin(&index)
		}
	}

	for index = 0; index < _map.MonstersCount; index++ {
		center, radius = _map.Monsters[index].GetHitbox()
		if rl.CheckCollisionCircleRec(center, radius, hitbox) {
			_map.CurrPlayer.X = savedX
			_map.CurrPlayer.Y = savedY
			_map.removeMonster(&index)
			_map.CurrPlayer.TakeDamage(5)
			return
		}
	}
}

func (_map *Map) PlayerHandleEffects() {
	_map.CurrPlayer.HandleWaterBootsSpeed()
	_map.CurrPlayer.HandleFireHelmetRange()
	_map.CurrPlayer.HandleInvisibleCapeSpeed()
	_map.CurrPlayer.HandleInvisibleCapeRange()
}

func (_map *Map) PlayerOnItem() {
	hitbox := _map.CurrPlayer.GetHitbox()
	var i int32
	for i = 0; i<_map.ItemsCount; i++ {
		canTakeItem := !_map.CurrPlayer.IsBagFull() || _map.Items[i].IsConsumable
		if rl.CheckCollisionRecs(_map.Items[i].GetHitbox(), hitbox) {
			if _map.Items[i].OnSale {
				_map.Items[i].DrawItemName(60, 180)
				_map.Items[i].DrawItemDescription(60, 300)
				_map.Items[i].DrawItemUpgrades(420, 300)
			}
			if !_map.CurrPlayer.HasItem(_map.Items[i]) {
				price := _map.Items[i].GetPrice()
				if (_map.Items[i].OnSale && _map.CurrPlayer.Money >= price) || !_map.Items[i].OnSale {
					if canTakeItem {
						util.ShowEnterKey(_map.Items[i].X + IBS + 13, _map.Items[i].Y + 5)
					}
					if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
						if canTakeItem {
							if _map.Items[i].OnSale {
								_map.CurrPlayer.UseMoney(price)
							}
							_map.CurrPlayer.AddInBag(_map.Items[i])
							_map.Items[i].ApplyEffect(_map)
							_map.removeItem(int32(i))
						}
					}
				}
			}
			break
		}
	}
}

func (_map *Map) TriggerSpells() {
	_map.CurrPlayer.TriggerSpells()
}

func (_map *Map) HandleSpells() {
	for i, _ := range _map.CurrPlayer.Elem.Spells {
		_map.CurrPlayer.Elem.Spells[i].ApplyEffect(_map)
	}
}

func (_map *Map) PlayerDraw() {
	_map.CurrPlayer.Draw()
}

func (_map *Map) WallsDraw() {
	for index := len(_map.Walls)-1; index >= 0; index-- {
		_map.Walls[index].Draw()
	}
}

func (_map *Map) ItemsDraw() {
	for index := _map.ItemsCount-1; index >= 0; index-- {
		_map.Items[index].Draw()
	}
}

func (_map *Map) CoinsDraw() {
	var index int32
	for index = 0; index < _map.CoinsCount; index++ {
		_map.Coins[index].Draw()
	}
}

func (_map *Map) ShotMove(index *int32) {
	var shouldBeRemoved bool = _map.Shots[*index].Move()
	if shouldBeRemoved {
		_map.removeShot(index)
	}
}

func (_map *Map) AddItem(item Item) {
	_map.Items[_map.ItemsCount] = item
	_map.ItemsCount++
}

func (_map *Map) removeItem(index int32) {
	for i := index; i<_map.ItemsCount-1; i++ {
		_map.Items[i] = _map.Items[i+1]
	}
	_map.ItemsCount--
}

func (_map *Map) removeCoin(index *int32) {
	_map.Coins[*index] = _map.Coins[_map.CoinsCount-1]
	_map.CoinsCount--
	_map.Coins = _map.Coins[:_map.CoinsCount]
	if *(index) > 0 {
		*(index)--
	}
}

func (_map *Map) removeMonster(index *int32) {
	_map.Monsters[*index] = _map.Monsters[_map.MonstersCount-1]
	_map.MonstersCount--
	if *(index) > 0 {
		*(index)--
	}
}

func (_map *Map) removeShot(index *int32) {
	_map.Shots[*index] = _map.Shots[_map.ShotsCount-1]
	_map.ShotsCount--
	if *(index) > 0 {
		*(index)--
	}
}

func (_map *Map) ShotCheckMoveCollision(index *int32) {
	hitbox := _map.Shots[*index].GetHitbox()
	for _, wall := range _map.Walls {
		if !wall.Crossable {
			if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
				_map.removeShot(index)
				return
			}
		}
	}
	var i int32
	var center rl.Vector2
	var radius float32

	if rl.CheckCollisionRecs(hitbox, _map.CurrPlayer.GetHitbox()) {
		if _map.Shots[*index].Owner != PLAYER {
			var damage int32 = util.GetDamage(_map.Shots[*index], _map.Shots[*index].BaseDamage, _map.CurrPlayer)
			fmt.Print("To player : ")
			fmt.Println(damage)
			_map.CurrPlayer.TakeDamage(damage)
		}
		_map.removeShot(index)
		return
	} else {
		for i = 0; i < _map.MonstersCount; i++ {
			center, radius = _map.Monsters[i].GetHitbox()
			if rl.CheckCollisionCircleRec(center, radius, hitbox) {
				if _map.Shots[*index].Owner != MONSTER {
					var damage int32 = util.GetDamage(_map.Shots[*index], _map.Shots[*index].BaseDamage, _map.Monsters[i])
					_map.Monsters[i].HandleDamageTaken(damage, _map.CurrPlayer)
					if _map.Monsters[i].IsDead() {
						_map.Monsters[i].SpreadLoots(_map)
						_map.CurrPlayer.AddExperience(_map.Monsters[i].GetExperience())
						_map.removeMonster(&i)
					}
				}
				_map.removeShot(index)
				return
			}
		}
	}
}

func (_map *Map) DiscountItems(value int32) {
	var i int32
	for i = 0; i<_map.ItemsCount; i++ {
		_map.Items[i].Discount += value
	}
}

func (_map *Map) aStar(walls []Wall) bool {
	var caseSize int32 = PBS + 1
	var rows, cols int32 = (_map.Width - _map.BorderSize*2) / caseSize, (_map.Height - _map.BorderSize/2) / caseSize
	var rowsint, colsint int = int(rows), int(cols)
	var i, j int32
	fmt.Println(len(_map.Opening))
	aStar := astar.NewAStar(int(rows), int(cols))
	p2p := astar.NewPointToPoint()
	var x string
	for i = 0; i<rows; i++ {
		for j = 0; j<cols; j++ {
			x = "."
			xx := _map.BorderSize + j * caseSize
			yy := _map.BorderSize + i * caseSize
			w := caseSize
			h := caseSize
			hitbox := rl.Rectangle {float32(xx), float32(yy), float32(w), float32(h)}
			for _, wall := range walls {
				if rl.CheckCollisionRecs(wall.GetHitbox(), hitbox) {
					aStar.FillTile(astar.Point{int(i), int(j)}, -1)
					x = "@"
					break
				}
			}
			fmt.Print(x)
		}
		fmt.Println()
	}

	if len(_map.Opening) == 1 {
		return true
	}

	var source, target []astar.Point

	if _map.Coords.X == 0 && _map.Coords.Y == 0 {
		// Players to each (0;0) map opening
		for _, ori := range _map.Opening {
			source = []astar.Point {astar.Point{rowsint-1, colsint-1}}
			target = OriToAstarCoord(ori, rowsint, colsint)
			if aStar.FindPath(p2p, source, target) == nil {
				return false
			}
		}

		// Player to (0;0) telep
		for _, telep := range _map.Teleporters {
			if telep.IsOk() {
				source = []astar.Point {astar.Point{int(_map.CurrPlayer.X * rows / _map.Width), int(_map.CurrPlayer.Y * cols / _map.Height)}}
				target = []astar.Point {astar.Point{int(telep.X * rows / _map.Width), int(telep.Y * cols / _map.Height)}}
				if aStar.FindPath(p2p, source, target) == nil {
					return false
				}
			}
		}
	}

	// Each opening to telporters
	openingLength := len(_map.Opening)
	for _, telep := range _map.Teleporters {
		if telep.IsOk() {
			var found bool = false
			for t := 0; t<openingLength; t++ {
				source = OriToAstarCoord(_map.Opening[t], rowsint, colsint)
				target = []astar.Point {astar.Point{int(telep.X * rows / _map.Width), int(telep.Y * cols / _map.Height)}}
				if aStar.FindPath(p2p, source, target) != nil {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	// Each opening to each opening
	for k := 0; k<openingLength; k++ {
		for l := k+1; l<openingLength; l++ {
			source = OriToAstarCoord(_map.Opening[k], rowsint, colsint)
			target = OriToAstarCoord(_map.Opening[l], rowsint, colsint)
			if aStar.FindPath(p2p, source, target) == nil {
				return false
			}
		}
	}
	return true
}

func (_map *Map) AddTeleporter(typee TeleporterType, walls []Wall) {
	var found = false
  for !found {
		found = true
		_map.Teleporters[typee].Init(r1.Int31() % 600 + 50, r1.Int31() % 600 + 50, typee)
		hitbox := _map.Teleporters[typee].GetHitbox()
		for _, wall := range walls {
			if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
				found = false
				break
			}
		}
	}
	fmt.Print(_map.Teleporters[typee].X)
	fmt.Print(" - ")
	fmt.Println(_map.Teleporters[typee].Y)
}

func (_map *Map) RemoveTelep(typee TeleporterType) {
	_map.Teleporters[typee].X = TELEPORTER_NOT_OK
	_map.Teleporters[typee].Y = TELEPORTER_NOT_OK
}

func (_map *Map) RemoveStairs() {
	_map.RemoveTelep(STAIRS)
}

func (_map *Map) RemoveShop() {
	_map.RemoveTelep(SHOP)
}

func (_map *Map) AddStairs(walls []Wall) {
	_map.AddTeleporter(STAIRS, walls)
	fmt.Println("\n\nSTAIRS\n\n")
}

func (_map *Map) AddShop(walls []Wall) {
	_map.AddTeleporter(SHOP, walls)
	fmt.Println("\n\nSHOP\n\n")
}

func (_map * Map) DrawShop() {
	rl.DrawText("$hop", 270, 50, 120, rl.NewColor(240, 149, 16, 255))
}