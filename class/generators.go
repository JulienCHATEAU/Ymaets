package class

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"github.com/gen2brain/raylib-go/raylib"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)
var oris []Orientation = []Orientation {NORTH, SOUTH, EAST, WEST}

func GenerateOris(remainingMapCount *int32, oppositeOri Orientation, oriss []Orientation, _maps map[Coord]*Map, currentCoord Coord) []Orientation {
	var opening []Orientation
	var possibleAmount int32 = *remainingMapCount
	var toCreate int32
	var orissLength int32 = int32(len(oriss))
	if possibleAmount > 0 {
		if possibleAmount > orissLength {
			if orissLength > 0 {
				possibleAmount = orissLength
				toCreate = (r1.Int31() % possibleAmount) + 1
			} else {
				toCreate = 0
			}
		}
	} else {
		toCreate = 0
	}
	opening = make([]Orientation, 0)
	for _, ori := range oris {
		if value, ok := _maps[GetNextCoord(ori, currentCoord)]; ok {
			if ContainsOri(value.Opening, GetOpositeOri(ori)) {
				opening = append(opening, ori)
			} else {
				oriss, _ = RemoveOri(oriss, ori)
			}
		}
	}
	var start int32 = int32(len(opening))
	var ori Orientation = oppositeOri
	var trouve bool
	var i, j int32
	var count int32 = 0
	for i = start-1; i<toCreate; i++ {
		if count > 10 {
			break
		}
		ori = ChooseInOris(oriss)
		trouve = true
		for j = 0; j<=i; j++ {
			if ori == opening[j] {
				trouve = false
				break
			}
		}
		if trouve {
			*remainingMapCount--
			opening = append(opening, ori)
		} else {
			count++
			i--
		}
	}
	return opening
}

