package main

import (
	"fmt"
	"math/rand"
	"time"
	ym "Ymaets/class"
	util "Ymaets/util"
	"github.com/gen2brain/raylib-go/raylib"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

var stageMapCount int32 = 10
var WINDOW_SIZE int32 = 800
var MENU_SIZE int32 = 300
var MINI_MAP_SIZE int32 = 30
var MINI_MAP_PATH_LOW_EDGE int32 = 2
var MINI_MAP_PATH_HIGH_EDGE int32 = 6
var MINI_STAGE_START_X int32 = WINDOW_SIZE + MENU_SIZE / 2
var MINI_STAGE_START_Y int32 = WINDOW_SIZE / 2
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

func getNextMiniMapCoord(x, y int32, ori ym.Orientation) (int32, int32) {
	var nextX, nextY int32 = x, y
	switch ori {
	case ym.NORTH:
		nextY -= MINI_MAP_PATH_HIGH_EDGE + MINI_MAP_SIZE
		break

	case ym.SOUTH:
		nextY += MINI_MAP_PATH_HIGH_EDGE + MINI_MAP_SIZE
		break

	case ym.EAST:
		nextX += MINI_MAP_PATH_HIGH_EDGE + MINI_MAP_SIZE
		break

	case ym.WEST:
		nextX -= MINI_MAP_PATH_HIGH_EDGE + MINI_MAP_SIZE
		break
	}
	return nextX, nextY
}

func drawMiniMap(centerX, centerY int32, current bool, oppositeOri ym.Orientation) {
	var x, y int32 = centerX - MINI_MAP_SIZE/2, centerY - MINI_MAP_SIZE/2
	rl.DrawRectangle(x, y, MINI_MAP_SIZE, MINI_MAP_SIZE, rl.NewColor(179, 164, 151, 255))
	var color rl.Color = rl.NewColor(108, 89, 72, 255)
	if current {
		color = rl.Red
	}
	rl.DrawRectangleLinesEx(util.ToRectangle(x, y, MINI_MAP_SIZE, MINI_MAP_SIZE), 2, color)
}

func drawPath(currentMapX, currentMapY int32, ori ym.Orientation) {
	var color rl.Color = rl.NewColor(108, 89, 72, 255)
	switch ori {
	case ym.NORTH:
		rl.DrawRectangle(currentMapX - MINI_MAP_PATH_LOW_EDGE/2, currentMapY - MINI_MAP_SIZE/2 - MINI_MAP_PATH_HIGH_EDGE, MINI_MAP_PATH_LOW_EDGE, MINI_MAP_PATH_HIGH_EDGE, color)
		break

	case ym.SOUTH:
		rl.DrawRectangle(currentMapX - MINI_MAP_PATH_LOW_EDGE/2, currentMapY + MINI_MAP_SIZE/2, MINI_MAP_PATH_LOW_EDGE, MINI_MAP_PATH_HIGH_EDGE, color)
		break

	case ym.EAST:
		rl.DrawRectangle(currentMapX + MINI_MAP_SIZE/2, currentMapY - MINI_MAP_PATH_LOW_EDGE/2, MINI_MAP_PATH_HIGH_EDGE, MINI_MAP_PATH_LOW_EDGE, color)
		break

	case ym.WEST:
		rl.DrawRectangle(currentMapX - MINI_MAP_SIZE/2 - MINI_MAP_PATH_HIGH_EDGE, currentMapY - MINI_MAP_PATH_LOW_EDGE/2, MINI_MAP_PATH_HIGH_EDGE, MINI_MAP_PATH_LOW_EDGE, color)
		break
	}	
}

func drawMiniStage2(_maps map[ym.Coord]*ym.Map, drawn_maps map[ym.Coord]bool, playerMapCoord ym.Coord, currentMapCoord ym.Coord, current bool, centerX, centerY int32, oppositeOri ym.Orientation) {
	if _maps[currentMapCoord].Visited && !drawn_maps[currentMapCoord] {
		drawn_maps[currentMapCoord] = true
		drawMiniMap(centerX, centerY, playerMapCoord.X == currentMapCoord.X && playerMapCoord.Y == currentMapCoord.Y, oppositeOri)
		remainingMaps, _ := ym.RemoveOri(_maps[currentMapCoord].Opening, oppositeOri)
		var nextX, nextY int32
		var nextCoord ym.Coord
		for _, opening := range remainingMaps {
			nextCoord = ym.GetNextCoord(opening, currentMapCoord)
			oppositeOri = ym.GetOpositeOri(opening)
			nextX, nextY = getNextMiniMapCoord(centerX, centerY, opening)
			drawPath(centerX, centerY, opening)
			drawMiniStage2(_maps, drawn_maps, playerMapCoord, nextCoord, false, nextX, nextY, oppositeOri)
		}
	}
}

func drawMiniStage(_maps map[ym.Coord]*ym.Map, playerMapCoord ym.Coord) {
	var drawn_maps map[ym.Coord]bool = make(map[ym.Coord]bool)
	drawMiniStage2(_maps, drawn_maps, playerMapCoord, ym.Coord{0, 0}, true, MINI_STAGE_START_X, MINI_STAGE_START_Y, ym.NONE);
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
	_maps[currentMapCoord].Visited = true;
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
				_maps[newMapIndex].Visited = true
				currentMapCoord = newMapIndex
			}
			_maps[currentMapCoord].CursorMove(mouseX, mouseY)
			_maps[currentMapCoord].CursorDraw()
			_maps[currentMapCoord].DrawMenu(MENU_SIZE)
			drawMiniStage(_maps, currentMapCoord)
		rl.EndDrawing()
	}
	fmt.Println(len(_maps))
	rl.CloseWindow()
}