package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else if nota > 5 {
		fmt.Println("Reprovado")
	} else {
		fmt.Println("Bomba")
	}
}

func main() {
	imprimirResultado(7)
	imprimirResultado(6)
}
