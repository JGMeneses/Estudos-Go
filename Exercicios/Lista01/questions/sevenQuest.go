package questions

import "fmt"

// Leia uma temperatura em graus Celsius e apresente-a convertida em graus Fahrenheit. A
// fórmula de conversão é F=(9*C+160)/5, sendo F a temperatura em Fahrenheit e C a temperatura
// em Celsius.

// Read a temperature in degrees Celsius and present it converted in degrees Fahrenheit. The conversion formula is F=(9*C+160)/5,
// with F being the temperature in Fahrenheit and C being the temperature in Celsius.

func Seven() {
	var degreesCelsius float64

	fmt.Println("Enter temperature in degrees Celsius:")
	fmt.Scan(&degreesCelsius)

	F := (9*degreesCelsius + 160) / 5

	fmt.Println("temperature in Fahrenheit: ", F)
}
