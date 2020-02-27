package class

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Orientation int32 
const (
	NORTH = iota + 1
	EAST
	SOUTH
	WEST
)

// Player body size
var PBS int32 = 28
// Player canon width
var PCW int32 = 6
// Player canon height
var PCH int32 = 8
// Player wheels size
var PWS float32 = 3
// Player shot width
var PSW int32 = 10
// Player shot height
var PSH int32 = 4
// Player shot speed
var PSS int32 = 5
// Player shot color
var PSC rl.Color = rl.Brown
// Player fire cooldown
var PFC int32 = 8
// Player move speed
var PMS int32 = 4

type Player struct {
	X 						int32
	Y 						int32
	Ori 					Orientation
	MoveSpeed		 	int32
	FireCooldown 	int32
	Move_keys 		[4]int32 // right, left, up, down
	Ori_keys 			[4]int32 // east, west, north, south
	Color 				rl.Color
}

func (player *Player) Init(x, y int32) {
		player.X = x
		player.Y = y
		player.Ori = WEST
		player.MoveSpeed = PMS
		player.FireCooldown = 0
		player.Move_keys = [4]int32{rl.KeyD, rl.KeyA, rl.KeyW, rl.KeyS}
		player.Ori_keys = [4]int32{rl.KeyRight, rl.KeyLeft, rl.KeyUp, rl.KeyDown}
		player.Color = rl.Blue
}

func (player *Player) ReduceCooldown() {
	if player.FireCooldown > 0 {
		player.FireCooldown--
	}
}

func (player *Player) GetCenter() (int32, int32) {
	return player.X + PBS / 2, player.Y + PBS / 2
}

func (player *Player) SetOriFromMouse(mouseX, mouseY int32) {
	centerX, centerY := player.GetCenter()
	if mouseX > mouseY && mouseX + mouseY < centerX + centerY {
		player.Ori = NORTH
	} else if mouseX < mouseY && mouseX + mouseY > centerX + centerY {
		player.Ori = SOUTH
	} else if mouseX > mouseY && mouseX + mouseY > centerX + centerY {
		player.Ori = EAST
	} else if mouseX < mouseY && mouseX + mouseY < centerX + centerY {
		player.Ori = WEST
	}
}

	// shot := Shot {50, 50, 4, 10, 3, NORTH, rl.Brown}
func (player *Player) GetShot() Shot {
	var shot Shot
	shot.Ori = player.Ori
	shot.Color = PSC
	shot.Speed = PSS
	shot.Width = PSH
	shot.Height = PSW
	switch player.Ori {
	case NORTH:
		shot.X = player.X + PBS/2-PCW/2
		shot.Y = player.Y - PCH - shot.Height
		break;

	case SOUTH:
		shot.X = player.X + PBS/2-PCW/2
		shot.Y = player.Y + PBS
		break;

	case EAST:
		shot.X = player.X + PBS
		shot.Y = player.Y + PBS/2-PCW/2
		break;

	case WEST:
		shot.X = player.X - PCH - shot.Width
		shot.Y = player.Y + PBS/2-PCW/2
		break;
	}
	return shot
}

func (player *Player) Draw() {
	// Body
	rl.DrawRectangle(player.X, player.Y, PBS, PBS, player.Color);
	// Canon
	switch player.Ori {
	case NORTH:
		rl.DrawRectangle(player.X + PBS/2-PCW/2, player.Y - PCH, PCW, PCH, rl.Black);
		break;

	case SOUTH:
		rl.DrawRectangle(player.X + PBS/2-PCW/2, player.Y + PBS, PCW, PCH, rl.Black);
		break;

	case EAST:
		rl.DrawRectangle(player.X + PBS, player.Y + PBS/2-PCW/2, PCH, PCW, rl.Black);
		break;

	case WEST:
		rl.DrawRectangle(player.X - PCH, player.Y + PBS/2-PCW/2, PCH, PCW, rl.Black);
		break;
	}
	// Wheels
	rl.DrawCircle(player.X, player.Y, PWS, rl.Black);
	rl.DrawCircle(player.X + PBS, player.Y, PWS, rl.Black);
	rl.DrawCircle(player.X, player.Y + PBS, PWS, rl.Black);
	rl.DrawCircle(player.X + PBS, player.Y + PBS, PWS, rl.Black);
}
