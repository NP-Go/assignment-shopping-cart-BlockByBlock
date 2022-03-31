package types

import "fmt"

type StoreInfo struct {
	Categories  []string
	AllItemInfo map[string]ItemInfo
}

func (storeInfo *StoreInfo) InitStoreInfo() {
	storeInfo.Categories = []string{"Household", "Food", "Drinks"}
	storeInfo.AllItemInfo = make(map[string]ItemInfo)
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

func (storeInfo *StoreInfo) AddItemInfo(itemName *string, itemInfo *ItemInfo) {
	storeInfo.AllItemInfo[*itemName] = *itemInfo
	fmt.Printf("Size of map %d", len(storeInfo.AllItemInfo))
}
