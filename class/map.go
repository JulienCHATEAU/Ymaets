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
} 

func (_map *Map) Init(windowSize int32) {
	_map.BorderSize = 5
	_map.Width = windowSize - _map.BorderSize
	_map.Height = windowSize - _map.BorderSize
	_map.ShotsCount = 0
	_map.Shots = make([]Shot, 50)
	_map.CurrPlayer.Init(_map.Width / 2 + 40, _map.Height / 2 + 40)
	_map.Curs.Init()
}

func (_map *Map) PlayerMove() {
	savedX := _map.CurrPlayer.X
	savedY := _map.CurrPlayer.Y
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
	if _map.CurrPlayer.X + PBS + PCH > _map.Width || _map.CurrPlayer.X - PCH < _map.BorderSize || _map.CurrPlayer.Y + PBS + PCH > _map.Height || _map.CurrPlayer.Y - PCH < _map.BorderSize {
		_map.CurrPlayer.X = savedX
		_map.CurrPlayer.Y = savedY
	}
}

func (_map *Map) CursorMove(mouseX, mouseY int32) {
	_map.Curs.X = mouseX
	_map.Curs.Y = mouseY
}

func (_map *Map) CursorDraw() {
	_map.Curs.Draw()
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

func (_map *Map) PlayerDraw() {
	_map.CurrPlayer.Draw()
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
	var margin int32 = 20
	if _map.Shots[*index].X - margin > _map.Width || _map.Shots[*index].X + margin < _map.BorderSize || _map.Shots[*index].Y - margin > _map.Height || _map.Shots[*index].Y + margin < _map.BorderSize {
		_map.Shots[*index] = _map.Shots[_map.ShotsCount-1]
		_map.ShotsCount--
		(*index)--
	}
}