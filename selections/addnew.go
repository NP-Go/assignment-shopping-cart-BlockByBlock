package selections

import (
	"fmt"

	"github.com/NP-Go/assignment-shopping-cart-BlockByBlock/types"
)

func addNewCategory(shoplist *types.StoreInfo) {
	fmt.Println(`
	Add New Category Name
	What is the New Category Name to add?
	`)

	newCategory := ""
	if _, err := fmt.Scan(&newCategory); err != nil && newCategory != "" {
		fmt.Println("No Input Found")
	}

	shoplist.AddCategory(newCategory)
}
