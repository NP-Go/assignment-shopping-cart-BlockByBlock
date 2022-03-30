package selections

import "fmt"

func Select(input *int) {
	switch *input {
	case 1:
		ViewShoppingList()
	case 2:
		fmt.Println("It's 2!")
	case 3:
		fmt.Println("It's 3!")
	case 4:
		fmt.Println("It's 4!")
	case 5:
		fmt.Println("It's 5!")
	case 6:
		fmt.Println("It's 6!")
	case 7:
		fmt.Println("It's 7!")
	default:
		fmt.Println("Option unavailable!")
	}
}
