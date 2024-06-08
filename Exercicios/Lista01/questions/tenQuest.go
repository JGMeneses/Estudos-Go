package questions

import "fmt"

// Leia dois valores inteiros para as variáveis A e B e efetue as operações de adição, subtração,
// multiplicação e divisão de A por B, apresentando ao final os resultados obtidos. Em seguida leia
// dois valores lógicos C e D e efetue as operações de negação (de cada um dos valores), conjunção
// (E) e disjunção (OU), apresentando ao final os resultados obtidos.

// Read two integer values into variables A and B and perform addition, subtraction,
// multiplication, and division of A by B, presenting the results obtained at the end. Then read
// two logical values C and D and perform the operations of negation (of each value), conjunction
// (AND), and disjunction (OR), presenting the results obtained at the end.
func Ten() {
	var A, B int
	var C, D bool

	fmt.Println("Enter two integer values for A and B:")
	fmt.Scan(&A, &B)

	fmt.Println("Addition:", A+B)
	fmt.Println("Subtraction:", A-B)
	fmt.Println("Multiplication:", A*B)
	if B != 0 {
		fmt.Println("Division:", A/B)
	} else {
		fmt.Println("Cannot divide by zero.")
	}

	fmt.Println("Enter two boolean values for C and D:")
	fmt.Scan(&C, &D)

	fmt.Println("Negation of C:", !C)
	fmt.Println("Negation of D:", !D)
	fmt.Println("Conjunction (C && D):", C && D)
	fmt.Println("Disjunction (C || D):", C || D)
}
