package main

import (
	"fmt"
	"math/rand"
	"time"
	ym "Ymaets/class"
	"github.com/gen2brain/raylib-go/raylib"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

var stageMapCount int32 = 10
var WINDOW_SIZE int32 = 800
var MENU_SIZE int32 = 300
var WINDOW_BCK rl.Color = rl.NewColor(245, 239, 220, 255) // Light Beige
var BORDER_COLOR rl.Color = rl.Gold

func removeImpossibleOris(_maps map[ym.Coord]*ym.Map, currentMapCoord ym.Coord, oris []ym.Orientation) []ym.Orientation {
	var coord ym.Coord
	for _, ori := range oris {
		coord = ym.GetNextCoord(ori, currentMapCoord)
		if _, ok := _maps[coord]; ok {
			oris, _ = ym.RemoveOri(oris, ori)
		}
	}
	return oris
}

func initStage(_maps map[ym.Coord]*ym.Map, player ym.Player, deeperProba int32, ori ym.Orientation, currentMapCoord ym.Coord, remainingMapCount *int32) map[ym.Coord]*ym.Map {
	fmt.Println()
	fmt.Println(_maps)
	fmt.Println(currentMapCoord)
	fmt.Printf("remainingMapCount : %d\n", *remainingMapCount)
	var oppositeOri ym.Orientation = ym.GetOpositeOri(ori)
	var openings []ym.Orientation
	if *remainingMapCount > 0 && r1.Int31() % 100 < deeperProba {
		var oris []ym.Orientation = []ym.Orientation {ym.NORTH, ym.SOUTH, ym.EAST, ym.WEST}
		oris = removeImpossibleOris(_maps, currentMapCoord, oris)
		openings = ym.GenerateOris(remainingMapCount, oppositeOri, oris, _maps, currentMapCoord)
		openings = ym.ShuffleOris(openings)
	} else {
		openings = []ym.Orientation {oppositeOri}
	}
	fmt.Println(openings)
	deeperProba -= (100 / stageMapCount)
	var _map *ym.Map = &ym.Map{}
	_map.CurrPlayer = player
	_map.Init(currentMapCoord, WINDOW_SIZE, openings)
	_map.InitBorders()
	_maps[currentMapCoord] = _map
	var nextCoord ym.Coord
	remainingMapsToCreate, _ := ym.RemoveOri(openings, oppositeOri)
	for _, opening := range remainingMapsToCreate {
		nextCoord = ym.GetNextCoord(opening, currentMapCoord)
		if _, ok := _maps[nextCoord]; !ok {
			_maps = initStage(_maps, player, deeperProba, opening, nextCoord, remainingMapCount)
		}
	}
	return _maps
}

func main() {
	var remainingMapCount int32 = stageMapCount
	remainingMapCount--
	var currentMapCoord ym.Coord = ym.Coord{0, 0}
	var _maps map[ym.Coord]*ym.Map
	var player ym.Player
	player.Init(WINDOW_SIZE - 50, WINDOW_SIZE - 50, ym.NORTH)
	for int32(len(_maps)) < stageMapCount - 1 {
		remainingMapCount = stageMapCount
		_maps = initStage(make(map[ym.Coord]*ym.Map), player, 100, ym.NONE, currentMapCoord, &remainingMapCount)
	}
	fmt.Println(len(_maps))
	fmt.Println(_maps)
	
	fmt.Println("Ymaets")
	rl.InitWindow(_maps[currentMapCoord].Width + MENU_SIZE, _maps[currentMapCoord].Height, "Ymaets")
	rl.HideCursor()
	rl.SetTargetFPS(60)
	var savedX int32
	var savedY int32
	var index int32 
	var changeMapOri ym.Orientation
	var newMapIndex ym.Coord

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
			rl.ClearBackground(WINDOW_BCK)
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()

			if rl.IsMouseButtonPressed(rl.MouseRightButton) {
				fmt.Printf("(%d, %d)\n", mouseX, mouseY)
				fmt.Printf("Map coord : {%d, %d}\n", currentMapCoord.X, currentMapCoord.Y)
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
				_maps[newMapIndex].CurrPlayer = _maps[currentMapCoord].CurrPlayer
				_maps[newMapIndex].Update(changeMapOri, WINDOW_SIZE)
				currentMapCoord = newMapIndex
			}
			_maps[currentMapCoord].CursorMove(mouseX, mouseY)
			_maps[currentMapCoord].CursorDraw()
			_maps[currentMapCoord].DrawMenu(MENU_SIZE)
		rl.EndDrawing()
	}
	fmt.Println(len(_maps))
	rl.CloseWindow()
}