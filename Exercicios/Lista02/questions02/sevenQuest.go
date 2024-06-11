package questions02

import (
	"errors"
	"fmt"
)

// Escreva um algoritmo que leia o código dos itens pedidos (um sanduiche e uma bebida) e a quantidade
// de cada um e calcule o valor a ser pago por aquele lanche.

// Write an algorithm to receive an order (a sandwich and a drink), receive the quantity of each item, and then calculate the total of the order.

type MenuItem struct {
	Description string
	Price       float64
}

var menu = map[int]MenuItem{
	100: {"Hot Dog", 1.10},
	101: {"Simple Bauru", 1.30},
	102: {"Bauru with Egg", 1.50},
	103: {"Hamburger", 1.10},
	104: {"Cheeseburger", 1.30},
	105: {"Soda", 1.00},
	106: {"Juice", 2.00},
	107: {"Nescau", 1.50},
}

func Seven() {
	printMenu()

	sandwichCode, sandwichQty, err := getOrder("Enter the sandwich code: ", "Enter the quantity: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	drinkCode, drinkQty, err := getOrder("Enter the drink code: ", "Enter the quantity: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	total, err := calculateTotal(sandwichCode, sandwichQty, drinkCode, drinkQty)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The total amount to be paid is: $%.2f\n", total)
}

func printMenu() {
	fmt.Println("Sandwiches:")
	fmt.Println("Code | Description     | Unit Price")
	fmt.Println("100  | Hot Dog         | $1.10")
	fmt.Println("101  | Simple Bauru    | $1.30")
	fmt.Println("102  | Bauru with Egg  | $1.50")
	fmt.Println("103  | Hamburger       | $1.10")
	fmt.Println("104  | Cheeseburger    | $1.30")
	fmt.Println("")
	fmt.Println("Drinks:")
	fmt.Println("Code | Description     | Unit Price")
	fmt.Println("105  | Soda            | $1.00")
	fmt.Println("106  | Juice           | $2.00")
	fmt.Println("107  | Nescau          | $1.50")
	fmt.Println("")
}

func getOrder(codePrompt, qtyPrompt string) (int, int, error) {
	var code, qty int
	fmt.Print(codePrompt)
	_, err := fmt.Scan(&code)
	if err != nil {
		return 0, 0, errors.New("error reading the item code")
	}

	fmt.Print(qtyPrompt)
	_, err = fmt.Scan(&qty)
	if err != nil {
		return 0, 0, errors.New("error reading the quantity")
	}
	if _, exists := menu[code]; !exists {
		return 0, 0, errors.New("invalid item code")
	}

	return code, qty, nil
}

func calculateTotal(sandwichCode, sandwichQty, drinkCode, drinkQty int) (float64, error) {
	sandwichItem, sandwichExists := menu[sandwichCode]
	drinkItem, drinkExists := menu[drinkCode]

	if !sandwichExists || !drinkExists {
		return 0, errors.New("One or more item codes are invalid")
	}

	total := (sandwichItem.Price*float64(sandwichQty) + (drinkItem.Price * float64(drinkQty)))

	return total, nil
}
