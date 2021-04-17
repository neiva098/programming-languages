package main

import (
	"fmt"
)

func main() {
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "Pedro"
	aprovados[3] = "Jao"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println(nome, cpf)
	}

	fmt.Println(aprovados[2])

	delete(aprovados, 2)

	fmt.Println(aprovados)
}
