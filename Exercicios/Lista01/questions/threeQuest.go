package questions

import "fmt"

// Read the value of a purchase at "Your Best Purchase" store and show the value of the installments
// according to the number of installments desired by the user. The store sells its products
// either for cash or in up to 10 installments without interest.
func Three() {
	var value int
	installments := 11
	fmt.Print("Enter the value of your purchase:")
	fmt.Scan(&value)

	for installments > 10 {
		fmt.Print("Enter the desired number of installments (up to 10x):")
		fmt.Scan(&installments)
	}
	if installments <= 0 {
		fmt.Printf("Your purchase will be for cash and the total amount to pay is R$ %d\n", value)
	} else {
		fmt.Printf("Your purchase will be in %d installments and the final amount to pay will be R$ %.2f per installment\n", installments, float64(value)/float64(installments))
	}
}
