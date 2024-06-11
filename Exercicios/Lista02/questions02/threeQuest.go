package questions02

import (
	"errors"
	"fmt"
)

// Calcule o peso ideal de uma pessoa. Dados de entrada: altura e gênero (“m”-masculino ou “f”-feminino).
// Utilize as seguintes fórmulas para cálculo do peso ideal:
// • Masculino = (72,7 x altura) - 58
// • Feminino = (62,1 x altura) - 44,7

func Three() {
	weight, gender, err := validInput("Enter the number: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	var idealWeight float64
	if gender == "m" {
		idealWeight = (72.7 * float64(weight)) - 58
	} else if gender == "f" {
		idealWeight = (62.1 * float64(weight)) - 44.7
	} else {
		fmt.Println("Invalid gender.")
		return
	}

	fmt.Println("Ideal weight:", idealWeight)
}

func validInput(prompt string) (int, string, error) {
	var weight int
	var genero string

	// Lendo o peso
	fmt.Print(prompt)
	_, err := fmt.Scan(&weight)
	if err != nil {
		return 0, "", errors.New("digite um valor inteiro")
	}

	// Lendo o gênero
	fmt.Print("Enter your gender: ")
	_, err = fmt.Scan(&genero)
	if err != nil {
		return 0, "", errors.New("digite um texto")
	}

	return weight, genero, nil

}
