package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 12332.55,
			"guga":     21221.3,
		},
		"J": {
			"jao": 212.1,
		},
	}

	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
