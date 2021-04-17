package main

import "fmt"

func main() {
	i := 0

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Laco infinito")
	}
}
