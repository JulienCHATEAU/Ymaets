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

var stageMapCount int32 = 7
var foundStairsMap bool = false

var MAP_SIZE int32 = 800
var MENU_SIZE int32 = 300
var MENU_BORDER_SIZE int32 = 10
var MENU_START_X = MAP_SIZE + MENU_BORDER_SIZE
var MINI_MAP_SIZE int32 = 30
var MINI_MAP_PATH_LOW_EDGE int32 = 2
var MINI_MAP_PATH_HIGH_EDGE int32 = 8
var MINI_STAGE_BORDER_SIZE int32 = 5
var MINI_STAGE_MARGIN int32 = 20
var MINI_STAGE_WIDTH int32 = MENU_SIZE - MENU_BORDER_SIZE * 2 - MINI_STAGE_MARGIN * 2
var MINI_STAGE_HEIGHT int32 = MINI_STAGE_WIDTH
var MINI_STAGE_MIN_X int32 = MENU_START_X + MINI_STAGE_MARGIN + MINI_STAGE_BORDER_SIZE
var MINI_STAGE_MAX_X int32 = MINI_STAGE_MIN_X + MINI_STAGE_WIDTH - MINI_STAGE_BORDER_SIZE * 2
var MINI_STAGE_MIN_Y int32 = MAP_SIZE / 2 - MINI_STAGE_HEIGHT / 2 + MINI_STAGE_BORDER_SIZE
var MINI_STAGE_MAX_Y int32 = MINI_STAGE_MIN_Y + MINI_STAGE_HEIGHT - MINI_STAGE_BORDER_SIZE * 2
var MINI_STAGE_START_X int32 = MINI_STAGE_MIN_X + MINI_STAGE_WIDTH / 2 - MINI_STAGE_BORDER_SIZE
var MINI_STAGE_START_Y int32 = MINI_STAGE_MIN_Y + MINI_STAGE_HEIGHT / 2 - MINI_STAGE_BORDER_SIZE
// var MINI_STAGE_START_X int32 = MAP_SIZE + MINI_MAP_SIZE + 10
var WINDOW_BCK rl.Color = rl.NewColor(245, 239, 220, 255) // Light Beige

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
	if *remainingMapCount > 0 {
		var oris []ym.Orientation = []ym.Orientation {ym.NORTH, ym.SOUTH, ym.EAST, ym.WEST}
		oris = removeImpossibleOris(_maps, currentMapCoord, oris)
		openings = ym.GenerateOris(remainingMapCount, oppositeOri, oris, _maps, currentMapCoord, r1.Int31() % 100 < deeperProba)
		openings = ym.ShuffleOris(openings)
	}
	fmt.Println(openings)
	var stairsProba int32 = 100 - deeperProba
	deeperProba -= (100 / stageMapCount)
	var _map *ym.Map = &ym.Map{}
	_map.CurrPlayer = player
	_map.Init(currentMapCoord, MAP_SIZE, openings)
	_map.InitBorders()
	if !foundStairsMap && r1.Int31() % 100 < stairsProba {
		foundStairsMap = true
		_map.AddStairs()
	} 
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
	var width int32 = MINI_MAP_SIZE
	var height int32 = MINI_MAP_SIZE
	var thick int32 = 2
	var color rl.Color = rl.NewColor(108, 89, 72, 255)
	if current {
		color = rl.Red
	}
	var outOfMiniMapBounds bool = true
	if x < MINI_STAGE_MIN_X && y < MINI_STAGE_MIN_Y {
		width -= (MINI_STAGE_MIN_X - x)
		x = MINI_STAGE_MIN_X
		height -= (MINI_STAGE_MIN_Y - y)
		y = MINI_STAGE_MIN_Y
		if height < 0 && width < 0 {
			height = 0
		}
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x, y, width - thick, height - thick, rl.NewColor(179, 164, 151, 255))
	} else if x < MINI_STAGE_MIN_X && y + height > MINI_STAGE_MAX_Y {
		width -= (MINI_STAGE_MIN_X - x)
		x = MINI_STAGE_MIN_X
		height = MINI_STAGE_MAX_Y - y
		if height < 0 && width < 0 {
			height = 0
		}
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x, y + thick, width - thick, height - thick, rl.NewColor(179, 164, 151, 255))
	} else if y + height > MINI_STAGE_MAX_Y && x + width > MINI_STAGE_MAX_X {
		height = MINI_STAGE_MAX_Y - y
		width = MINI_STAGE_MAX_X - x
		if height < 0 && width < 0 {
			height = 0
		}
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x + thick, y + thick, width - thick, height - thick, rl.NewColor(179, 164, 151, 255))
	} else if x + width > MINI_STAGE_MAX_X && y < MINI_STAGE_MIN_Y {
		height -= (MINI_STAGE_MIN_Y - y)
		y = MINI_STAGE_MIN_Y
		width = MINI_STAGE_MAX_X - x
		if height < 0 && width < 0 {
			height = 0
		}
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x + thick, y, width - thick, height - thick, rl.NewColor(179, 164, 151, 255))
	} else if x < MINI_STAGE_MIN_X {
		width -= (MINI_STAGE_MIN_X - x)
		x = MINI_STAGE_MIN_X
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x, y + thick, width - thick, height - thick*2, rl.NewColor(179, 164, 151, 255))
	} else if x + width > MINI_STAGE_MAX_X {
		width = MINI_STAGE_MAX_X - x
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x + thick, y + thick, width - thick, height - thick*2, rl.NewColor(179, 164, 151, 255))
	} else if y < MINI_STAGE_MIN_Y {
		height -= (MINI_STAGE_MIN_Y - y)
		y = MINI_STAGE_MIN_Y
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x + thick, y, width - thick*2, height - thick, rl.NewColor(179, 164, 151, 255))
	} else if y + height > MINI_STAGE_MAX_Y {
		height = MINI_STAGE_MAX_Y - y
		rl.DrawRectangle(x, y, width, height, color)
		rl.DrawRectangle(x + thick, y + thick, width - thick*2, height - thick, rl.NewColor(179, 164, 151, 255))
	} else {
		outOfMiniMapBounds = false
	}
	if !outOfMiniMapBounds {
		rl.DrawRectangle(x, y, width, height, rl.NewColor(179, 164, 151, 255))
		rl.DrawRectangleLinesEx(util.ToRectangle(x, y, width, height), thick, color)
	}
}

