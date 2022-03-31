package types

import "fmt"

type StoreInfo struct {
	Categories  []string
	AllItemInfo map[string]ItemInfo
}

func (storeInfo *StoreInfo) InitStoreInfo() {
	storeInfo.Categories = []string{"Household", "Food", "Drinks"}
	storeInfo.AllItemInfo = make(map[string]ItemInfo)

	// default items - lazy
	storeInfo.AllItemInfo["egg"] = ItemInfo{Category: 1, Qty: 10, UnitCost: 2.55}
	storeInfo.AllItemInfo["tea"] = ItemInfo{Category: 2, Qty: 37, UnitCost: 1.20}
	storeInfo.AllItemInfo["coffee"] = ItemInfo{Category: 2, Qty: 49, UnitCost: 1.40}
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
