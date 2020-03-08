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

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

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
	NextStage			Stairs
	CoinsCount		int32
	Coins 				[]Coin
	MonstersCount	int32
	Monsters 			[]Monster
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
	//Obstacles
	_map.Walls = append(_map.Walls, GeneratePossibleWalls(_map)...)
}

func (_map *Map) Init(coord Coord, windowSize int32, opening []Orientation) {
	_map.Visited = false
	_map.Coords = coord
	_map.BorderSize = 20
	_map.Width = windowSize
	_map.Height = windowSize
	_map.Opening = opening
	_map.Curs.Init()
	_map.CoinsCount = 0
	_map.Coins = make([]Coin, 0)
	_map.MonstersCount = 4
	_map.Monsters = make([]Monster, 50)
	_map.Monsters[0].Init(50, 50) 
	_map.Monsters[1].Init(150, 350) 
	_map.Monsters[2].Init(250, 50) 
	_map.Monsters[3].Init(100, 450)
	_map.NextStage.Init(-1, -1)
	_map.ShotsCount = 0
	_map.Shots = make([]Shot, 50)
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

func (_map *Map) DrawMenu(size int32, borderSize int32) {
	var textStarting int32 = 50 
	var textCount int32 = 0
	rl.DrawRectangle(_map.Width, 0, size, _map.Height, rl.NewColor(65, 87, 106, 255))
	rl.DrawRectangle(_map.Width + borderSize, borderSize, size - borderSize*2, _map.Height - borderSize*2, rl.RayWhite)
	rl.DrawText("HP : " + strconv.Itoa(int(_map.CurrPlayer.Hp)) + " / " + strconv.Itoa(int(_map.CurrPlayer.MaxHp)), _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount++
	rl.DrawText("Move speed : " + strconv.Itoa(int(_map.CurrPlayer.Speed)) + " / " + strconv.Itoa(int(_map.CurrPlayer.MaxSpeed)), _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount++
	rl.DrawText("Money : " + strconv.Itoa(int(_map.CurrPlayer.Money)) + " coin", _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
	textCount++
	textCount++
	rl.DrawText("Mini map : ", _map.Width + 30, textStarting + 50 * textCount, 20, rl.DarkGray)
}

func (_map *Map) MonsterMove(index int32) {
	if util.PointsDistance(_map.Monsters[index].X, _map.Monsters[index].Y, _map.CurrPlayer.X, _map.CurrPlayer.Y) <= _map.Monsters[index].AggroDist {
		var dx int32 = 0
		var dy int32 = 0
		if _map.Monsters[index].X < _map.CurrPlayer.X {
			dx = _map.Monsters[index].MoveSpeed
		} else {
			dx = -_map.Monsters[index].MoveSpeed
		}
		if _map.Monsters[index].Y < _map.CurrPlayer.Y {
			dy = _map.Monsters[index].MoveSpeed
		} else {
			dy = -_map.Monsters[index].MoveSpeed
		}
		_map.Monsters[index].X += dx
		_map.Monsters[index].Y += dy
	}	
}

func (_map *Map) MonsterCheckMoveCollision(index *int32, savedX, savedY int32) {
	center, radius := _map.Monsters[*index].GetHitbox()
	for _, wall := range _map.Walls {
		if !wall.Walkable {
			if rl.CheckCollisionCircleRec(center, radius, wall.GetHitbox()) {
				_map.Monsters[*index].X = savedX
				_map.Monsters[*index].Y = savedY
				return
			}
		}
	}

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
		_map.removeMonster(index)
		_map.CurrPlayer.TakeDamage(5)
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
			if !oneKeyPressed && _map.CurrPlayer.Speed < _map.CurrPlayer.MaxSpeed {
				_map.CurrPlayer.Speed += 1
			}
			oneKeyPressed = true
			*(dests[index]) += ops[index] * _map.CurrPlayer.Speed;
		}
	}
	if !oneKeyPressed {
		if _map.CurrPlayer.Speed > 0 {
			*(dests[lastKeyPressedIndex]) += ops[lastKeyPressedIndex] * _map.CurrPlayer.Speed;
			_map.CurrPlayer.Speed -= 1
		}
	}
	return _map.getChangeMapOri()
}

func (_map *Map) PlayerOri(mouseX, mouseY int32) {
	if mouseX > 0 && mouseX < _map.Width && mouseY > 0 && mouseY < _map.Height {
		_map.CurrPlayer.SetOriFromMouse(mouseX, mouseY)
	}
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
	if (rl.IsMouseButtonDown(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeySpace)) && _map.CurrPlayer.Animations.Values[FIRE_COOLDOWN] == 0 {
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
		if !wall.Walkable {
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

func (_map *Map) PlayerCheckMoveCollision(savedX, savedY int32) {
	hitbox := _map.CurrPlayer.GetHitbox()
	for index, _ := range _map.Walls {
		if rl.CheckCollisionRecs(hitbox, _map.Walls[index].GetHitbox()) {
			if !_map.Walls[index].Walkable {
				_map.CurrPlayer.X = savedX
				_map.CurrPlayer.Y = savedY
				return
			} else {
				if _map.Walls[index].Animations.Values[DAMAGE_DEALT] == 0 {
					_map.CurrPlayer.TakeDamage(_map.Walls[index].WalkDamage)
					_map.Walls[index].Animations.Values[DAMAGE_DEALT] = LDT
				}
			}
		}
	}
	var index int32
	var center rl.Vector2
	var radius float32

	for index = 0; index < _map.CoinsCount; index++ {
		center, radius = _map.Coins[index].GetHitbox()
		if rl.CheckCollisionCircleRec(center, radius, hitbox) {
			_map.CurrPlayer.Money += _map.Coins[index].Value
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

func (_map *Map) PlayerDraw() {
	_map.CurrPlayer.Draw()
}

func (_map *Map) WallsDraw() {
	for index := len(_map.Walls)-1; index >= 0; index-- {
		_map.Walls[index].Draw()
	}
}

func (_map *Map) CoinsDraw() {
	var index int32
	for index = 0; index < _map.CoinsCount; index++ {
		_map.Coins[index].Draw()
	}
}

func (_map *Map) ShotMove(index *int32) {
	if _map.Shots[*index].TravelDist >= _map.Shots[*index].Range {
		_map.removeShot(index)
		return
	}
	_map.Shots[*index].TravelDist += _map.Shots[*index].Speed
	switch _map.Shots[*index].Ori {
	case NORTH:
		_map.Shots[*index].Y -= _map.Shots[*index].Speed
		break;

	case SOUTH:
		_map.Shots[*index].Y += _map.Shots[*index].Speed
		break;

	case EAST:
		_map.Shots[*index].X += _map.Shots[*index].Speed		
		break;

	case WEST:
		_map.Shots[*index].X -= _map.Shots[*index].Speed		
		break;
	}
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

	for i = 0; i < _map.MonstersCount; i++ {
		center, radius = _map.Monsters[i].GetHitbox()
		if rl.CheckCollisionCircleRec(center, radius, hitbox) {
			_map.Monsters[i].TakeDamage(25)
			if _map.Monsters[i].Hp == 0 {
				coins := _map.Monsters[i].SpreadCoins()
				_map.CoinsCount += int32(len(coins))
				_map.Coins = append(_map.Coins, coins...)
				_map.removeMonster(&i)
			}
			_map.removeShot(index)
			return
		}
	}
}

func (_map *Map) aStar(walls []Wall) bool {
	rows := int(_map.Width / PBS)
	cols := int(_map.Height / PBS)
	fmt.Println(len(_map.Opening))
	aStar := astar.NewAStar(rows, cols)
	p2p := astar.NewPointToPoint()
	for i := 0; i<rows; i++ {
		for j := 0; j<cols; j++ {
			for _, wall := range walls {
				if !wall.Walkable && rl.CheckCollisionRecs(wall.GetHitbox(), rl.Rectangle {float32(i * rows), float32(j*rows), float32(PBS), float32(PBS)}) {
					aStar.FillTile(astar.Point{i, j}, -1)
					break
				}
			}
		}
	}

	if len(_map.Opening) == 1 {
		return true
	}

	var source, target []astar.Point
	if _map.Coords.X == 0 && _map.Coords.Y == 0 {
		for _, ori := range _map.Opening {
			source = OriToAstarCoord(ori, rows, cols)
			target = []astar.Point {astar.Point{rows-1, cols-1}}	
			if aStar.FindPath(p2p, source, target) == nil {
				return false
			}
		}
	}

	openingLength := len(_map.Opening)
	for k := 0; k<openingLength; k++ {
		for l := k+1; l<openingLength; l++ {
			source = OriToAstarCoord(_map.Opening[k], rows, cols)
			target = OriToAstarCoord(_map.Opening[l], rows, cols)
			if aStar.FindPath(p2p, source, target) == nil {
				return false
			}
		}
	}
	return true
}

func (_map *Map) AddStairs() {
	var found = false
	hitbox := _map.NextStage.GetHitbox()
  for !found {
		_map.NextStage.Init(r1.Int31() % 600 + 50, r1.Int31() % 600 + 50)
		for _, wall := range _map.Walls {
			if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
				break
			}
		}
		found = true
	}
}