func drawPath(currentMapX, currentMapY int32, ori ym.Orientation) {
	var color rl.Color = rl.NewColor(98, 79, 62, 255)
	var x, y, width, height int32
	switch ori {
	case ym.NORTH:
		x = currentMapX - MINI_MAP_PATH_LOW_EDGE/2
		y = currentMapY - MINI_MAP_SIZE/2 - MINI_MAP_PATH_HIGH_EDGE
		width = MINI_MAP_PATH_LOW_EDGE
		height = MINI_MAP_PATH_HIGH_EDGE
		if y < MINI_STAGE_MIN_Y {
			height -= (MINI_STAGE_MIN_Y - y)
			y = MINI_STAGE_MIN_Y
		}
		break

	case ym.SOUTH:
		x = currentMapX - MINI_MAP_PATH_LOW_EDGE/2
		y = currentMapY + MINI_MAP_SIZE/2
		width = MINI_MAP_PATH_LOW_EDGE
		height = MINI_MAP_PATH_HIGH_EDGE
		if y + MINI_MAP_PATH_HIGH_EDGE > MINI_STAGE_MAX_Y {
			height = MINI_STAGE_MAX_Y - y
		}
		break

	case ym.EAST:
		x = currentMapX + MINI_MAP_SIZE/2
		y = currentMapY - MINI_MAP_PATH_LOW_EDGE/2
		width = MINI_MAP_PATH_HIGH_EDGE
		height = MINI_MAP_PATH_LOW_EDGE
		if x + MINI_MAP_PATH_HIGH_EDGE > MINI_STAGE_MAX_X {
			width = MINI_STAGE_MAX_X - x
		}
		break

	case ym.WEST:
		x = currentMapX - MINI_MAP_SIZE/2 - MINI_MAP_PATH_HIGH_EDGE
		y = currentMapY - MINI_MAP_PATH_LOW_EDGE/2
		width = MINI_MAP_PATH_HIGH_EDGE
		height = MINI_MAP_PATH_LOW_EDGE
		if x < MINI_STAGE_MIN_X {
			width -= (MINI_STAGE_MIN_X - x)
			x = MINI_STAGE_MIN_X
		}
		break
	}
	if x < MINI_STAGE_MIN_X || x > MINI_STAGE_MAX_X || y < MINI_STAGE_MIN_Y || y > MINI_STAGE_MAX_Y {
		return
	}
	rl.DrawRectangle(x, y, width, height,  color)
}

func drawMiniStage2(_maps map[ym.Coord]*ym.Map, drawn_maps map[ym.Coord]bool, playerMapCoord ym.Coord, currentMapCoord ym.Coord, centerX, centerY int32, oppositeOri ym.Orientation) bool {
	var current bool = playerMapCoord.X == currentMapCoord.X && playerMapCoord.Y == currentMapCoord.Y
	if current {
		if centerX - MINI_MAP_SIZE/2 < MINI_STAGE_MIN_X {
			return false
		}
		if centerX - MINI_MAP_SIZE/2 + MINI_MAP_SIZE > MINI_STAGE_MAX_X {
			return false
		}
		if centerY - MINI_MAP_SIZE/2 < MINI_STAGE_MIN_Y {
			return false
		}
		if centerY - MINI_MAP_SIZE/2 + MINI_MAP_SIZE > MINI_STAGE_MAX_Y {
			return false
		}
	}
	if _maps[currentMapCoord].Visited && !drawn_maps[currentMapCoord] {
		drawn_maps[currentMapCoord] = true
		drawMiniMap(centerX, centerY, current, oppositeOri)
		remainingMaps, _ := ym.RemoveOri(_maps[currentMapCoord].Opening, oppositeOri)
		var nextX, nextY int32
		var nextCoord ym.Coord
		for _, opening := range remainingMaps {
			nextCoord = ym.GetNextCoord(opening, currentMapCoord)
			oppositeOri = ym.GetOpositeOri(opening)
			nextX, nextY = getNextMiniMapCoord(centerX, centerY, opening)
			drawPath(centerX, centerY, opening)
			var ok bool = drawMiniStage2(_maps, drawn_maps, playerMapCoord, nextCoord, nextX, nextY, oppositeOri)
			if !ok {
				return false
			}
		}
	}
	return true
}

