package questions

import "fmt"

// Leia dois valores inteiros nas variáveis Val1 e Val2, troque os seus conteúdos e escreva o
// resultado. Exemplo: Se o valor lido foi 10 para a variável Val1 e 11 para a variável Val2, depois
// da troca o algoritmo terá que escreva Val1 = 11 e Val2 = 10.

// Read two integer values into the variables Val1 and Val2, swap their contents, and write the result.
// Example: If the value read was 10 for variable Val1 and 11 for variable Val2, after the swap
// the algorithm should write Val1 = 11 and Val2 = 10.
func Nine() {
	var val1, val2, val3 int

	fmt.Print("Enter the value for val1: ")
	fmt.Scan(&val1)
	fmt.Print("Enter the value for val1: ")
	fmt.Scan(&val2)

	val3 = val1

	val1 = val2

	val2 = val3

	fmt.Print("Val1: ", val1, " Val2: ", val2)

}
