package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func addItem(shoplist *types.StoreInfo) {
	var (
		name     string
		category string
		qty      int
		cost     float64
	)

	fmt.Println("What is the name of your item")
	fmt.Scanln(&name)

	fmt.Println("What category does it belongs to")
	fmt.Scanln(&category)

	fmt.Println("How many units are there?")
	fmt.Scanln(&qty)

	fmt.Println("How much does it costs per unit?")
	fmt.Scanln(&cost)

	shoplist.AddItemInfo(&name, &types.ItemInfo{Category: 0, Qty: qty, UnitCost: cost})
}
