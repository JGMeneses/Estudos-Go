package main

import (
	"fmt"
	"time"
)

func minhaGoroutine() {
	// Implemente sua lógica de goroutine aqui
	for i := 0; i < 5; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(time.Second)
	}
}

func main() {
	// Inicia uma nova goroutine
	go minhaGoroutine()

	// Lógica principal do programa
	for i := 0; i < 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(time.Second)
	}
}
