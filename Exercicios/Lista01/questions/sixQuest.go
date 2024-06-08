package questions

import "fmt"

// O custo ao consumidor de um carro novo é a soma do custo de fábrica com a percentagem do
// distribuidor e dos impostos, ambos aplicados ao custo de fábrica. Supondo que a percentagem
// do distribuidor seja de 28% e os impostos de 45%, escreva um algoritmo que leia o custo de
// fábrica de um carro e escreva o custo ao consumidor.

// The cost to the consumer of a new car is the sum of the factory cost with the percentage of the distributor and the taxes,
// both applied to the factory cost. Assuming the percentage of the distributor is 28% and the taxes are 45%, write an algorithm that reads the factory
// cost of a car and writes the consumer cost.

func Six() {
	var factoryCost float64

	fmt.Print("Enter the factory cost: ")
	fmt.Scan(&factoryCost)

	distributorCost := factoryCost * 0.28
	taxCost := factoryCost * 0.45

	consumerCost := factoryCost + distributorCost + taxCost

	fmt.Println("The consumer cost is:", consumerCost)
}
