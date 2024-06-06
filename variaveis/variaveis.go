package main

import "fmt"

func main() {
	// Tipos Básicos
	var inteiro int = 10
	var flutuante float64 = 3.14
	var booleano bool = true
	var texto string = "Olá, Mundo!"

	// Tipos Compostos
	var array [3]int = [3]int{1, 2, 3}
	var slice []int = []int{1, 2, 3, 4, 5}

	// Inicializando um map
	mapa := make(map[string]int)
	mapa["chave"] = 1

	// Struct
	type Pessoa struct {
		Nome  string
		Idade int
	}
	var pessoa Pessoa = Pessoa{"Alice", 30}

	// Ponteiro
	var ponteiro *int = &inteiro

	// Channel
	var ch chan int = make(chan int)

	fmt.Println(inteiro, flutuante, booleano, texto)
	fmt.Println(array, slice)
	fmt.Println(mapa)
	fmt.Println(pessoa)
	fmt.Println(ponteiro)
	fmt.Println(ch)
}
