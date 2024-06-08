package questions

import "fmt"

//Escreva o valor em reais (R$) de um valor lido em dólares (US$). O algoritmo deverá solicitar o
//valor da cotação do dólar e também a quantidade de dólares que o usuário deseja converter.

// Write the value in reais (R$) of a value read in dollars (US$). The algorithm should request the
// dollar exchange rate and also the quantity of dollars that the user wants to convert.

func Eight() {
	var rate, quantityDollars float64

	fmt.Print("Enter  dollar exchange rate: ")
	fmt.Scan(&rate)

	fmt.Print("Enter the quantity of dollars you want: ")
	fmt.Scan(&quantityDollars)

	valueReais := quantityDollars * rate
	fmt.Print("The value of dollars converted to Real was: ", valueReais)

}
