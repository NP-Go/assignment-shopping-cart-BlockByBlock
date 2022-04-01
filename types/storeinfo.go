package types

import "fmt"

type StoreInfo struct {
	Categories  []string
	AllItemInfo map[string]ItemInfo
}

func (storeInfo *StoreInfo) InitStoreInfo() {
	storeInfo.Categories = []string{"Household", "Food", "Drinks"}
	storeInfo.AllItemInfo = make(map[string]ItemInfo)

	// test data
	storeInfo.AllItemInfo["Fork"] = ItemInfo{Category: 0, Qty: 4, UnitCost: 3}
	storeInfo.AllItemInfo["Plates"] = ItemInfo{Category: 0, Qty: 4, UnitCost: 3}
	storeInfo.AllItemInfo["Cups"] = ItemInfo{Category: 0, Qty: 5, UnitCost: 3}

	storeInfo.AllItemInfo["Bread"] = ItemInfo{Category: 1, Qty: 5, UnitCost: 2}
	storeInfo.AllItemInfo["Cake"] = ItemInfo{Category: 1, Qty: 3, UnitCost: 1}

	storeInfo.AllItemInfo["Coke"] = ItemInfo{Category: 2, Qty: 5, UnitCost: 2}
	storeInfo.AllItemInfo["Sprite"] = ItemInfo{Category: 2, Qty: 5, UnitCost: 2}
}

func (storeInfo *StoreInfo) AddCategory(newCategory *string) {
	currentSize := len(storeInfo.Categories)

	// check if category exist
	for i, category := range storeInfo.Categories {
		if category == *newCategory {
			fmt.Printf("Category: Food already exist at index %d", i)
			return
		}
	}

	storeInfo.Categories = append(storeInfo.Categories, *newCategory)

	fmt.Printf("New category: %s added at index %d", *newCategory, currentSize)
}

func (storeInfo *StoreInfo) GetCategoryIndex(category *string) int {
	for i, currCategory := range storeInfo.Categories {
		if currCategory == *category {
			return i
		}
	}
	return -1
}

func (storeInfo *StoreInfo) PrintAll() {
	fmt.Println("Shopping List Contents:")

	categoryName := "Others"

	for itemName, itemInfo := range storeInfo.AllItemInfo {
		if itemInfo.Category != -1 {
			categoryName = storeInfo.Categories[itemInfo.Category]
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

func (storeInfo *StoreInfo) PrintItemInfo() {
	fmt.Println("Print Current Data")

	for itemName, itemInfo := range storeInfo.AllItemInfo {
		fmt.Printf("%s - %v\n", itemName, itemInfo)
	}
}
