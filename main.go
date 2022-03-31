package main

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/selections"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/utils"
)

func main() {
	thisShop := new(types.StoreInfo)
	thisShop.InitStoreInfo()

	// temp
	fmt.Println(thisShop)

	for {
		utils.WelcomeMenu()
		selection := utils.GetMenuInput()
		selections.Select(&selection, thisShop)
	}
}
