package main

import (
	"fmt"
	ym "Ymaets/class"
	"github.com/gen2brain/raylib-go/raylib"
)

var WINDOW_SIZE int32 = 800
var MENU_SIZE int32 = 300
var WINDOW_BCK rl.Color = rl.NewColor(245, 239, 220, 255) // Light Beige
var BORDER_COLOR rl.Color = rl.Gold

func getOpositeOri(ori ym.Orientation) ym.Orientation {
	if ori == ym.NORTH {
		return ym.SOUTH
	} else if ori == ym.SOUTH {
		return ym.NORTH
	} else if ori == ym.WEST {
		return ym.EAST
	} else if ori == ym.EAST {
		return ym.WEST
	} else {
		return ym.NONE
	}
}

func main() {

	var _maps []ym.Map = make([]ym.Map, 1)
	var currentMapIndex = 0
	_maps[currentMapIndex].Init(WINDOW_SIZE, []ym.Orientation{ym.NORTH, ym.SOUTH, ym.WEST, ym.EAST})
	_maps[currentMapIndex].CurrPlayer.Init(_maps[currentMapIndex].Width - 50, _maps[currentMapIndex].Height - 50, ym.NORTH)

	var mapsLink []map[ym.Orientation]int = make([]map[ym.Orientation]int, 1)
	mapsLink[0] = make(map[ym.Orientation]int)
	mapsLink[0][ym.NORTH] = -1
	mapsLink[0][ym.SOUTH] = -1
	mapsLink[0][ym.EAST] = -1
	mapsLink[0][ym.WEST] = -1

	fmt.Println("Ymaets")
	// screenW := rl.GetScreenWidth()
	// screenH := rl.GetScreenHeight()
	rl.InitWindow(_maps[currentMapIndex].Width + MENU_SIZE, _maps[currentMapIndex].Height, "Ymaets")
	rl.HideCursor()
	rl.SetTargetFPS(60)
	var savedX int32
	var savedY int32
	var index int32 
	var changeMapOri ym.Orientation

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
			rl.ClearBackground(WINDOW_BCK)
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()

			_maps[currentMapIndex].WallsDraw()

			for index = 0; index < _maps[currentMapIndex].ShotsCount; index++ {
				_maps[currentMapIndex].Shots[index].Draw()
				_maps[currentMapIndex].ShotMove(&index)
				_maps[currentMapIndex].ShotCheckMoveCollision(&index)
			}

			
			for index = 0; index < _maps[currentMapIndex].MonstersCount; index++ {
				savedX = _maps[currentMapIndex].Monsters[index].X
				savedY = _maps[currentMapIndex].Monsters[index].Y
				_maps[currentMapIndex].Monsters[index].Draw()
				_maps[currentMapIndex].MonsterMove(index)
				_maps[currentMapIndex].MonsterCheckMoveCollision(&index, savedX, savedY)
			}

			savedX = _maps[currentMapIndex].CurrPlayer.X
			savedY = _maps[currentMapIndex].CurrPlayer.Y
			savedOri := _maps[currentMapIndex].CurrPlayer.Ori
			_maps[currentMapIndex].PlayerOri(mouseX, mouseY)
			_maps[currentMapIndex].PlayerCheckOriCollision(savedOri)
			changeMapOri = _maps[currentMapIndex].PlayerMove()
			if changeMapOri == ym.NONE {
				_maps[currentMapIndex].PlayerCheckMoveCollision(savedX, savedY)
				_maps[currentMapIndex].PlayerFire()
				_maps[currentMapIndex].PlayerDraw()
			} else {
				var newMapIndex int = mapsLink[currentMapIndex][changeMapOri]
				if newMapIndex == -1 {
					newMapIndex = len(_maps)
					var _map ym.Map
					var opening []ym.Orientation = make([]ym.Orientation, 1)
					var oppositeOri ym.Orientation = getOpositeOri(changeMapOri)
					opening[0] = oppositeOri
					_map.Init(WINDOW_SIZE, opening)
					_map.MapChangeInit(_maps[currentMapIndex], changeMapOri, WINDOW_SIZE, opening)
					_maps = append(_maps, _map)

					mapsLink[currentMapIndex][changeMapOri] = newMapIndex
					var newMapLink map[ym.Orientation]int = make(map[ym.Orientation]int)
					newMapLink[ym.NORTH] = -1
					newMapLink[ym.SOUTH] = -1
					newMapLink[ym.EAST] = -1
					newMapLink[ym.WEST] = -1
					mapsLink = append(mapsLink, newMapLink)
					mapsLink[newMapIndex][oppositeOri] = currentMapIndex
					currentMapIndex = newMapIndex
					} else {
					_maps[newMapIndex].Update(_maps[currentMapIndex], changeMapOri, WINDOW_SIZE)
					currentMapIndex = newMapIndex
				}
			}
			_maps[currentMapIndex].CursorMove(mouseX, mouseY)
			_maps[currentMapIndex].CursorDraw()
			_maps[currentMapIndex].DrawMenu(MENU_SIZE)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}