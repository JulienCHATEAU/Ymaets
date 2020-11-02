package class

import (
	// "fmt"
	util "Ymaets/util"
	rl "github.com/gen2brain/raylib-go/raylib"
	// "os"
)

const (
	CORE_ANIMATION = iota
	ANIMATION_COUNT
)

type AnimFunction func(remainingTime int32, spell *Spell)

type Spell struct {
	Name 						string
	Description 		string
	Cooldown				int32
	CurrCooldown 		int32
	SpellShot				Shot
	ActiveTime			int32
	CurrActiveTime	int32
	Animations			Timers
	Effect					func(spell *Spell, _map *Map)
	Draw						func(spell *Spell)
	AnimFunctions		map[int32]AnimFunction
}

func (spell *Spell) Init(cooldown, activeTime int32, name, desc string, spellFunction func(spell *Spell, _map *Map), drawFunction func(spell *Spell), animFunctions map[int32]AnimFunction) {
	spell.Name = name
	spell.Description = desc
	spell.Cooldown = cooldown
	spell.CurrCooldown = 0
	spell.ActiveTime = activeTime
	spell.CurrActiveTime = 0
	spell.Effect = spellFunction
	spell.Draw = drawFunction
	spell.Animations.Init(ANIMATION_COUNT)
	spell.AnimFunctions = animFunctions
}

func (spell *Spell) ApplyEffect(_map *Map) {
	spell.DecreaseCooldown(1)
	if spell.IsActive() {
		spell.Move()
		spell.Draw(spell)
		spell.HandleAnimations()
		spell.Effect(spell, _map)
		spell.DecreaseActiveTime(1)
	}
}

func (spell *Spell) HandleAnimation(notEnded []int32) {
	for _, anim := range notEnded {
		spell.AnimFunctions[anim](spell.Animations.Values[anim], spell)
	}
}

func (spell *Spell) HandleAnimations() {
	notEnded, _ := spell.Animations.Decrement()
	spell.HandleAnimation(notEnded)
}

func (spell *Spell) Move() {
	var shotEndLife = spell.SpellShot.Move()
	if shotEndLife {
		spell.SetInactive()
		//end life effect
		//...
	} 
}

func (spell *Spell) IsActive() bool {
		return spell.CurrActiveTime > 0 || spell.CurrActiveTime == ALWAYS_ACTIVE_AFTER_CAST
}

func (spell *Spell) SetInactive() {
	spell.Animations.Values[CORE_ANIMATION] = 0
	spell.SpellShot.TravelDist = 0
	spell.CurrActiveTime = 0
}

func (spell *Spell) IsAvailable() bool {
	return spell.CurrCooldown == 0
}

func (spell *Spell) GetHitbox() (rl.Vector2, float32) {
	return rl.Vector2 {float32(spell.SpellShot.X), float32(spell.SpellShot.Y)}, float32(spell.SpellShot.Width)
}

func (spell *Spell) Trigger(player *Player) {
	if spell.IsAvailable() {
		spell.SpellShot.Ori = player.Ori
		spell.SpellShot.Stats = player.GetStats()
		spell.Animations.Values[CORE_ANIMATION] = spell.ActiveTime
		spell.SetShot(player)
		spell.CurrActiveTime = spell.ActiveTime
		spell.CurrCooldown = spell.Cooldown
	}
}

func (spell *Spell) SetShot(player *Player) {
	spell.SpellShot.SetCoordFromPlayer(player)
	switch spell.SpellShot.Ori {
		case NORTH:
			spell.SpellShot.X += spell.SpellShot.Width/2
			break
	
		case SOUTH:
			spell.SpellShot.Y += spell.SpellShot.Height*2
			spell.SpellShot.X += spell.SpellShot.Width/2
			break
	
		case EAST:
			spell.SpellShot.X += spell.SpellShot.Width*2
			spell.SpellShot.Y += spell.SpellShot.Height/2
			break
	
		case WEST:
			spell.SpellShot.Y += spell.SpellShot.Height/2
			break
	}
}

func (spell *Spell) DecreaseActiveTime(time int32) {
	if spell.CurrActiveTime != ALWAYS_ACTIVE_AFTER_CAST {
		spell.CurrActiveTime -= time
		if spell.CurrActiveTime < 0 {
			spell.CurrActiveTime = 0
		}
	}
}

func (spell *Spell) DecreaseCooldown(time int32) {
		spell.CurrCooldown -= time
		if spell.CurrCooldown < 0 {
			spell.CurrCooldown = 0
		}
}

func (spell *Spell) HandleCollision(_map *Map) {
	spellCenter, spellRadius := spell.GetHitbox()
	for _, wall := range _map.Walls {
		if !wall.Crossable {
			if rl.CheckCollisionCircleRec(spellCenter, spellRadius, wall.GetHitbox()) {
				spell.SetInactive()
				return
			}
		}
	}
	var i int32
	var center rl.Vector2
	var radius float32
	for i = 0; i < _map.MonstersCount; i++ {
		center, radius = _map.Monsters[i].GetHitbox()
		if rl.CheckCollisionCircles(center, radius, spellCenter, spellRadius) {
			var damage int32 = util.GetDamage(spell.SpellShot, spell.SpellShot.BaseDamage, _map.Monsters[i])
			_map.Monsters[i].HandleDamageTaken(damage, _map.CurrPlayer)
			if _map.Monsters[i].IsDead() {
				_map.Monsters[i].SpreadLoots(_map)
				_map.CurrPlayer.AddExperience(_map.Monsters[i].GetExperience())
				_map.removeMonster(&i)
			}
			spell.SetInactive()
			return
		}
	}
}