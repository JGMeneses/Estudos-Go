package main

import "fmt"

func main() {
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

/*

Opção padrão
func main() {
    var nome string
    var idade int

    fmt.Print("Digite seu nome: ")
    fmt.Scanln(&nome)

    fmt.Print("Digite sua idade: ")
    fmt.Scanln(&idade)

    fmt.Printf("Olá, %s! Você tem %d anos.\n", nome, idade)
}
*/
