package types

import "fmt"

type StoreInfo struct {
	Categories []string
}

func (storeInfo *StoreInfo) SetCategories(categories []string) {
	storeInfo.Categories = categories
}

func (storeInfo *StoreInfo) AddCategory(newCategory string) {
	currentSize := len(storeInfo.Categories)

	// check if category exist
	for i, category := range storeInfo.Categories {
		if category == newCategory {
			fmt.Printf("Category: Food already exist at index %d", i)
			return
		}
	}

	storeInfo.Categories = append(storeInfo.Categories, newCategory)

	fmt.Printf("New category: %s added at index %d", newCategory, currentSize)
}
