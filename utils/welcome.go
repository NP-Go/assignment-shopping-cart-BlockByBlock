package utils

import (
	"fmt"
)

// printed on runtime prompt
func WelcomeMenu() {
	msg := `
	Shopping List Application
	==========================
	1. View entire shopping list.
	2. Generate Shopping List Report.
	3. Add Items.
	4. Modify Items.
	5. Delete Item.
	6. Print Current Data.
	7. Add New Category Name.
	`

	fmt.Printf("%s", msg)
	fmt.Println("Select your choice: ")

	getMenuInput()
}

func getMenuInput() int {
	input := 0
	if _, err := fmt.Scan(&input); err != nil {
		fmt.Println("Input failed")
	}

	return input
}
