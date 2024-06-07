package main

import (
	"fmt"
	questions "projeto001/Exercicios/Lista01/questions"
)

func main() {
	// Call the functions from firstQuest.go and secondQuest.go as needed
	var choice int

	fmt.Print("Which question do you want to access? \n 1- First question \n 2- Second question")
	fmt.Print("\n 3- Third question \n 4- Fourth question\n 5- Five question")
	fmt.Scan(&choice) // Read user's choice

	switch choice {
	case 1:
		questions.One()
	case 2:
		questions.Two()
	case 3:
		questions.Three()
	case 4:
		questions.Four()
	case 5:
		questions.Five()
	default:
		fmt.Println("Invalid option!")
	}
}
