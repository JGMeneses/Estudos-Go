package questions

import "fmt"

// Read the age of a person expressed in years, months, and days and display it expressed only in days.
// Note: Consider each month to have 30 days.

func One() {
	var days, months, year int

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&year)
	fmt.Print("Enter the number of months: ")
	fmt.Scan(&months)
	fmt.Print("Enter the number of days: ")
	fmt.Scan(&days)

	finalAge := year*360 + (months*30 + days)

	fmt.Print(finalAge)
}
