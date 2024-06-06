package questoes

import "fmt"

// Leia a idade de uma pessoa expressa em anos, meses e dias e mostre-a expressa apenas em dias.
// Obs: Considere cada mês com 30 dias.
func One() {
	var dias, meses, ano int

	fmt.Print("Quantos anos tem:")
	fmt.Scan(&ano)
	fmt.Print("meses:")
	fmt.Scan(&meses)
	fmt.Print("dias:")
	fmt.Scan(&dias)

	idadeFinal := ano*360 + (meses*30 + dias)

	fmt.Print(idadeFinal)
}
