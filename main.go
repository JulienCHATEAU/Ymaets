package main

import (
	ym "Ymaets/class"
	"github.com/gen2brain/raylib-go/raylib"
)

var WINDOW_SIZE int32 = 800
var WINDOW_BCK rl.Color = rl.NewColor(245, 239, 220, 255) // Light Beige
var PLAYERS_COUNT int32 = 1
var BORDER_COLOR rl.Color = rl.Gold

func main() {

	var _map ym.Map
	_map.BorderSize = 5
	_map.Width = WINDOW_SIZE - _map.BorderSize
	_map.Height = WINDOW_SIZE - _map.BorderSize
	_map.Players = make([]ym.Player, PLAYERS_COUNT)
	_map.ShotsCount = 0
	_map.Shots = make([]ym.Shot, 50)
	_map.Players[0] = ym.Player {
		_map.Width / 2 + 40,
		_map.Height / 2 + 40,
		ym.WEST,
		0,
		[4]int32{rl.KeyD, rl.KeyA, rl.KeyW, rl.KeyS},
		[4]int32{rl.KeyRight, rl.KeyLeft, rl.KeyUp, rl.KeyDown},
		rl.Red}

	rl.InitWindow(_map.Width, _map.Height, "Ymaets")
	rl.SetTargetFPS(60)
	
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
			rl.ClearBackground(WINDOW_BCK)
			rl.DrawRectangle(0, 0, _map.Width, _map.BorderSize, BORDER_COLOR)
			rl.DrawRectangle(0, 0, _map.BorderSize, _map.Width, BORDER_COLOR)
			rl.DrawRectangle(0, _map.Height - _map.BorderSize, _map.Width, _map.BorderSize, BORDER_COLOR)
			rl.DrawRectangle(_map.Width - _map.BorderSize, 0, _map.BorderSize, _map.Width, BORDER_COLOR)

			var index int32 
			for index = 0; index < _map.ShotsCount; index++ {
				_map.Shots[index].Draw()
				_map.ShotMove(&index)
			}
			for index, player := range _map.Players {
				_map.PlayerMove(index, 4)
				_map.PlayerFire(index)
				player.Draw()
			}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}