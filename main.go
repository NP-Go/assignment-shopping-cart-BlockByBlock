package main

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/selections"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/utils"
)

func main() {
	thisShop := new(types.StoreInfo)
	thisShop.SetCategories([]string{"Household", "Food", "Drinks"})

	// temp
	fmt.Println(thisShop)

	utils.WelcomeMenu()
	for {
		selection := utils.GetMenuInput()
		selections.Select(&selection)
	}
}
