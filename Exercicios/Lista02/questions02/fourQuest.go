package questions02

import (
	"errors"
	"fmt"
)

// In a certain State, for vehicle transfers, DETRAN charges a 1% tax for cars
// manufactured before 1990 and a 1.5% tax for those manufactured from 1990 onwards,
// this tax is levied on the car's price. Develop an algorithm that reads the car's year
// and price, calculates and writes the tax to be paid.

func Four() {

	price, year, err := valid("Enter the price of car: ", "Enter the year of car: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := applyingTax(price, year)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(value)
}

func valid(prompt string, prompt2 string) (int, int, error) {
	var price, year int

	fmt.Print(prompt)
	_, err := fmt.Scan(&price)
	if err != nil {
		return 0, 0, errors.New("enter the integer value")
	}

	fmt.Print(prompt2)
	_, err = fmt.Scan(&year)
	if err != nil {
		return 0, 0, errors.New("enter the integer value please")
	}

	return price, year, nil
}

func applyingTax(price, year int) (float64, error) {
	var value float64
	if year >= 1990 {
		value = float64(price) * 0.015
	} else {
		value = float64(price) * 0.01
	}
	value = float64(price) + value
	return value, nil
}
