package main

import (
	"fmt"
	ym "Ymaets/class"
	"github.com/gen2brain/raylib-go/raylib"
)

var WINDOW_SIZE int32 = 800
var WINDOW_BCK rl.Color = rl.NewColor(245, 239, 220, 255) // Light Beige
var BORDER_COLOR rl.Color = rl.Gold

func main() {

	var _map ym.Map
	_map.Init(WINDOW_SIZE)

	fmt.Println("Ymaets")
	// screenW := rl.GetScreenWidth()
	// screenH := rl.GetScreenHeight()
	rl.InitWindow(_map.Width, _map.Height, "Ymaets")
	// rl.InitWindow(int32(screenW), int32(screenH), "Ymaets")
	rl.HideCursor()
	rl.SetTargetFPS(60)
	
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
			rl.ClearBackground(WINDOW_BCK)
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()

			var index int32 
			for index = 0; index < _map.ShotsCount; index++ {
				_map.Shots[index].Draw()
				_map.ShotMove(index)
				_map.ShotCheckMoveCollision(&index)
			}

			var savedX int32
			var savedY int32
			for index = 0; index < _map.MonstersCount; index++ {
				savedX = _map.Monsters[index].X
				savedY = _map.Monsters[index].Y
				_map.Monsters[index].Draw()
				_map.MonsterMove(index)
				_map.MonsterCheckMoveCollision(&index, savedX, savedY)
			}

			_map.WallsDraw()

			savedX = _map.CurrPlayer.X
			savedY = _map.CurrPlayer.Y
			savedOri := _map.CurrPlayer.Ori
			_map.PlayerOri(mouseX, mouseY)
			_map.PlayerCheckOriCollision(savedOri)
			_map.PlayerMove()
			_map.PlayerCheckMoveCollision(savedX, savedY)
			_map.PlayerFire()
			_map.PlayerDraw()

			_map.CursorMove(mouseX, mouseY)
			_map.CursorDraw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}