package class

import (
	"github.com/gen2brain/raylib-go/raylib"
	// "fmt"
	// "time"
)

func InitFire(e map[string]Element) {
	var element1Name = "Fire"
	var fireElem Element
	fireElem.Init(element1Name)

	/* Fire classic spells */
	var fireClassicSpells []Spell = make([]Spell, 1)

	// Spark
	var sparkAnimFunctions map[int32]AnimFunction = make(map[int32]AnimFunction)
	sparkAnimFunctions[CORE_ANIMATION] = func(remainingTime int32, spell *Spell) {
		shot := spell.SpellShot
		var delay int32 = 20
		if remainingTime % delay > delay/2 {
			DrawSparkAnimation(shot.X, shot.Y, shot.Width, shot.Height, 5, -3, 15, shot.Ori)
		} else {
			DrawSparkAnimation(shot.X, shot.Y, shot.Width, shot.Height, 1, 1, 30, shot.Ori)
		}
	}

	fireClassicSpells[0].SpellShot.InitForSpell(rl.Red, 5, 8, 8, 300, 15, PLAYER)
	fireClassicSpells[0].Init(150, ALWAYS_ACTIVE_AFTER_CAST, "Spark", "Throws a spark in the current direction", func(spell *Spell, _map *Map) {
		spell.HandleCollision(_map)
	}, func(spell *Spell) {
		shot := spell.SpellShot
		rl.DrawCircle(shot.X, shot.Y, float32(shot.Width), rl.Orange);
		rl.DrawCircle(shot.X+1, shot.Y+1, float32(shot.Width-2), rl.Red);
		rl.DrawCircle(shot.X+2, shot.Y+1, float32(shot.Width-3), rl.Black);
		DrawSparkBack(shot.X, shot.Y, shot.Width, shot.Height, shot.Ori)
	}, sparkAnimFunctions)
	

	/* Fire Polyvalent spells */
	var firePolyvalentSpells []Spell = make([]Spell, 1)

	// Fire Protection
	firePolyvalentSpells[0].Init(300, 150, "Fire protection", "Reduces damages received from fire sources", func(spell *Spell, _map *Map) {
		
	}, func(spell *Spell) {
		
	}, sparkAnimFunctions)


	/* Fire Ultimate spells */
	var fireUltimateSpells []Spell = make([]Spell, 1)

	// Heatwave
	fireUltimateSpells[0].Init(300, 150, "Heatwave", "Deals damages for 10 sec to all the ennemies present around you", func(spell *Spell, _map *Map) {

	}, func(spell *Spell) {
		
	}, sparkAnimFunctions)

	// Fire default spell deck
	fireElem.Spells[CLASSIC] = fireClassicSpells[0]
	fireElem.Spells[POLYVALENT] = firePolyvalentSpells[0]
	fireElem.Spells[ULTIMATE] = fireUltimateSpells[0]
	e[element1Name] = fireElem
}




/********/
/* UTIL */
/********/


func DrawSparkBack(x, y, width, height int32, ori Orientation) {
	switch ori {
	case NORTH:
		rl.DrawRectangle(x-width/2, y+height-1, width, 2, rl.Orange);
		rl.DrawRectangle(x-width/2+1, y+height+1, width-2, 2, rl.Orange);
		rl.DrawRectangle(x-width/2, y+height, 2, 4, rl.NewColor(230, 15, 10, 100));
		rl.DrawRectangle(x-width/2+1, y, 2, 3, rl.NewColor(230, 15, 10, 100));
		break;

	case SOUTH:
		rl.DrawRectangle(x-width/2, y-height-1, width, 2, rl.Orange);
		rl.DrawRectangle(x-width/2+1, y-height-3, width-2, 2, rl.Orange);
		rl.DrawRectangle(x-width/2, y-height-4, 2, 4, rl.NewColor(230, 15, 10, 100));
		rl.DrawRectangle(x-width/2+1, y-2, 2, 3, rl.NewColor(230, 15, 10, 100));
		break;

	case EAST:
		rl.DrawRectangle(x-width-1, y-height/2, 2, height, rl.Orange);
		rl.DrawRectangle(x-width-3, y-height/2+1, 2, height-2, rl.Orange);
		rl.DrawRectangle(x-width-4, y-height/2, 4, 2, rl.NewColor(230, 15, 10, 100));
		rl.DrawRectangle(x-2, y-height/2+1, 3, 2, rl.NewColor(230, 15, 10, 100));
		break;

	case WEST:
		rl.DrawRectangle(x+width-1, y-height/2, 2, height, rl.Orange);
		rl.DrawRectangle(x+width+1, y-height/2+1, 2, height-2, rl.Orange);
		rl.DrawRectangle(x+width, y-height/2, 4, 2, rl.NewColor(230, 15, 10, 100));
		rl.DrawRectangle(x, y-height/2+1, 3, 2, rl.NewColor(230, 15, 10, 100));
		break;
	}
}

func DrawSparkAnimation(x, y, width, height, sparkMargin, animMargin int32, dcolor uint8, ori Orientation) {
	switch ori {
	case NORTH:
		rl.DrawRectangle(x+animMargin, y+height+3+sparkMargin, 3, 5, rl.NewColor(220+dcolor, 15, 10, 100));
		break;

	case SOUTH:
		rl.DrawRectangle(x+animMargin, y-height-3-sparkMargin, 3, 5, rl.NewColor(220+dcolor, 15, 10, 100));
		break;

	case EAST:
		rl.DrawRectangle(x-width-3-sparkMargin, y+animMargin, 5, 3, rl.NewColor(220+dcolor, 15, 10, 100));
		break;
		
	case WEST:
		rl.DrawRectangle(x+width+3+sparkMargin, y+animMargin, 5, 3, rl.NewColor(220+dcolor, 15, 10, 100));
		break;
	}
}