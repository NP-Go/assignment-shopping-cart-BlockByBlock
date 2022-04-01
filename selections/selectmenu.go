package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func SelectMenu(input *int, shoplist *types.StoreInfo) {
	switch *input {
	case 1:
		shoplist.PrintAll()
	case 2:
		generateReport(shoplist)
	case 3:
		addItem(shoplist)
	case 4:
		modifyItem(shoplist)
	case 5:
		deleteItem(shoplist)
	case 6:
		shoplist.PrintItemInfo()
	case 7:
		addNewCategory(shoplist)
	default:
		fmt.Println("Option unavailable!")
	}
}
