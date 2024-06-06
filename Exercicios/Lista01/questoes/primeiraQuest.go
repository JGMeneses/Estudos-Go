package questoes

import "fmt"

// One é uma função que calcula a idade em dias
func One() {
	var dias, meses, ano int

	fmt.Print("Quantos anos tem:")
	fmt.Scan(&ano)
	fmt.Print("Digite sua idade em meses:")
	fmt.Scan(&meses)
	fmt.Print("Digite sua idade em dias:")
	fmt.Scan(&dias)

	idadeFinal := ano*360 + (meses*30 + dias)

	fmt.Print(idadeFinal)
}
