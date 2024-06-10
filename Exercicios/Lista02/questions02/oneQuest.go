package questions02

import (
	"errors"
	"fmt"
)

// Given three values X, Y, and Z, check and write if they can be the lengths of the sides
// of a triangle. If they can, determine and write if the triangle is equilateral,
// isosceles, or scalene. If they do not form a triangle, write this message.
// Consider the following properties:
// The length of each side of a triangle is less than the sum of the other two sides.
// • Equilateral: all three sides have equal lengths.
// • Isosceles: two sides have equal lengths.
// • Scalene: all three sides have different lengths.

// Dados três valores X, Y e Z, verifique e escreva se eles
// podem ser os comprimentos dos lados de um
// triângulo e, se forem, verificar e escrever se é um triângulo equilátero,
// isósceles ou escalenos. Se eles
// não formarem um triângulo, escrever esta mensagem. Considere as seguintes
// propriedades:
// • O comprimento de cada lado em um triângulo é menor que a soma dos outros
// dois lados;
// • Equiláteros: tem os comprimentos dos três lados iguais;
// • Isósceles: tem os comprimentos de dois lados iguais;
// • Escaleno: tem os comprimentos dos três lados diferentes.

func One() {
	x, err := getInput("Enter X: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	y, err := getInput("Enter Y: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	z, err := getInput("Enter Z:")
	if err != nil {
		fmt.Println(err)
		return
	}

	if isValidTriangle(x, y, z) {
		triangleType := triangleType(x, y, z)
		fmt.Println("The triangle is", triangleType)
	} else {
		fmt.Println("The values do not form a triangle")

	}
}

func isValidTriangle(x, y, z int) bool {
	return x < y+z && y < z+x && z < x+y
}

func getInput(prompt string) (int, error) {
	var value int
	fmt.Print(prompt)
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, errors.New("invalid input, please enter an integer value")
	}
	if value <= 0 {
		return 0, errors.New("invalid input, please enter a positive integer value")
	}
	return value, nil
}

func triangleType(x, y, z int) string {
	if x == y && y == z {
		return "equilateral"
	}
	if x == y || y == z || x == z {
		return "isosceles"
	}
	return "scalene"
}
