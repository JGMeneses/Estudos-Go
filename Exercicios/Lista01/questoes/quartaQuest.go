package questoes

import "fmt"

func Four() {
	var numero, percentual float64

	fmt.Print("Digite um número inteiro N: ")
	fmt.Scan(&numero)
	fmt.Print("Digite um percentual: ")
	fmt.Scan(&percentual)

	result := (numero) * (percentual / 100.0)

	fmt.Println("Valor final:", result)
}
