package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func generateReport(shoplist *types.StoreInfo) {
	var input int

	fmt.Println(`
	Generate Report
	1. Total Cost of each category
	2. List of item by category
	3. Main Menu
	`)
	fmt.Scanln(&input)

	if input == 1 {
		fmt.Println("Total cost by Category")
	}

	if input == 2 {
		fmt.Println("List by Category")

		for itemName, itemInfo := range shoplist.AllItemInfo {
			for i := 0; i < len(shoplist.Categories); i++ {
				if itemInfo.Category == i {
					categoryName := shoplist.Categories[itemInfo.Category]

					fmt.Printf(
						"Category: %s - Item: %s Quantity: %d Unit Cost: %f\n",
						categoryName,
						itemName,
						itemInfo.Qty,
						itemInfo.UnitCost,
					)
				}

			}
		}
	}

	return
}
