package class

import (
	// "github.com/gen2brain/raylib-go/raylib"
	// "fmt"
	// "time"
)

const (
	CLASSIC = iota
	POLYVALENT
	ULTIMATE
	ELEMENT_SPELL_COUNT
	ALWAYS_ACTIVE_AFTER_CAST = 555555
)

type Element struct {
	Name 		string
	Spells 	[]Spell // classic, polyvalent, utlimate
}

func (elem *Element) Init(name string) {
	elem.Name = name
	elem.Spells = make([]Spell, ELEMENT_SPELL_COUNT)
}

var elements map[string]Element = InitElements()

func InitElements() map[string]Element {
	var e map[string]Element = make(map[string]Element)
	InitFire(e)
	return e
}

func GetDefaultElement(name string) Element {
	return elements[name]
}