package class

import (
	"Ymaets/util"
	"github.com/gen2brain/raylib-go/raylib"
)

type EntityType int32
const (
	PLAYER = iota
	MONSTER
)

type Shot struct {
	X 					int32
	Y 					int32
	Width 			int32
	Height			int32
	Speed				int32
	Owner				EntityType
	BaseDamage 	int32
	Stats 			util.Stat
	Ori 				Orientation
	Range 			int32
	TravelDist 	int32
	Color 			rl.Color
}

func (shot *Shot) InitForSpell(color rl.Color, speed, width, height, rangee, baseDamage int32, owner EntityType) {
	shot.Width = width
	shot.Speed = speed
	shot.Height = height
	shot.Range = rangee
	shot.Owner = owner
	shot.Color = color
	shot.BaseDamage = baseDamage
}

func (shot *Shot) Init(ori Orientation, color rl.Color, speed, width, height, rangee, baseDamage int32, owner EntityType, stats util.Stat) {
	shot.Ori = ori
	shot.Color = color
	shot.Width = width
	shot.Speed = speed
	shot.Height = height
	shot.Range = rangee
	shot.Owner = owner
	shot.BaseDamage = baseDamage
	shot.Stats = stats
	shot.TravelDist = 0
}

func (shot *Shot) Draw() {
	if shot.Ori == NORTH || shot.Ori == SOUTH {
		rl.DrawRectangle(shot.X, shot.Y, shot.Width, shot.Height, shot.Color);
	} else {
		rl.DrawRectangle(shot.X, shot.Y, shot.Height, shot.Width, shot.Color);
	}
}

func (shot *Shot) Move() bool {
	if shot.Speed > 0 {
		if shot.TravelDist >= shot.Range {
			return true
		}
		shot.TravelDist += shot.Speed
		switch shot.Ori {
		case NORTH:
			shot.Y -= shot.Speed
			break;
	
		case SOUTH:
			shot.Y += shot.Speed
			break;
	
		case EAST:
			shot.X += shot.Speed		
			break;
	
		case WEST:
			shot.X -= shot.Speed		
			break;
		}
	}
	return false
}

func (shot *Shot) SetCoordFromPlayer(player *Player) {
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
}

func (shot *Shot) GetHitbox() rl.Rectangle {
	return rl.Rectangle{float32(shot.X), float32(shot.Y), float32(shot.Width), float32(shot.Height)}
}

func (shot Shot) GetStats() util.Stat {
	return shot.Stats
}