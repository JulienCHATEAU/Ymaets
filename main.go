package main

import (
	"fmt"
	"math/rand"
	"time"
	ym "Ymaets/class"
	"github.com/gen2brain/raylib-go/raylib"
)

type Coord struct {
	X int32
	Y int32
}

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

var WINDOW_SIZE int32 = 800
var MENU_SIZE int32 = 300
var WINDOW_BCK rl.Color = rl.NewColor(245, 239, 220, 255) // Light Beige
var BORDER_COLOR rl.Color = rl.Gold

func getOpositeOri(ori ym.Orientation) ym.Orientation {
	var oppositeOri ym.Orientation
	switch ori {
		case ym.NORTH:
			oppositeOri = ym.SOUTH
			break
		case ym.SOUTH:
			oppositeOri = ym.NORTH
			break
		case ym.WEST:
			oppositeOri = ym.EAST
			break
		case ym.EAST:
			oppositeOri = ym.WEST
			break

		default:
			oppositeOri = ym.NONE
			break
	}
	return oppositeOri
}

func getNextCoord(ori ym.Orientation, coord Coord) Coord {
	var newCoord Coord = coord
	switch ori {
		case ym.NORTH:
			newCoord.Y++
			break
		case ym.SOUTH:
			newCoord.Y--
			break
		case ym.WEST:
			newCoord.X--
			break
		case ym.EAST:
			newCoord.X++
			break
	}
	return newCoord
}

func generateOri(remainingMapCount, notCreatedYet *int, oppositeOri ym.Orientation, oris []ym.Orientation) []ym.Orientation {
	var opening []ym.Orientation
	possibleAmount := *remainingMapCount - *notCreatedYet
	var toCreate int
	orisLength := len(oris)
	if possibleAmount > 0 {
		if possibleAmount > orisLength {
			possibleAmount = orisLength
		}
		toCreate = (r1.Int() % possibleAmount) + 1
	} else {
		toCreate = 0
	}
	*remainingMapCount--
	*notCreatedYet += toCreate - 1
	opening = make([]ym.Orientation, toCreate+1)
	opening[0] = oppositeOri
	var ori ym.Orientation = oppositeOri
	var trouve bool
	for i := 1; i<toCreate+1; i++ {
		ori = oris[r1.Int() % orisLength]
		trouve = true
		for j := 0; j<i; j++ {
			if ori == opening[j] {
				trouve = false
				break
			}
		}
		if trouve {
			opening[i] = ori
		} else {
			i--
		}
	}
	return opening
}


func removeOrientation(oris []ym.Orientation, ori ym.Orientation) []ym.Orientation {
	for index, val := range oris {
		if val == ori {
			oris[index] = oris[len(oris)-1]
			return oris[:len(oris)-1]
		}
	}
	return oris
}

func removeImpossibleOris(_maps map[Coord]*ym.Map, currentMapCoord Coord, oris []ym.Orientation) []ym.Orientation {
	var coord Coord
	for _, ori := range oris {
		coord = getNextCoord(ori, currentMapCoord)
		if _, ok := _maps[coord]; ok {
			oris = removeOrientation(oris, ori)
		}
	}
	return oris
}

func main() {

	var remainingMapCount int = 10
	var notCreatedYet int = 1
	var currentMapCoord Coord = Coord{0, 0}
	var _maps map[Coord]*ym.Map = make(map[Coord]*ym.Map)
	var _map *ym.Map = &ym.Map{}
	_map.Init(WINDOW_SIZE, []ym.Orientation {ym.NORTH})
	_map.CurrPlayer.Init(_map.Width - 50, _map.Height - 50, ym.NORTH)
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
	var newMapIndex Coord
	var oppositeOri ym.Orientation
	var opening []ym.Orientation

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
			rl.ClearBackground(WINDOW_BCK)
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()

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
				newMapIndex = getNextCoord(changeMapOri, currentMapCoord)
				if _, ok := _maps[newMapIndex]; ok {
					_maps[newMapIndex].Update(*_maps[currentMapCoord], changeMapOri, WINDOW_SIZE)
					currentMapCoord = newMapIndex
				} else {
					var nextMap *ym.Map = &ym.Map{}
					oppositeOri = getOpositeOri(changeMapOri)
					var oris []ym.Orientation = []ym.Orientation {ym.NORTH, ym.SOUTH, ym.EAST, ym.WEST}
					fmt.Println(oris)
					oris = removeOrientation(oris, oppositeOri)
					fmt.Println(oris)
					oris = removeImpossibleOris(_maps, newMapIndex, oris)
					fmt.Println(oris)
					opening = generateOri(&remainingMapCount, &notCreatedYet, oppositeOri, oris)
					nextMap.Init(WINDOW_SIZE, opening)
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