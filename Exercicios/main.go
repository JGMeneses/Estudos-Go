package main

import (
	"fmt"
	"os"
	"os/exec"
	questions "projeto001/Exercicios/Lista01/questions"
	questions02 "projeto001/Exercicios/Lista02/questions02"
)

func main() {

	var listaChoice, questionChoice int

	fmt.Println("Which list do you want to access?")
	fmt.Println("1- List 1")
	fmt.Println("2- List 2")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&listaChoice)
	clearScreen()
	switch listaChoice {
	case 1:
		fmt.Println("Which question do you want to access from List 1?")
		fmt.Print("Which question do you want to access? \n 1- First question \n 2- Second question")
		fmt.Print("\n 3- Third question \n 4- Fourth question\n 5- Five question \n 6- Five question")
		fmt.Print("\n 7- Third question \n 8- Fourth question\n 9- Five question \n 10- Five question\n")
		fmt.Scan(&questionChoice) // Read user's choice
		clearScreen()             // Clear the screen
		switch questionChoice {
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
	case 2:
		fmt.Println("Which question do you want to access from List 2?")
		fmt.Print("Which question do you want to access? \n 1- First question \n 2- Second question")
		fmt.Print("\n 3- Third question \n 4- Fourth question\n 5- Five question \n 6- Five question")
		fmt.Print("\n 7- Third question \n 8- Fourth question\n 9- Five question \n 10- Five question\n")
		fmt.Scan(&questionChoice) // Read user's choice
		clearScreen()             // Clear the screen
		switch questionChoice {
		case 1:
			questions02.One()
		case 2:
			questions02.Two()
		case 3:
			questions02.Three()
		case 4:
			questions02.Four()
		case 5:
			questions02.Five()
		case 6:
			questions02.Six()
		case 7:
			questions02.Seven()
		default:
			fmt.Println("Invalid option!")
		}
	}
}

// clearScreen clears the terminal screen.
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
