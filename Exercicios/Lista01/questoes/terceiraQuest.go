package questoes

// Leia o valor de uma compra nas Lojas Sua Melhor Compra e mostre o valor das prestações de
// acordo com a quantidade de parcelas desejada pelo usuário. A loja está vendendo seus produtos
// a vista ou parcelado em até 10 vezes sem juros.
import "fmt"

func Three() {
	var valor int
	parcelas := 11
	fmt.Print("Digite o valor da sua compra:")
	fmt.Scan(&valor)

	for parcelas > 10 {
		fmt.Print("Digite quantidade de parcelas desejadas em até 10x:")
		fmt.Scan(&parcelas)
	}
	if parcelas <= 0 {
		fmt.Print("Então sua compra será avista e o valor total para pagar é R$ ", valor)
	} else {
		fmt.Print("Sua compra vai ser parcelada em:", parcelas, "o valor final para pagar será de R$ ", valor/parcelas)
	}
}
