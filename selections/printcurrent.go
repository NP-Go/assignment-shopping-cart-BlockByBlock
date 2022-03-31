package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func printCurrentData(shoplist *types.StoreInfo) {
	fmt.Println("Print Current Data")

	for itemName, itemInfo := range shoplist.AllItemInfo {
		fmt.Printf("%s - %v\n", itemName, itemInfo)
	}
}
