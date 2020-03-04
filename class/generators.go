package class

import (
	// "fmt"
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

func GenerateWallWithOri(_map *Map, ori Orientation, x, y, lowEdge, highEdge, nextLowEdge int32) (Wall, int32, int32) {
	var wall Wall
	var nextX int32
	var nextY int32
	var currX int32
	var currY int32
	switch ori {
	case NORTH:
		currY = y - highEdge
		if currY < _map.BorderSize {
			highEdge -= (currY - _map.BorderSize)
			currY = _map.BorderSize
		}
		wall.InitWall(x, currY, lowEdge, highEdge, rl.Gray)
		nextX = x 
		nextY = currY
		break

	case EAST:
		if x + highEdge > _map.Width - _map.BorderSize {
			highEdge = _map.Width - _map.BorderSize - x
		}
		wall.InitWall(x, y, highEdge, lowEdge, rl.Gray)
		nextX = x + highEdge - nextLowEdge
		nextY = y
		break

	case SOUTH:
		if y + highEdge > _map.Height - _map.BorderSize {
			highEdge = _map.Height - _map.BorderSize - y
		}
		wall.InitWall(x, y, lowEdge, highEdge, rl.Gray)
		nextX = x
		nextY = y + highEdge - nextLowEdge
		break

	case WEST:
		currX = x - highEdge
		if currX < _map.BorderSize {
			highEdge -= (currX - _map.BorderSize)
			currX = _map.BorderSize
		}
		wall.InitWall(currX, y, highEdge, lowEdge, rl.Gray)
		nextX = currX
		nextY = y
		break
	}
	return wall, nextX, nextY
}

func GenerateBigWall(_map *Map, bigWallSurface, x, y, cornerCount int32, currentWalls []Wall, openningHitboxes []rl.Rectangle) []Wall {
	var wallCount int32 = cornerCount + 1
	var ori Orientation = NONE
	var oppositeOri Orientation
	var bigWall []Wall = make([]Wall, wallCount)
	var i int32
	var lowEdge int32 = r1.Int31() % 30 + 30
	var nextLowEdge int32
	var highEdge int32
	var remainingSurface = bigWallSurface
	var wallSurface int32
	var newX int32
	var newY int32
	var collision bool
	var orisss []Orientation
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
			bigWall[i], newX, newY = GenerateWallWithOri(_map, ori, x, y, lowEdge, highEdge, nextLowEdge)
			collision = rl.CheckCollisionRecs(_map.CurrPlayer.GetHitbox(), bigWall[i].GetHitbox())
			if collision {
				return bigWall[:i]
			}
			for _, currentWall := range currentWalls {
				collision = rl.CheckCollisionRecs(currentWall.GetHitbox(), bigWall[i].GetHitbox())
				if collision {
					return bigWall[:i]
				}
			}
			for _, openningHitbox := range openningHitboxes {
				collision = rl.CheckCollisionRecs(openningHitbox, bigWall[i].GetHitbox())
				if collision {
					return bigWall[:i]
				}
			}
			lowEdge = nextLowEdge
			x = newX
			y = newY
			oppositeOri = GetOpositeOri(ori)
		}
	}
	return bigWall
}

func GenerateWalls(_map *Map) []Wall {
	var freeSurface int32 = _map.GetFreeSurface()
	var obstacleSurfaceProportion int32 = (r1.Int31() % 5) + 5
	var obstacleSurface int32 = freeSurface * obstacleSurfaceProportion / 100

	var wallSurfaceProportion int32 = 100
	// var wallSurfaceProportion int32 = (r1.Int31() % 50)
	var wallSurface int32 = obstacleSurface * wallSurfaceProportion / 100
	// obstacleSurface -= wallSurface
	// var lavaSurfaceProportion int32 = (r1.Int31() % 50)
	// var lavaSurface int32 = obstacleSurface * lavaSurfaceProportion
	// obstacleSurface -= lavaSurface
	// var waterSurfaceProportion int32 = obstacleSurface
	var walls []Wall = make([]Wall, 0)
	var cornerCount int32 = r1.Int31() % 2 + 2
	var openningHitboxes []rl.Rectangle = _map.GetOpeningHitboxes()
	walls = append(walls, GenerateBigWall(_map, wallSurface, 400, 600, cornerCount, walls, openningHitboxes)...)
	cornerCount = r1.Int31() % 2 + 2
	walls = append(walls, GenerateBigWall(_map, wallSurface, 200, 200, cornerCount, walls, openningHitboxes)...)
	cornerCount = r1.Int31() % 2 + 2
	walls = append(walls, GenerateBigWall(_map, wallSurface, 500, 200, cornerCount, walls, openningHitboxes)...)
	return walls
}