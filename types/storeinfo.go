package types

import "fmt"

type StoreInfo struct {
	Categories  []string
	AllItemInfo map[string]ItemInfo
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
