package main

import (
	ym "Ymaets/class"
	"github.com/gen2brain/raylib-go/raylib"
)

var WINDOW_SIZE int32 = 800
var WINDOW_BCK rl.Color = rl.NewColor(245, 239, 220, 255) // Light Beige
var BORDER_COLOR rl.Color = rl.Gold

func main() {

	var _map ym.Map
	_map.Init(WINDOW_SIZE)

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