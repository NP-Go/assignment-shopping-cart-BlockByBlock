package main

import (
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/selections"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/utils"
)

type itemInfo struct {
	category int
	qty      int
	unitCost float64
}

type storeInfo struct {
	categories []string
	itemMap    map[string]itemInfo
}

var thisShop storeInfo

func init() {
	thisShop.categories = []string{"Household", "Food", "Drinks"}
}

func main() {
	utils.WelcomeMenu()
	for {
		selection := utils.GetMenuInput()
		selections.Select(&selection)
	}
}
