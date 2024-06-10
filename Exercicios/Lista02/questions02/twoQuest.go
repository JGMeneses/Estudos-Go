package questions02

import (
	"errors"
	"fmt"
)

func Two() {

	num, err := inputValue("Enter the number: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	if evenOrOdd(num) {
		if num > 100 {
			fmt.Println("The number is greater than 100")
		} else {
			fmt.Println("The number is not greater than 100")
		}
	} else {
		if num < 0 {
			fmt.Println("The number is negative")
		} else {
			fmt.Println("The number is not negative")
		}
	}
}

func inputValue(prompt string) (int, error) {
	var value int
	fmt.Print(prompt)
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, errors.New("invalid input, please enter an integer value")
	}

	return value, nil
}

func evenOrOdd(num int) bool {
	return num%2 == 0
}
