package questions02

import (
	"errors"
	"fmt"
)

// A escola “APRENDER” faz o pagamento de seus professores por hora/aula. Faça um algoritmo que calcule
// e escreva o salário de um professor, sabendo que o valor da hora/aula segue a tabela abaixo:

// the “APRENDER” school pays its teachers by the hour. make an algorithm that calculates and writes a teacher's salary.

func Six() {
	hours, lvl, err := valid("Enter work hours: ", "Enter teacher level:")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := salary(hours, lvl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The teacher's salary is:", value)
}

func salary(hours, lvl int) (int, error) {
	var value int

	switch lvl {
	case 1:
		value = hours * 12
	case 2:
		value = hours * 17
	case 3:
		value = hours * 25
	default:
		return 0, errors.New("there are only 3 levels of teachers")
	}

	return value, nil
}
