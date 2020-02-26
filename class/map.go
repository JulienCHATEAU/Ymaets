package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Map struct {
	Width 			int32
	Height 			int32
	BorderSize 	int32
	Players 		[]Player
	ShotsCount 	int32
	Shots 			[]Shot
} 

func (_map *Map) Init(windowSize int32) {
	_map.BorderSize = 5
	_map.Width = windowSize - _map.BorderSize
	_map.Height = windowSize - _map.BorderSize
	_map.Players = make([]Player, 1)
	_map.ShotsCount = 0
	_map.Shots = make([]Shot, 50)
	_map.Players[0] = Player {
		_map.Width / 2 + 40,
		_map.Height / 2 + 40,
		WEST,
		0,
		[4]int32{rl.KeyD, rl.KeyA, rl.KeyW, rl.KeyS},
		[4]int32{rl.KeyRight, rl.KeyLeft, rl.KeyUp, rl.KeyDown},
		rl.Red}
}

func (_map *Map) PlayerMove(index int, delta int32) {
	savedX := _map.Players[index].X
	savedY := _map.Players[index].Y
	if rl.IsKeyDown(_map.Players[index].Move_keys[0]) {
		_map.Players[index].X += delta;
	}
	if rl.IsKeyDown(_map.Players[index].Move_keys[1]) {
		_map.Players[index].X -= delta;
	}
	if rl.IsKeyDown(_map.Players[index].Move_keys[2]) {
		_map.Players[index].Y -= delta;
	}
	if rl.IsKeyDown(_map.Players[index].Move_keys[3]) {
		_map.Players[index].Y += delta;
	}
	if rl.IsKeyDown(_map.Players[index].Ori_keys[0]) {
		_map.Players[index].Ori = EAST;
	}
	if rl.IsKeyDown(_map.Players[index].Ori_keys[1]) {
		_map.Players[index].Ori = WEST;
	}
	if rl.IsKeyDown(_map.Players[index].Ori_keys[2]) {
		_map.Players[index].Ori = NORTH;
	}
	if rl.IsKeyDown(_map.Players[index].Ori_keys[3]) {
		_map.Players[index].Ori = SOUTH;
	}
	if _map.Players[index].X + PBS + PCH > _map.Width || _map.Players[index].X - PCH < _map.BorderSize || _map.Players[index].Y + PBS + PCH > _map.Height || _map.Players[index].Y - PCH < _map.BorderSize {
		_map.Players[index].X = savedX
		_map.Players[index].Y = savedY
	}
}

func (_map *Map) PlayerFire(index int) {
	if rl.IsKeyDown(rl.KeySpace) && _map.Players[index].FireCooldown == 0 {
		shot := _map.Players[index].GetShot()
		if int32(len(_map.Shots)) > _map.ShotsCount {
			_map.Shots[_map.ShotsCount] = shot
		} else {
			_map.Shots = append(_map.Shots, shot)
		}
		_map.ShotsCount++
		_map.Players[index].FireCooldown = PFC
	}
	_map.Players[index].ReduceCooldown()
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