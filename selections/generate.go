package selections

import (
	"fmt"
	"sync"

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
		var wg sync.WaitGroup

		for i := 0; i < len(shoplist.Categories); i++ {
			wg.Add(1)
			go func(categoryIndex int) {
				defer wg.Done()
				sumOfCategory(shoplist, categoryIndex)
			}(i)
		}

		wg.Wait()
	}

	if input == 2 {
		fmt.Println("List by Category")

		for i := 0; i < len(shoplist.Categories); i++ {
			categoryName := shoplist.Categories[i]
			for itemName, itemInfo := range shoplist.AllItemInfo {
				if itemInfo.Category == i {
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

func sumOfCategory(shoplist *types.StoreInfo, i int) {
	categoryName := shoplist.Categories[i]
	totalCost := 0.00

	for _, itemInfo := range shoplist.AllItemInfo {
		if itemInfo.Category == i {
			totalCost += float64(itemInfo.Qty) * itemInfo.UnitCost
		}
	}

	fmt.Printf("%s cost: %f\n", categoryName, totalCost)
}
