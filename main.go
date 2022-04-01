package main

import (
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/selections"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/utils"
)

func main() {
	thisShop := new(types.StoreInfo)
	thisShop.InitStoreInfo()

	for {
		utils.WelcomeMenu()
		selection := utils.GetMenuInput()

		if selection > 0 {
			selections.SelectMenu(&selection, thisShop)
		}
	}
}
