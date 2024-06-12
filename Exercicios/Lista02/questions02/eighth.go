package questions02

import (
	"errors"
	"fmt"
)

// Elabore um algoritmo que leia um número inteiro e uma letra (“a” – antecessores ou “s” – sucessores) e
// mostre a soma dos seus próximos 10 antecessores ou sucessores de acordo com a letra digitada.

// Create an algorithm that reads an integer and a letter (“a” - predecessors or “s” - successors) and
// display the sum of its next 10 predecessors or successors according to the letter entered.

func Eighth() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	err := letterAndNumValidator("Enter the letter: ", number)
	if err != nil {
		fmt.Print("Error in the function")
	}
}

func letterAndNumValidator(prompt string, num int) error {
	var letter string
	fmt.Println(prompt)
	_, err := fmt.Scan(&letter)
	if err != nil {
		return errors.New("enter a string")
	}

	switch letter {
	case "a":
		if num <= 0 {
			return errors.New("the sum of its predecessors is negative")
		} else {
			result := sumPredecessors(num)
			fmt.Println("\n Final sum", result)
		}

	case "s":
		if num <= 0 {
			return errors.New("the sum of its predecessors is negative")
		} else {
			result := sumSucessors(num)
			fmt.Println("\n Final sum", result)
		}
	default:
	}
	return err
}

func sumPredecessors(num int) int {
	var sum int
	for i := num - 1; i >= num-10; i-- {
		fmt.Print(i, ",")

		sum += i
	}

	return sum

}

func sumSucessors(num int) int {
	var sum int
	for i := num - 1; i <= num+10; i++ {
		fmt.Print(i, ",")

		sum += i
	}

	return sum

}
