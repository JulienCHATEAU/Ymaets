package class

import (
	"time"
	"math/rand"
	"github.com/gen2brain/raylib-go/raylib"
)

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

type Map struct {
	Width 				int32
	Height 				int32
	BorderSize 		int32
	CurrPlayer 		Player
	Curs			 		Cursor
	ShotsCount 		int32
	Shots 				[]Shot
	Walls 				[]Wall
	MonstersCount	int32
	Monsters 			[]Monster
} 

func (_map *Map) Init(windowSize int32) {
	_map.BorderSize = 10
	_map.Width = windowSize - _map.BorderSize
	_map.Height = windowSize - _map.BorderSize
	_map.ShotsCount = 0
	_map.Shots = make([]Shot, 50)
	_map.CurrPlayer.Init(_map.Width - 50, _map.Height - 50)
	_map.Curs.Init()
	_map.MonstersCount = 1
	_map.Monsters = make([]Monster, 50)
	_map.Monsters[0].Init(50, 50) 
	//Borders
	_map.Walls = make([]Wall, 7)
	_map.Walls[0].InitBorder(0, 0, _map.Width, _map.BorderSize)
	_map.Walls[1].InitBorder(0, 0, _map.BorderSize, _map.Width)
	_map.Walls[2].InitBorder(0, _map.Height - _map.BorderSize, _map.Width, _map.BorderSize)
	_map.Walls[3].InitBorder(_map.Width - _map.BorderSize, 0, _map.BorderSize, _map.Width)
	//Obstacles
	_map.Walls[4].Init(150, 150, 40, 30, rl.Gray)
	_map.Walls[5].Init(500, 170, 20, 50, rl.Gray)
	_map.Walls[6].Init(600, 540, 25, 45, rl.Gray)
}

func (_map *Map) MonsterMove(index int32) {
	speed := _map.Monsters[index].MoveSpeed
	dx := (random.Int31() % speed) - speed/2
	dy := (random.Int31() % speed)
	_map.Monsters[index].X += dx
	_map.Monsters[index].Y += dy
}

func (_map *Map) MonsterCheckMoveCollision(index, savedX, savedY int32) {
	center, radius := _map.Monsters[index].GetHitbox()
	for _, wall := range _map.Walls {
		if rl.CheckCollisionCircleRec(center, radius, wall.GetHitbox()) {
			_map.Monsters[index].X = savedX
			_map.Monsters[index].Y = savedY
			return
		}
	}
	playerHitbox := _map.CurrPlayer.GetHitbox()
	if rl.CheckCollisionCircleRec(center, radius, playerHitbox) {
		_map.Monsters[index].X = savedX
		_map.Monsters[index].Y = savedY
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

func (_map *Map) PlayerMove() {
	var oneKeyPressed bool = false
	var dests [4]*int32 = [4]*int32 {&(_map.CurrPlayer.X), &(_map.CurrPlayer.X), &(_map.CurrPlayer.Y), &(_map.CurrPlayer.Y)}
	for index, key := range _map.CurrPlayer.Move_keys {
		if rl.IsKeyDown(key) {
			lastKeyPressedIndex = index
			if !oneKeyPressed && _map.CurrPlayer.MoveSpeed < _map.CurrPlayer.MaxSpeed {
				_map.CurrPlayer.MoveSpeed += 1
			}
			oneKeyPressed = true
			*(dests[index]) += ops[index] * _map.CurrPlayer.MoveSpeed;
		}
	}
	if !oneKeyPressed {
		if _map.CurrPlayer.MoveSpeed > 0 {
			*(dests[lastKeyPressedIndex]) += ops[lastKeyPressedIndex] * _map.CurrPlayer.MoveSpeed;
			_map.CurrPlayer.MoveSpeed -= 1
		}
	}
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
	if (rl.IsMouseButtonDown(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeySpace)) && _map.CurrPlayer.FireCooldown == 0 {
		shot := _map.CurrPlayer.GetShot()
		if int32(len(_map.Shots)) > _map.ShotsCount {
			_map.Shots[_map.ShotsCount] = shot
		} else {
			_map.Shots = append(_map.Shots, shot)
		}
		_map.ShotsCount++
		_map.CurrPlayer.FireCooldown = PFC
	}
	_map.CurrPlayer.ReduceCooldown()
}

func (_map *Map) PlayerCheckOriCollision(savedOri Orientation) {
	hitbox := _map.CurrPlayer.GetHitbox()
	for _, wall := range _map.Walls {
		if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
			_map.CurrPlayer.Ori = savedOri
			return
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
	for _, wall := range _map.Walls {
		if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
			_map.CurrPlayer.X = savedX
			_map.CurrPlayer.Y = savedY
			return
		}
	}
	var index int32
	var center rl.Vector2
	var radius float32
	for index = 0; index < _map.MonstersCount; index++ {
		center, radius = _map.Monsters[index].GetHitbox()
		if rl.CheckCollisionCircleRec(center, radius, hitbox) {
			_map.CurrPlayer.X = savedX
			_map.CurrPlayer.Y = savedY
			return
		}
	}
}

func (_map *Map) PlayerDraw() {
	_map.CurrPlayer.Draw()
}

func (_map *Map) WallsDraw() {
	for _, wall := range _map.Walls {
		wall.Draw()
	}
}

func (_map *Map) ShotMove(index int32) {
	switch _map.Shots[index].Ori {
	case NORTH:
		_map.Shots[index].Y -= _map.Shots[index].Speed
		break;

	case SOUTH:
		_map.Shots[index].Y += _map.Shots[index].Speed
		break;

	case EAST:
		_map.Shots[index].X += _map.Shots[index].Speed		
		break;

	case WEST:
		_map.Shots[index].X -= _map.Shots[index].Speed		
		break;
	}
}

func (_map *Map) ShotCheckMoveCollision(index *int32) {
	hitbox := _map.Shots[*index].GetHitbox()
	for _, wall := range _map.Walls {
		if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
			_map.Shots[*index] = _map.Shots[_map.ShotsCount-1]
			_map.ShotsCount--
			(*index)--
			return
		}
	}
	var i int32
	var center rl.Vector2
	var radius float32
	for i = 0; i < _map.MonstersCount; i++ {
		center, radius = _map.Monsters[i].GetHitbox()
		if rl.CheckCollisionCircleRec(center, radius, hitbox) {
			_map.Monsters[i].Color = rl.Orange
			_map.Shots[*index] = _map.Shots[_map.ShotsCount-1]
			_map.ShotsCount--
			(*index)--
			return
		}
	}
}