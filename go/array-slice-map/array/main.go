package main

import "fmt"

func main() {
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 0, 1, 2
	fmt.Println(notas)

	soma := 0.0

	for i := 0; i < len(notas); i++ {
		soma += notas[i]
	}

	media := soma / float64(len(notas))
	fmt.Println(media)
}