func GenerateWallWithOri(_map *Map, ori Orientation, x, y, lowEdge, highEdge, nextLowEdge int32) (int32, int32, int32, int32, int32, int32) {
	var nextX int32
	var nextY int32
	var tmp int32
	var currX int32 = x
	var currY int32 = y
	switch ori {
	case NORTH:
		currY = y - highEdge
		if currY < _map.BorderSize {
			fmt.Printf("NORTH : out of map %d, %d, %d, %d, %d, %d, %d\n", x, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
			highEdge -= (currY - _map.BorderSize)
			currY = _map.BorderSize
		}
		fmt.Printf("NORTH : %d, %d, %d, %d, %d, %d, %d\n", x, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
		nextX = currX 
		nextY = currY
		break

	case EAST:
		if currX + highEdge > _map.Width - _map.BorderSize {
			fmt.Printf("EAST : out of map %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
			highEdge = _map.Width - _map.BorderSize - currX
		}
		if currY + lowEdge > _map.Height - _map.BorderSize {
			fmt.Printf("EAST : out of map %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
			lowEdge = _map.Height - _map.BorderSize - currY
		}
		if currX < _map.BorderSize {
			fmt.Printf("EAST : out of map %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
			highEdge += (currX - _map.BorderSize)
			currX = _map.BorderSize
		}
		if currY < _map.BorderSize {
			fmt.Printf("EAST : out of map %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
			lowEdge += (currY - _map.BorderSize)
			currY = _map.BorderSize
		}
		fmt.Printf("EAST : %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
		nextX = currX + highEdge - nextLowEdge
		nextY = currY
		tmp = highEdge
		highEdge = lowEdge
		lowEdge = tmp
		break

	case SOUTH:
		if currY + highEdge > _map.Height - _map.BorderSize {
			fmt.Printf("SOUTH : out of map %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
			highEdge = _map.Height - _map.BorderSize - currY
		}
		fmt.Printf("SOUTH : %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
		nextX = currX
		nextY = currY + highEdge - nextLowEdge
		break

	case WEST:
		currX = currX - highEdge
		if currX < _map.BorderSize {
			fmt.Printf("WEST : out of map %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
			highEdge -= (currX - _map.BorderSize)
			currX = _map.BorderSize
		}
		fmt.Printf("WEST : %d, %d, %d, %d, %d, %d, %d\n", currX, currY, lowEdge, highEdge, _map.Width, _map.Height, _map.BorderSize)
		nextX = currX
		nextY = currY
		tmp = highEdge
		highEdge = lowEdge
		lowEdge = tmp
		break
	}
	return currX, currY, lowEdge, highEdge, nextX, nextY
}

func GenerateBigWall(_map *Map, bigWallSurface, x, y, cornerCount int32) []Wall {
	var wallCount int32 = cornerCount + 1
	var ori Orientation = NONE
	var oppositeOri Orientation
	var bigWall []Wall = make([]Wall, wallCount)
	var i int32
	var lowEdge int32 = r1.Int31() % 60 + 40
	var nextLowEdge int32
	var highEdge int32
	var remainingSurface = bigWallSurface
	var wallSurface int32
	var newX int32
	var newY int32
	var collision bool
	var orisss []Orientation
	var currX, currY, currLowEdge, currHighEdge int32
	for i = 0; i < wallCount; i++ {
		wallSurface = bigWallSurface * (r1.Int31() % 11 + (100 / wallCount) - 11) / 100
		if remainingSurface - wallSurface < 0 {
			wallSurface = remainingSurface
		}
		remainingSurface -= wallSurface
		if wallSurface > 0 {
			orisss, _ = RemoveOri(oris, oppositeOri)
			orisss, _ = RemoveOri(orisss, ori)
			ori = ChooseInOris(orisss)
			nextLowEdge = r1.Int31() % 30 + 30
			highEdge = wallSurface / lowEdge
			currX, currY, currLowEdge, currHighEdge, newX, newY = GenerateWallWithOri(_map, ori, x, y, lowEdge, highEdge, nextLowEdge)
			bigWall[i].InitWall(currX, currY, currLowEdge, currHighEdge, rl.Gray)
			collision = rl.CheckCollisionRecs(_map.CurrPlayer.GetHitbox(), bigWall[i].GetHitbox())
			if collision {
				return bigWall[:i]
			}
			lowEdge = nextLowEdge
			x = newX
			y = newY
			oppositeOri = GetOpositeOri(ori)
		}
	}
	return bigWall
}

func GeneratePylon(_map *Map, wallSurface, centerX, centerY, cornerCount int32) []Wall {
	var pylonLength int32 = cornerCount * 2 + 1
	var pylon []Wall = make([]Wall, pylonLength)
	var wEqualsH = int32(math.Sqrt(float64(wallSurface)))
	var wHDiff int32 = 10
	var centerRectWidth = wEqualsH + (((r1.Int31() % wHDiff)) * wEqualsH / 100)
	var centerRectHeight = wEqualsH * 2 - centerRectWidth
	var x int32 = centerX - centerRectWidth/2
	var y int32 = centerY - centerRectHeight/2
	currX, currY, currLowEdge, currHighEdge, _, _ := GenerateWallWithOri(_map, EAST, x, y, centerRectHeight, centerRectWidth, 0)
	pylon[0].InitWall(currX, currY, currLowEdge, currHighEdge, rl.Gray)
	var i int32
	var anglesLowEdge int32 = r1.Int31() % (centerRectWidth / 15) + 7
	var iangles int32
	for i = 1; i<cornerCount+1; i++ {
		iangles = i * anglesLowEdge
		currX, currY, currLowEdge, currHighEdge, _, _ = GenerateWallWithOri(_map, EAST, x + iangles, y - iangles, anglesLowEdge, centerRectWidth - iangles*2, 0)
		pylon[i].InitWall(currX, currY, currLowEdge, currHighEdge, rl.Gray)
		currX, currY, currLowEdge, currHighEdge, _, _ = GenerateWallWithOri(_map, EAST, x + iangles, y + centerRectHeight + iangles - anglesLowEdge, anglesLowEdge, centerRectWidth - iangles*2, 0)
		pylon[pylonLength - i].InitWall(currX, currY, currLowEdge, currHighEdge, rl.Gray)
	}
	return pylon
}


func GenerateLake(_map *Map, rectCount, maxWidth, minWidth, maxHeight, minHeight, startX, startY int32, lava bool) []Wall {
	var lake []Wall = make([]Wall, rectCount)
	var i int32
	var x int32 = startX
	var y int32 = startY
	var width int32
	var height int32
	var xDelta int32 = 30
	var currX, currY, currLowEdge, currHighEdge int32
	for i = 0; i<rectCount; i++ {
		width = (r1.Int31() % (maxWidth - minWidth)) + minWidth
		height = (r1.Int31() % (maxHeight - minHeight)) + minHeight
		currX, currY, currLowEdge, currHighEdge, _, _ = GenerateWallWithOri(_map, EAST, x, y, height, width, 0)
		if lava {
			lake[i].InitLava(currX, currY, currLowEdge, currHighEdge)
		} else {
			lake[i].InitWater(currX, currY, currLowEdge, currHighEdge)
		}
		x += (r1.Int31() % (xDelta*2)) - xDelta
		y += height
	}
	return lake
}

func GenerateWalls(_map *Map) []Wall {
	var freeSurface int32 = _map.GetFreeSurface()
	var obstacleSurfaceProportion int32 = (r1.Int31() % 5) + 5
	var obstacleSurface int32 = freeSurface * obstacleSurfaceProportion / 100

	var wallSurfaceProportion int32 = 80
	// var wallSurfaceProportion int32 = (r1.Int31() % 50)
	var wallSurface int32 = obstacleSurface * wallSurfaceProportion / 100
	// obstacleSurface -= wallSurface
	// var lavaSurfaceProportion int32 = (r1.Int31() % 50)
	// var lavaSurface int32 = obstacleSurface * lavaSurfaceProportion
	// obstacleSurface -= lavaSurface
	// var waterSurfaceProportion int32 = obstacleSurface
	var walls []Wall = make([]Wall, 0)
	var pylonCornerCount int32 = r1.Int31() % 4 + 1
	var bigWallCornerCount int32 = r1.Int31() % 2 + 1
	fmt.Print("wallSurface : ")
	fmt.Println(wallSurface)
	walls = append(walls, GenerateLake(_map, 5, 150, 100, 50, 30, 100, 200, false)...)
	walls = append(walls, GenerateLake(_map, 10, 150, 100, 40, 20, 400, 400, true)...)
	walls = append(walls, GeneratePylon(_map, wallSurface / 2, 400, 400, pylonCornerCount)...)
	// walls = append(walls, GeneratePylon(_map, wallSurface / 2, 50, 400, pylonCornerCount)...)
	bigWallCornerCount = r1.Int31() % 2 + 1
	walls = append(walls, GenerateBigWall(_map, wallSurface / 2, 400, 600, bigWallCornerCount)...)
	// walls = append(walls, GenerateBigWall(_map, wallSurface, 200, 200, cornerCount, walls, openningHitboxes)...)
	// cornerCount = r1.Int31() % 2 + 1
	// walls = append(walls, GenerateBigWall(_map, wallSurface, 500, 200, cornerCount, walls, openningHitboxes)...)
	return walls
}