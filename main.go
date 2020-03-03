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

func removeImpossibleOris(_maps map[ym.Coord]*ym.Map, currentMapCoord ym.Coord, oris []ym.Orientation) []ym.Orientation {
	var coord ym.Coord
	for _, ori := range oris {
		coord = ym.GetNextCoord(ori, currentMapCoord)
		if _, ok := _maps[coord]; ok {
			oris = ym.RemoveOri(oris, ori)
		}
	}
	return oris
}

func main() {

	var remainingMapCount int = 10
	var notCreatedYet int = 1
	var currentMapCoord ym.Coord = ym.Coord{0, 0}
	var _maps map[ym.Coord]*ym.Map = make(map[ym.Coord]*ym.Map)
	var _map *ym.Map = &ym.Map{}
	_map.Init(WINDOW_SIZE, []ym.Orientation {ym.NORTH})
	_map.CurrPlayer.Init(_map.Width - 50, _map.Height - 50, ym.NORTH)
	_map.InitBorders()
	_maps[currentMapCoord] = _map
	remainingMapCount--

	fmt.Println("Ymaets")
	rl.InitWindow(_maps[currentMapCoord].Width + MENU_SIZE, _maps[currentMapCoord].Height, "Ymaets")
	rl.HideCursor()
	rl.SetTargetFPS(60)
	var savedX int32
	var savedY int32
	var index int32 
	var changeMapOri ym.Orientation
	var newMapIndex ym.Coord
	var oppositeOri ym.Orientation
	var opening []ym.Orientation

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
			rl.ClearBackground(WINDOW_BCK)
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()

			if rl.IsMouseButtonPressed(rl.MouseRightButton) {
				fmt.Printf("(%d, %d)", mouseX, mouseY)
			}

			_maps[currentMapCoord].WallsDraw()

			for index = 0; index < _maps[currentMapCoord].ShotsCount; index++ {
				_maps[currentMapCoord].Shots[index].Draw()
				_maps[currentMapCoord].ShotMove(&index)
				_maps[currentMapCoord].ShotCheckMoveCollision(&index)
			}

			
			for index = 0; index < _maps[currentMapCoord].MonstersCount; index++ {
				savedX = _maps[currentMapCoord].Monsters[index].X
				savedY = _maps[currentMapCoord].Monsters[index].Y
				_maps[currentMapCoord].Monsters[index].Draw()
				_maps[currentMapCoord].MonsterMove(index)
				_maps[currentMapCoord].MonsterCheckMoveCollision(&index, savedX, savedY)
			}

			savedX = _maps[currentMapCoord].CurrPlayer.X
			savedY = _maps[currentMapCoord].CurrPlayer.Y
			savedOri := _maps[currentMapCoord].CurrPlayer.Ori
			_maps[currentMapCoord].PlayerOri(mouseX, mouseY)
			_maps[currentMapCoord].PlayerCheckOriCollision(savedOri)
			changeMapOri = _maps[currentMapCoord].PlayerMove()
			if changeMapOri == ym.NONE {
				_maps[currentMapCoord].PlayerCheckMoveCollision(savedX, savedY)
				_maps[currentMapCoord].PlayerFire()
				_maps[currentMapCoord].PlayerDraw()
			} else {
				newMapIndex = ym.GetNextCoord(changeMapOri, currentMapCoord)
				if _, ok := _maps[newMapIndex]; ok {
					_maps[newMapIndex].Update(*_maps[currentMapCoord], changeMapOri, WINDOW_SIZE)
					currentMapCoord = newMapIndex
				} else {
					var nextMap *ym.Map = &ym.Map{}
					oppositeOri = ym.GetOpositeOri(changeMapOri)
					var oris []ym.Orientation = []ym.Orientation {ym.NORTH, ym.SOUTH, ym.EAST, ym.WEST}
					oris = ym.RemoveOri(oris, oppositeOri)
					oris = removeImpossibleOris(_maps, newMapIndex, oris)
					opening = ym.GenerateOris(&remainingMapCount, &notCreatedYet, oppositeOri, oris, _maps, newMapIndex)
					nextMap.Init(WINDOW_SIZE, opening)
					nextMap.InitBorders()
					nextMap.MapChangeInit(*_maps[currentMapCoord], changeMapOri, WINDOW_SIZE, opening)
					_maps[newMapIndex] = nextMap
					currentMapCoord = newMapIndex
				}
			}
			_maps[currentMapCoord].CursorMove(mouseX, mouseY)
			_maps[currentMapCoord].CursorDraw()
			_maps[currentMapCoord].DrawMenu(MENU_SIZE)
		rl.EndDrawing()
	}
	fmt.Println(len(_maps))
	rl.CloseWindow()
}