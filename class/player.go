package class

import (
	"github.com/gen2brain/raylib-go/raylib"
	util "Ymaets/util"
)


type Coord struct {
	X int32
	Y int32
}

func GetNextCoord(ori Orientation, coord Coord) Coord {
	var newCoord Coord = coord
	switch ori {
		case NORTH:
			newCoord.Y++
			break
		case SOUTH:
			newCoord.Y--
			break
		case WEST:
			newCoord.X--
			break
		case EAST:
			newCoord.X++
			break
	}
	return newCoord
}

type Orientation int32 
const (
	NONE = iota - 1
	NORTH
	SOUTH
	EAST
	WEST
)

func GetOpositeOri(ori Orientation) Orientation {
	var oppositeOri Orientation
	switch ori {
		case NORTH:
			oppositeOri = SOUTH
			break
		case SOUTH:
			oppositeOri = NORTH
			break
		case WEST:
			oppositeOri = EAST
			break
		case EAST:
			oppositeOri = WEST
			break

		default:
			oppositeOri = NONE
			break
	}
	return oppositeOri
}

func RemoveOri(oris []Orientation, ori Orientation) ([]Orientation, bool) {
	var oriss []Orientation = make([]Orientation, len(oris))
	for index, ori := range oris {
		oriss[index] = ori
	}
	for index, val := range oriss {
		if val == ori {
			oriss[index] = oriss[len(oriss)-1]
			return oriss[:len(oriss)-1], true
		}
	}
	return oriss, false
}

func ContainsOri(oris []Orientation, ori Orientation) bool {
	for _, val := range oris {
		if val == ori {
			return true
		}
	}
	return false
}

func ChooseInOris(oris []Orientation) Orientation {
	return oris[r1.Int() % len(oris)]
}

func ShuffleOris(oris []Orientation) []Orientation {
	r1.Shuffle(len(oris), func(i, j int) { oris[i], oris[j] = oris[j], oris[i] })
	return oris
}

// Player body size
var PBS int32 = 28
// Player canon width
var PCW int32 = 6
// Player canon height
var PCH int32 = 8
// Player shot width
var PSW int32 = 10
// Player shot height
var PSH int32 = 4
// Player shot speed (px/frame)
var PSS int32 = 5
// Player shot range (px)
var PSR int32 = 250
// Player shot color
var PSC rl.Color = rl.Red
// Player fire cooldown
var PFC int32 = 8
// Player move speed
var PMS int32 = 4
// Player health max
var PHM int32 = 100
// Player timers count
var PTC int32 = 2

type PlayerTimers int32
const (
	PLAYER_TAKE_DAMAGE = iota
	FIRE_COOLDOWN
)

type Player struct {
	X 						int32
	Y 						int32
	Ori 					Orientation
	Speed				 	int32
	MaxSpeed		 	int32
	Hp						int32
	MaxHp					int32
	Money					int32
	Move_keys 		[4]int32 // right, left, up, down
	Ori_keys 			[4]int32 // east, west, north, south
	Color 				rl.Color
	Animations		Timers
}

func (player *Player) Init(x, y int32, ori Orientation) {
		player.X = x
		player.Y = y
		player.Ori = ori
		player.Speed = 0
		player.MaxSpeed = PMS
		player.Hp = PHM
		player.MaxHp = PHM
		player.Money = 0
		player.Move_keys = [4]int32{rl.KeyD, rl.KeyA, rl.KeyW, rl.KeyS}
		player.Ori_keys = [4]int32{rl.KeyRight, rl.KeyLeft, rl.KeyUp, rl.KeyDown}
		player.Color = rl.Blue
		player.Animations.Init(PTC)
}

func (player *Player) GetHitbox() rl.Rectangle {
	var x int32 = 0
	var y int32 = 0
	width := PBS
	height := PBS
	switch player.Ori {
	case NORTH:
		x = player.X
		y = player.Y - PCH
		height += PCH
		break

	case SOUTH:
		x = player.X
		y = player.Y
		height += PCH
		break

	case EAST:
		x = player.X
		y = player.Y
		width += PCH
		break

	case WEST:
		x = player.X - PCH
		y = player.Y
		width += PCH
		break
	}
	return util.ToRectangle(x, y, width, height)
}

func (player *Player) GetCenter() (int32, int32) {
	return player.X + PBS / 2, player.Y + PBS / 2
}

func (player *Player) SetOriFromMouse(mouseX, mouseY int32) {
	centerX, centerY := player.GetCenter()
	if mouseX - mouseY > centerX - centerY && mouseX + mouseY < centerX + centerY {
		player.Ori = NORTH
	} else if mouseX - mouseY < centerX - centerY && mouseX + mouseY > centerX + centerY {
		player.Ori = SOUTH
	} else if mouseX - mouseY > centerX - centerY && mouseX + mouseY > centerX + centerY {
		player.Ori = EAST
	} else if mouseX - mouseY < centerX - centerY && mouseX + mouseY < centerX + centerY {
		player.Ori = WEST
	}
}

	// shot := Shot {50, 50, 4, 10, 3, NORTH, rl.Brown}
func (player *Player) GetShot() Shot {
	var shot Shot
	shot.Init(player.Ori, PSC, PSS, PSH, PSW, PSR)
	switch player.Ori {
	case NORTH:
		shot.X = player.X + PBS/2-PCW/2
		shot.Y = player.Y - PCH - shot.Height
		break

	case SOUTH:
		shot.X = player.X + PBS/2-PCW/2
		shot.Y = player.Y + PBS
		break

	case EAST:
		shot.X = player.X + PBS
		shot.Y = player.Y + PBS/2-PCW/2
		break

	case WEST:
		shot.X = player.X - PCH - shot.Width
		shot.Y = player.Y + PBS/2-PCW/2
		break
	}
	return shot
}

func (player *Player) TakeDamage(damage int32) {
	if damage > 0 {
		player.Hp -= damage
		if player.Hp - damage < 0 {
			player.Hp = 0
		}
		player.Animations.Values[PLAYER_TAKE_DAMAGE] = 5
	}
}

func (player *Player) HandleAnimation(notEnded []int32) {
	for i := 0; i<len(notEnded); i++ {
		switch notEnded[i] {
		case PLAYER_TAKE_DAMAGE:
			rl.DrawRectangleLinesEx(util.ToRectangle(player.X, player.Y, PBS, PBS), 3, rl.Red)
			break

		//ADD ANIMATION HANDLER HERE
		}
	}
}

func (player *Player) Draw() {
	// Body
	rl.DrawRectangle(player.X, player.Y, PBS, PBS, player.Color);
	// Canon
	switch player.Ori {
	case NORTH:
		rl.DrawRectangle(player.X + PBS/2-PCW/2, player.Y - PCH, PCW, PCH, rl.Black);
		break

	case SOUTH:
		rl.DrawRectangle(player.X + PBS/2-PCW/2, player.Y + PBS, PCW, PCH, rl.Black);
		break

	case EAST:
		rl.DrawRectangle(player.X + PBS, player.Y + PBS/2-PCW/2, PCH, PCW, rl.Black);
		break

	case WEST:
		rl.DrawRectangle(player.X - PCH, player.Y + PBS/2-PCW/2, PCH, PCW, rl.Black);
		break
	}
	// Health bar
	util.DrawHealthBar(player.Hp, player.MaxHp, player.X, player.Y - PCH, PBS)
	//Animations
	notEnded, _ := player.Animations.Decrement()
	player.HandleAnimation(notEnded)
}
