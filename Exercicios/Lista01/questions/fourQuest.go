package questions

import "fmt"

// Read an integer number N and a percentage, and write the value of the percentage
// applied to the number.
func Four() {
	var number, percentage float64

	fmt.Print("Enter an integer number N: ")
	fmt.Scan(&number)
	fmt.Print("Enter a percentage: ")
	fmt.Scan(&percentage)

	result := number * (percentage / 100.0)

	fmt.Println("Final value:", result)
}
