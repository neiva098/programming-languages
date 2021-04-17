package main

import (
	"fmt"
)

func main() {
	funcsESalarios := map[string]float64{
		"jao": 1212.4,
		"ana": 1212.99,
	}

	fmt.Println(funcsESalarios)

	funcsESalarios["rafael"] = 812981.0

	delete(funcsESalarios, "inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