func drawMiniStage(_maps map[ym.Coord]*ym.Map, playerMapCoord ym.Coord, currentStage int32) {
	var drawn_maps map[ym.Coord]bool = make(map[ym.Coord]bool)
	borders := util.ToRectangle(MINI_STAGE_MIN_X - MINI_STAGE_BORDER_SIZE, MINI_STAGE_MIN_Y - MINI_STAGE_BORDER_SIZE, MINI_STAGE_WIDTH, MINI_STAGE_HEIGHT)
	rl.DrawRectangleLinesEx(borders, MINI_STAGE_BORDER_SIZE, rl.NewColor(47, 70, 91, 255))
	rl.DrawRectangle(MINI_STAGE_MIN_X, MINI_STAGE_MIN_Y, MINI_STAGE_WIDTH - MINI_STAGE_BORDER_SIZE*2, MINI_STAGE_HEIGHT - MINI_STAGE_BORDER_SIZE*2, rl.NewColor(215, 215, 215, 255))
	if !drawMiniStage2(_maps, drawn_maps, playerMapCoord, ym.Coord{0, 0}, MINI_STAGE_START_X, MINI_STAGE_START_Y, ym.NONE) {
		_maps[playerMapCoord].DrawMenu(MENU_SIZE, MENU_BORDER_SIZE, currentStage)
		drawn_maps = make(map[ym.Coord]bool)
		rl.DrawRectangleLinesEx(borders, MINI_STAGE_BORDER_SIZE, rl.NewColor(47, 70, 91, 255))
		rl.DrawRectangle(MINI_STAGE_MIN_X, MINI_STAGE_MIN_Y, MINI_STAGE_WIDTH - MINI_STAGE_BORDER_SIZE*2, MINI_STAGE_HEIGHT - MINI_STAGE_BORDER_SIZE*2, rl.NewColor(215, 215, 215, 255))
		drawMiniStage2(_maps, drawn_maps, playerMapCoord, playerMapCoord, MINI_STAGE_START_X, MINI_STAGE_START_Y, ym.NONE);
	}
}

func newStage(currentStage *int32, currentMapCoord *ym.Coord, player ym.Player) map[ym.Coord]*ym.Map {
	var remainingMapCount int32
	var _maps map[ym.Coord]*ym.Map
	foundStairsMap = false
	(*currentStage)++
	*currentMapCoord = ym.Coord{0, 0}
	for int32(len(_maps)) < stageMapCount - 1 {
		fmt.Println("START NEW STAGE GENERATION")
		remainingMapCount = stageMapCount
		_maps = initStage(make(map[ym.Coord]*ym.Map), player, 100, ym.NONE, *currentMapCoord, &remainingMapCount)
	}
	if !foundStairsMap {
		_maps[*currentMapCoord].AddStairs()
	}
	_maps[*currentMapCoord].Visited = true;
	return _maps
}

func main() {
	var currentStage int32 = 0
	var currentMapCoord ym.Coord
	var _maps map[ym.Coord]*ym.Map
	var player ym.Player
	player.Init(MAP_SIZE - 50, MAP_SIZE - 50, ym.NORTH)
	_maps = newStage(&currentStage, &currentMapCoord, player)
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
				fmt.Println(_maps[currentMapCoord].Coins)
				fmt.Println(_maps[currentMapCoord].CoinsCount)
			}

			_maps[currentMapCoord].CoinsDraw()
			_maps[currentMapCoord].WallsDraw()

			if _maps[currentMapCoord].NextStage.X != -1 {
				_maps[currentMapCoord].NextStage.Draw()
			}

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
				_maps[newMapIndex].Update(changeMapOri, MAP_SIZE)
				_maps[newMapIndex].Visited = true
				currentMapCoord = newMapIndex
			}
			_maps[currentMapCoord].CursorMove(mouseX, mouseY)
			_maps[currentMapCoord].CursorDraw()
			_maps[currentMapCoord].DrawMenu(MENU_SIZE, MENU_BORDER_SIZE, currentStage)
			drawMiniStage(_maps, currentMapCoord, currentStage)

			if _maps[currentMapCoord].IsPlayerOnStairs() {
				_maps[currentMapCoord].CurrPlayer.X = MAP_SIZE - 50
				_maps[currentMapCoord].CurrPlayer.Y = MAP_SIZE - 50
				_maps[currentMapCoord].CurrPlayer.Ori = ym.NORTH
				_maps = newStage(&currentStage, &currentMapCoord, _maps[currentMapCoord].CurrPlayer)
			}

		rl.EndDrawing()
	}
	fmt.Println(len(_maps))
	rl.CloseWindow()
}