package main

import (
	"fmt"
	"projeto001/Exercicios/Lista01/questoes"
)

func main() {
	// Chame as funções dos arquivos primeiraQuest.go e segundaQuest.go conforme necessário
	var escolha int

	fmt.Print("Qual questão quer acessar? \n 1- Primeira questão \n 2- Segunda questão \n 3- Terceira questão")
	fmt.Scan(&escolha) // Ler a escolha do usuário

	switch escolha {
	case 1:
		questoes.One()
	case 2:
		questoes.Two()
	case 3:
		questoes.Three()
	default:
		fmt.Println("Opção inválida!")
	}
}
