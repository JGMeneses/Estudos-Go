package questoes

import "fmt"

// Calcule e escreva a quantidade de dinheiro gasta por um fumante a partir da leitura das
// informações número de anos que ele fuma, quantidade de cigarros fumados por dia e preço de
// uma carteira.
func Two() {
	var anosFumando, qtdDiaria, preco float64

	fmt.Println("Quantos anos fuma:")
	fmt.Scan(&anosFumando)
	fmt.Println("quantidade de cigarros fumados por dia :")
	fmt.Scan(&qtdDiaria)
	fmt.Println(" preço da carteira:")
	fmt.Scan(&preco)
	gasto := (qtdDiaria / 12 * (anosFumando * 365))
	fmt.Print("Seu gasto com cigarros até hoje: ", gasto)
}
