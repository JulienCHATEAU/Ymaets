package class

import (
	// "fmt"
	"strings"
	"github.com/gen2brain/raylib-go/raylib"
	util "Ymaets/util"
)

var tabs []string = []string {"Description", "Upgrades"}

type BagMenu struct {
	BorderSize		int32
	X 						int32
	Y 						int32
	Width					int32
	Heigth				int32
	CurrMap				*Map
	SelectedItem	int32
	SelectedTab	int32
	IsListFocused bool
}

func (bagMenu *BagMenu) Init(x, y, borderSize, width, heigth int32, _map *Map) {
	bagMenu.BorderSize = borderSize
	bagMenu.X = x
	bagMenu.Y = y
	bagMenu.Width = width
	bagMenu.Heigth = heigth
	bagMenu.SelectedItem = 0
	bagMenu.SelectedTab = 0
	bagMenu.IsListFocused = true
	bagMenu.CurrMap = _map
}

func (bagMenu *BagMenu) HandleFocus() {
	if rl.IsKeyPressed(rl.KeyUp) {
		bagMenu.IsListFocused = true
	} else if rl.IsKeyPressed(rl.KeyDown) {
		bagMenu.IsListFocused = false
	}
	if bagMenu.CurrMap.CurrPlayer.BagSize >= 1 {
		item := bagMenu.CurrMap.CurrPlayer.Bag[bagMenu.SelectedItem]
		if rl.IsKeyPressed(rl.KeyQ) {
			if bagMenu.SelectedItem == bagMenu.CurrMap.CurrPlayer.BagSize - 1 && bagMenu.SelectedItem > 0 {
				bagMenu.SelectedItem--
			}
			bagMenu.CurrMap.CurrPlayer.RemoveFromBag(item)
			item.RemoveEffect(bagMenu.CurrMap)
			item.X = bagMenu.CurrMap.CurrPlayer.X
			item.Y = bagMenu.CurrMap.CurrPlayer.Y
			bagMenu.CurrMap.AddItem(item)
		} else if rl.IsKeyPressed(rl.KeyU) {
			if bagMenu.CurrMap.CurrPlayer.UpgradePoint >= item.Level {
				canLevelUp := bagMenu.CurrMap.CurrPlayer.Bag[bagMenu.SelectedItem].LevelUp(bagMenu.CurrMap)
				if canLevelUp {
					bagMenu.CurrMap.CurrPlayer.UpgradePoint -= item.Level
				}
			}
		}
	}
	if bagMenu.IsListFocused {
		if bagMenu.CurrMap.CurrPlayer.BagSize >= 1 {
			if rl.IsKeyPressed(rl.KeyLeft) {
				bagMenu.SelectedItem--
				if bagMenu.SelectedItem < 0 {
					bagMenu.SelectedItem = 0
					// bagMenu.SelectedItem = bagMenu.CurrMap.CurrPlayer.BagSize - 1
				}
			} else if rl.IsKeyPressed(rl.KeyRight) {
				bagMenu.SelectedItem++
				if bagMenu.SelectedItem > bagMenu.CurrMap.CurrPlayer.BagSize - 1 {
					bagMenu.SelectedItem = bagMenu.CurrMap.CurrPlayer.BagSize - 1
					// bagMenu.SelectedItem = 0
				}
			}
		}
	} else {
		if rl.IsKeyPressed(rl.KeyLeft) {
			bagMenu.SelectedTab--
			if bagMenu.SelectedTab < 0 {
				bagMenu.SelectedTab = int32(len(tabs)) - 1
			}
		} else if rl.IsKeyPressed(rl.KeyRight) {
			bagMenu.SelectedTab++
			if bagMenu.SelectedTab > int32(len(tabs)) - 1 {
				bagMenu.SelectedTab = 0
			}
		}
	}
}

func (bagMenu *BagMenu) drawTabContent(currItem Item, currX, currY int32) {
	currX += 20
	rl.DrawText(string(currItem.Name), currX, currY + 50, 23, rl.NewColor(144, 12, 63, 255))
	rl.DrawRectangle(currX, currY + 85, 140, 2, rl.DarkGray)
	currY += 120
	switch tabs[bagMenu.SelectedTab] {
	case "Description":
		var lineCount int32 = 0
		var maxLineChar int32 = 30
		var currLine string = ""
		var currWordLength int32
		var currLineLength int32
		for _, word := range strings.Split(currItem.Description, " ") {
			currWordLength = int32(len(word))
			currLineLength = int32(len(currLine))
			if currLineLength + currWordLength + 1 > maxLineChar {
				rl.DrawText(currLine, currX, currY + lineCount * 30, 20, rl.DarkGray)
				lineCount++
				currLine = word
			} else {
				if currLine != "" {
					currLine += " "
				}
				currLine += word
			}
		}
		rl.DrawText(currLine, currX, currY + lineCount * 30, 20, rl.DarkGray)
		break

	case "Upgrades":
		var i int32
		var color rl.Color = rl.DarkGray
		for i = 0; i<IML; i++ {
			if i+1 <= currItem.Level {
				color = rl.NewColor(200, 32, 16, 255)
			} else {
				color = rl.DarkGray
			}
			rl.DrawText(currItem.GetLevelUpDescription(i), currX, currY + i * 30, 18, color)
		}
		break
	}
}

