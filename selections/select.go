package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func Select(input *int, shoplist *types.StoreInfo) {
	switch *input {
	case 1:
		viewShoppingList()
	case 2:
		fmt.Println("It's 2!")
	case 3:
		addItem(shoplist)
	case 4:
		fmt.Println("It's 4!")
	case 5:
		fmt.Println("It's 5!")
	case 6:
		fmt.Println("It's 6!")
	case 7:
		addNewCategory(shoplist)
	default:
		fmt.Println("Option unavailable!")
	}
}
