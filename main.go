package main

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/selections"
	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/utils"
)

func init() {
	// init struct here
	fmt.Println("")
}

func main() {
	utils.WelcomeMenu()
	selection := utils.GetMenuInput()
	selections.Select(&selection)
}