func (bagMenu *BagMenu) Draw() {
	var startX, startY int32 = bagMenu.X + bagMenu.BorderSize, bagMenu.Y + bagMenu.BorderSize
	var contentWidth, contentHeight = bagMenu.Width - bagMenu.BorderSize*2, bagMenu.Heigth - bagMenu.BorderSize*2
	rl.DrawRectangle(bagMenu.X, bagMenu.Y, bagMenu.Width, bagMenu.Heigth, rl.NewColor(65, 87, 106, 255))
	rl.DrawRectangle(startX, startY, contentWidth, contentHeight, rl.RayWhite)
	var currDY int32 = 10
	rl.DrawText("Bag", startX + 190, startY + currDY, 25, rl.DarkGray)
	var infoBorderSize int32 = 5
	var itemListMargin int32 = 30
	var itemListHeigth int32 = 60
	currDY += 55
	var listBorderColor rl.Color = rl.LightGray
	if bagMenu.IsListFocused {
		listBorderColor = rl.Gray
	}
	rl.DrawRectangle(startX + itemListMargin, startY + currDY, contentWidth - itemListMargin*2, itemListHeigth, listBorderColor)
	rl.DrawRectangle(startX + itemListMargin + infoBorderSize, startY + currDY + infoBorderSize, contentWidth - itemListMargin*2 - infoBorderSize*2, itemListHeigth - infoBorderSize*2, rl.LightGray)

	var bagSize = bagMenu.CurrMap.CurrPlayer.BagSize
	var itemMarginLeft, itemMarginTop int32 = 15, 5
	var currItemX, currItemY int32 = startX + itemListMargin + itemMarginLeft + 8, startY + currDY + itemMarginTop
	v1, v2, v3 := rl.Vector2{float32(startX + itemListMargin - 8), float32(startY + currDY + itemListHeigth/2)}, rl.Vector2{float32(startX + itemListMargin + itemMarginLeft - 7), float32(startY + currDY + itemListHeigth/2 - 10)}, rl.Vector2{float32(startX + itemListMargin + itemMarginLeft - 7), float32(startY + currDY + itemListHeigth/2 + 10)}
	if bagMenu.SelectedItem >= 7 {
		rl.DrawTriangle(v1, v3, v2, rl.NewColor(120, 100, 95, 255))
	}
	if bagMenu.SelectedItem < bagSize - 1 && bagSize > 7 {
		v1.X += float32(contentWidth - itemListMargin*2 + 16)
		v2.X += float32(contentWidth - itemListMargin*2 - 17)
		v3.X += float32(contentWidth - itemListMargin*2 - 17)
		rl.DrawTriangle(v1, v2, v3, rl.NewColor(120, 100, 95, 255))
	}
	if bagSize > 0 {
		var itemSize = bagMenu.CurrMap.CurrPlayer.Bag[0].Size + 5
		var currItemY = startY + currDY + itemListHeigth - (itemListHeigth-itemSize)/2 - itemSize
		var i int32 = 0
		if bagMenu.SelectedItem >= 7 {
			i = bagMenu.SelectedItem
		}
		var count int32 = 0
		for i = i; i < bagSize; i++ {
			if count >= 7 {
				break
			}
			item := bagMenu.CurrMap.CurrPlayer.Bag[i]
			item.X = currItemX
			item.Y = currItemY
			item.Size = itemSize
			if i == bagMenu.SelectedItem {
				item.Selected = true
			}
			item.Draw()
			currItemX += itemSize + itemMarginLeft
			count++
		}
	} else {
		rl.DrawText("Empty", currItemX, currItemY + 15, 20, rl.DarkGray)
	}

	currDY += itemListHeigth + 30
	var infoBorderColor rl.Color = rl.LightGray
	if !bagMenu.IsListFocused {
		infoBorderColor = rl.Gray
	}
	rl.DrawRectangle(startX + itemListMargin, startY + currDY, contentWidth - itemListMargin*2, contentHeight - currDY - 60, infoBorderColor)
	rl.DrawRectangle(startX + itemListMargin + infoBorderSize, startY + currDY + infoBorderSize, contentWidth - itemListMargin*2 - infoBorderSize*2, contentHeight - currDY - 60 - infoBorderSize*2, rl.LightGray)
	
	var currX int32 = startX + itemListMargin + infoBorderSize
	var currY int32 = startY + currDY + infoBorderSize

	var tabWidth int32 = 120
	var tabHeight int32 = 35
	var tabMargin int32 = 8
	for index, tab := range tabs {
		i := int32(index)
		if i == bagMenu.SelectedTab {
			rl.DrawRectangle(currX + i * tabWidth, currY, tabWidth, tabHeight, rl.DarkGray)
			rl.DrawText(tab, currX + i * tabWidth + tabMargin, currY + tabMargin, 19, rl.LightGray)
		} else {
			rl.DrawRectangle(currX + i * tabWidth, currY, tabWidth, tabHeight, rl.LightGray)
			rl.DrawText(tab, currX + i * tabWidth + tabMargin, currY + tabMargin, 19, rl.DarkGray)
		}
	}

	currY += 10
	currItem := bagMenu.CurrMap.CurrPlayer.Bag[bagMenu.SelectedItem]
	if bagSize > 0 {
		bagMenu.drawTabContent(currItem, currX, currY)
	}

	currY = startY + contentHeight - 30 - infoBorderSize*2
	startX += 20
	rl.DrawText("Close : ", startX + 20, currY, 19, rl.DarkGray)
	util.ShowBackspaceKey(startX + 90, currY)
	if bagSize > 0 {
		rl.DrawText("Drop : ", startX + 140, currY, 19, rl.DarkGray)
		util.ShowClassicKey(startX + 205, currY, "A")
		if bagMenu.CurrMap.CurrPlayer.UpgradePoint >= currItem.Level && currItem.CanLevelUp() {
			rl.DrawText("Upgrade : ", startX + 250, currY, 19, rl.DarkGray)
			util.ShowClassicKey(startX + 350, currY, "U")		
		}
	}

}