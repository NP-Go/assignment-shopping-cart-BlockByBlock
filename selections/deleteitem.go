package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func deleteItem(shoplist *types.StoreInfo) {
	var itemName string

	fmt.Println(`
	Delete Item.
	Enter item name to delete:
	`)
	fmt.Scanln(&itemName)

	_, exist := shoplist.AllItemInfo[itemName]

	if exist {
		delete(shoplist.AllItemInfo, itemName)
		fmt.Printf("Deleted %s", itemName)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
}
