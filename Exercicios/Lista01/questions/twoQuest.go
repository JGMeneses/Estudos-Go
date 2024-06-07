package questions

import "fmt"

// Calculate and write the amount of money spent by a smoker based on reading the
// information: number of years they have been smoking, quantity of cigarettes smoked per day, and price of a pack.
func Two() {
	var yearsSmoking, dailyQuantity, price float64

	fmt.Println("How many years have you been smoking:")
	fmt.Scan(&yearsSmoking)
	fmt.Println("How many cigarettes do you smoke per day:")
	fmt.Scan(&dailyQuantity)
	fmt.Println("Price of a pack of cigarettes:")
	fmt.Scan(&price)

	expense := dailyQuantity / 20 * (yearsSmoking * 365) * price
	fmt.Print("Your expense on cigarettes until today: ", expense)
}
