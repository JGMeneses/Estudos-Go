package questions02

import (
	"fmt"
)

//  Uma empresa concederá um aumento de salário aos seus funcionários, variável de acordo com o cargo,
// conforme a tabela abaixo. Faça um algoritmo que leia o salário e o código do cargo de um funcionário e
// calcule o novo salário. Se o cargo do funcionário não estiver na tabela, ele deverá receber 40% de
// aumento. Mostre o salário antigo, o novo salário e a diferença.

func Five() {
	salary, cod, err := valid("Enter your salary: ", "Enter the cod your job:")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := applyingIncrease(salary, cod)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(value)
}

func applyingIncrease(salary, cod int) (float64, error) {
	var value float64

	if cod == 101 {
		value = float64(salary) * 0.1
	} else if cod == 102 {
		value = float64(salary) * 0.2
	} else if cod == 103 {
		value = float64(salary) * 0.3
	} else {
		value = float64(salary) * 0.4
	}

	value = float64(salary) + value
	return value, nil
}
