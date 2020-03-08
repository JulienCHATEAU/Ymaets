package class

import (
	// "fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

type Coin struct {
	X 					int32
	Y 					int32
	Radius			float32
	Value			 	int32
}

func (coin *Coin) Init(x, y int32) {
	coin.X = x
	coin.Y = y
	coin.Radius = 4
	coin.Value = r1.Int31() % 20 + 50
}

func (coin *Coin) Draw() {
	// Coin
	rl.DrawCircle(coin.X+1, coin.Y, coin.Radius, rl.NewColor(218, 161, 54, 255))
	rl.DrawCircle(coin.X, coin.Y, coin.Radius, rl.Gold)
	rl.DrawCircle(coin.X, coin.Y, coin.Radius - 3, rl.NewColor(218, 161, 54, 255))
	// rl.DrawText("$", coin.X - 1, coin.Y - 1, 1, rl.DarkGray)
}

func (coin *Coin) GetHitbox() (rl.Vector2, float32) {
	return rl.Vector2 {float32(coin.X), float32(coin.Y)}, coin.Radius
}