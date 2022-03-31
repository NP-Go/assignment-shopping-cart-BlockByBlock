package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func Select(input *int, shoplist *types.StoreInfo) {
	switch *input {
	case 1:
		viewShoppingList(shoplist)
	case 2:
		fmt.Println("It's 2!")
	case 3:
		addItem(shoplist)
	case 4:
		modifyItem(shoplist)
	case 5:
		deleteItem(shoplist)
	case 6:
		printCurrentData(shoplist)
	case 7:
		addNewCategory(shoplist)
	default:
		fmt.Println("Option unavailable!")
	}
}
