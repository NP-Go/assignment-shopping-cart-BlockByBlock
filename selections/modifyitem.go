package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func modifyItem(shoplist *types.StoreInfo) {
	var (
		itemName    string
		newItemName string
		newCategory string
		newQty      int
		newUnitCost float64
	)

	fmt.Println(`
	Modify Items.
	Which item would you wish to modify?
	`)
	fmt.Scanln(&itemName)

	itemInfo, exist := shoplist.AllItemInfo[itemName]

	if !exist {
		fmt.Println("Item dont exist")
		return
	}

	// for invalid category index
	categoryName := "Others"
	if itemInfo.Category != -1 {
		categoryName = shoplist.Categories[itemInfo.Category]
	}

	fmt.Printf(
		"Current item name is %s - Category is %s - Quantity is %d - Unit Cost %f\n",
		itemName,
		categoryName,
		itemInfo.Qty,
		itemInfo.UnitCost,
	)

	fmt.Println("Enter new name. Enter for no change")
	fmt.Scanln(&newItemName)
	fmt.Println("Enter new category. Enter for no change")
	fmt.Scanln(&newCategory)
	fmt.Println("Enter new quantity. Enter for no change")
	fmt.Scanln(&newQty)
	fmt.Println("Enter new unit cost. Enter for no change")
	fmt.Scanln(&newUnitCost)

	if newCategory == "" {
		fmt.Println("No changes to category made")
	} else {
		itemInfo.Category = shoplist.GetCategoryIndex(&newCategory)
		shoplist.AllItemInfo[itemName] = itemInfo
	}

	// does not handle set qty 0
	if newQty == 0 {
		fmt.Println("No changes to quantity made")
	} else {
		itemInfo.Qty = newQty
		shoplist.AllItemInfo[itemName] = itemInfo
	}

	if newUnitCost == 0 {
		fmt.Println("No changes to unit cost made")
	} else {
		itemInfo.UnitCost = newUnitCost
		shoplist.AllItemInfo[itemName] = itemInfo
	}

	if newItemName == "" {
		fmt.Println("No changes to item name made")
	} else {
		itemInfo = shoplist.AllItemInfo[itemName] // handle prev changes
		delete(shoplist.AllItemInfo, itemName)
		shoplist.AllItemInfo[newItemName] = itemInfo
	}
}
