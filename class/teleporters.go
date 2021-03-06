package class

import (
	// "fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"Ymaets/util"
)

// Teleporter body size
var SBS int32 = 40
var TELEPORTER_NOT_OK int32 = -100 

var S int32 = 8
var T int32 = 2

type TeleporterType string 
const (
	STAIRS = "Stairs"
	SHOP = "Shop"
	RETURN_STAGE = "ReturnStage"
)

type Teleporter struct {
	Type 				TeleporterType
	X 					int32
	Y 					int32
}

func (telep *Teleporter) Init(x, y int32, typee TeleporterType) {
	telep.Type = typee
	telep.X = x
	telep.Y = y
}

func (telep *Teleporter) DrawShop(opacity uint8) {
	green := rl.NewColor(48, 125, 74, 255)
	rl.DrawRectangle(telep.X, telep.Y, SBS, SBS, green)
	green.A = opacity
	var border int32 = 3
	width := SBS - border * 2
	rl.DrawRectangle(telep.X + border, telep.Y + border, width, width, rl.NewColor(200, 200, 200, opacity))
	rl.DrawCircle(telep.X + border + width/2, telep.Y + width - 8, 8, green)
	rl.DrawText("$", telep.X + border + width/2 - 3, telep.Y + width - 12, 5, rl.NewColor(255, 203, 0, opacity))
	rl.DrawRectangle(telep.X + border + width/2 - 3, telep.Y + width/2 - 4, 6, 6, green)
	rl.DrawCircle(telep.X + border + width/2 - 1, telep.Y + border + 3, 2, rl.NewColor(227, 123, 216, opacity))
	rl.DrawCircle(telep.X + border + width/2 + 2, telep.Y + 10, 2, rl.NewColor(227, 123, 216, opacity))
}

func (telep *Teleporter) Draw() {
	switch telep.Type {
	case STAIRS:
		rl.DrawRectangle(telep.X, telep.Y, SBS+2, SBS+2, rl.NewColor(50, 50, 50, 255))
		rl.DrawRectangle(telep.X, telep.Y + 2, S, SBS - 2, rl.LightGray)
		rl.DrawRectangle(telep.X + 1 * S, telep.Y + 2, T, SBS - 2, rl.NewColor(110, 110, 110, 255))

		rl.DrawRectangle(telep.X + 1 * S + 1 * T, telep.Y + 7, S, SBS - 7, rl.LightGray)
		rl.DrawRectangle(telep.X + 2 * S + 1 * T, telep.Y + 7, T, SBS - 7, rl.NewColor(110, 110, 110, 255))

		rl.DrawRectangle(telep.X + 2 * S + 2 * T, telep.Y + 12, S, SBS - 12, rl.LightGray)
		rl.DrawRectangle(telep.X + 3 * S + 2 * T, telep.Y + 12, T, SBS - 12, rl.NewColor(110, 110, 110, 255))

		rl.DrawRectangle(telep.X + 3 * S + 3 * T, telep.Y + 17, S, SBS - 17, rl.LightGray)
		rl.DrawRectangle(telep.X + 4 * S + 3 * T, telep.Y + 17, T, SBS - 17, rl.NewColor(110, 110, 110, 255))
		break

	case SHOP:
		telep.DrawShop(255)
		break

	case RETURN_STAGE:
		telep.DrawShop(180)
		rl.DrawRectangle(telep.X + 25, telep.Y + 10, 5, 16, rl.Black)
		rl.DrawRectangle(telep.X + 15, telep.Y + 20, 12, 6, rl.Black)
		v1, v2, v3 := rl.Vector2{float32(telep.X + 9), float32(telep.Y + 23)}, rl.Vector2{float32(telep.X + 15), float32(telep.Y + 16)}, rl.Vector2{float32(telep.X + 15), float32(telep.Y + 30)}
		rl.DrawTriangle(v1, v3, v2, rl.Black)
		break
	}
}

func (telep *Teleporter) IsOk() bool {
	return telep.X != TELEPORTER_NOT_OK
}

func (telep *Teleporter) GetHitbox() rl.Rectangle {
	return util.ToRectangle(telep.X, telep.Y, SBS+2, SBS+2)
}