package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Map struct {
	Width 			int32
	Height 			int32
	BorderSize 	int32
	CurrPlayer 	Player
	Curs			 	Cursor
	ShotsCount 	int32
	Shots 			[]Shot
	Walls 			[]Wall
} 

func (_map *Map) Init(windowSize int32) {
	_map.BorderSize = 10
	_map.Width = windowSize - _map.BorderSize
	_map.Height = windowSize - _map.BorderSize
	_map.ShotsCount = 0
	_map.Shots = make([]Shot, 50)
	_map.CurrPlayer.Init(_map.Width / 2 + 40, _map.Height / 2 + 40)
	_map.Curs.Init()
	_map.Walls = make([]Wall, 7)
	//Borders
	_map.Walls[0].InitBorder(0, 0, _map.Width, _map.BorderSize)
	_map.Walls[1].InitBorder(0, 0, _map.BorderSize, _map.Width)
	_map.Walls[2].InitBorder(0, _map.Height - _map.BorderSize, _map.Width, _map.BorderSize)
	_map.Walls[3].InitBorder(_map.Width - _map.BorderSize, 0, _map.BorderSize, _map.Width)
	//Obstacles
	_map.Walls[4].Init(150, 150, 40, 30, rl.Gray)
	_map.Walls[5].Init(500, 170, 20, 50, rl.Gray)
	_map.Walls[6].Init(600, 540, 25, 45, rl.Gray)
}


func (_map *Map) CursorMove(mouseX, mouseY int32) {
	_map.Curs.X = mouseX
	_map.Curs.Y = mouseY
}

func (_map *Map) CursorDraw() {
	_map.Curs.Draw()
}

func (_map *Map) PlayerMove() {
	if rl.IsKeyDown(_map.CurrPlayer.Move_keys[0]) {
		_map.CurrPlayer.X += _map.CurrPlayer.MoveSpeed;
	}
	if rl.IsKeyDown(_map.CurrPlayer.Move_keys[1]) {
		_map.CurrPlayer.X -= _map.CurrPlayer.MoveSpeed;
	}
	if rl.IsKeyDown(_map.CurrPlayer.Move_keys[2]) {
		_map.CurrPlayer.Y -= _map.CurrPlayer.MoveSpeed;
	}
	if rl.IsKeyDown(_map.CurrPlayer.Move_keys[3]) {
		_map.CurrPlayer.Y += _map.CurrPlayer.MoveSpeed;
	}
}

func (_map *Map) PlayerOri(mouseX, mouseY int32) {
	_map.CurrPlayer.SetOriFromMouse(mouseX, mouseY)
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
	if rl.IsMouseButtonDown(rl.MouseLeftButton) && _map.CurrPlayer.FireCooldown == 0 {
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
}

func (_map *Map) PlayerDraw() {
	_map.CurrPlayer.Draw()
}

func (_map *Map) WallsDraw() {
	for _, wall := range _map.Walls {
		wall.Draw()
	}
}

func (_map *Map) ShotMove(index *int32) {
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
	_map.ShotsCheckCollision(index)
}

func (_map *Map) ShotsCheckCollision(index *int32) {
	hitbox := _map.Shots[*index].GetHitbox()
	for _, wall := range _map.Walls {
		if rl.CheckCollisionRecs(hitbox, wall.GetHitbox()) {
			_map.Shots[*index] = _map.Shots[_map.ShotsCount-1]
			_map.ShotsCount--
			(*index)--
			return
		}
	}
}