package main

import (
	"fmt"
	questions "projeto001/Exercicios/Lista01/questions"
)

func main() {
	// Call the functions from firstQuest.go and secondQuest.go as needed
	var choice int

	fmt.Print("Which question do you want to access? \n 1- First question \n 2- Second question")
	fmt.Print("\n 3- Third question \n 4- Fourth question\n 5- Five question \n 6- Five question")
	fmt.Print("\n 7- Third question \n 8- Fourth question\n 9- Five question \n 10- Five question\n")
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
	case 6:
		questions.Six()
	case 7:
		questions.Seven()
	default:
		fmt.Println("Invalid option!")
	}
}
