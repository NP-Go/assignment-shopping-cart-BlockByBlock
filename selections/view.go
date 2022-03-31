package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func viewShoppingList(shoplist *types.StoreInfo) {
	fmt.Println("Shopping List Contents:")

	categoryName := "Others"

	for itemName, itemInfo := range shoplist.AllItemInfo {
		if itemInfo.Category != -1 {
			categoryName = shoplist.Categories[itemInfo.Category]
		}

		fmt.Printf(
			"Category: %s - Item: %s Quantity: %d Unit Cost: %f\n",
			categoryName,
			itemName,
			itemInfo.Qty,
			itemInfo.UnitCost,
		)
	}
}
