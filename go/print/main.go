package main

import "fmt"

func main() {
	fmt.Print("ola")
	fmt.Print(" linha")

	fmt.Println("cria linha")
	fmt.Println("cria linha de novo")

	x := 3.141516

	fmt.Println("x = ", x)

	xs := fmt.Sprint(x)

	fmt.Println("x = " + xs)

	fmt.Printf("O valor de x é %.2f", x)

	a, b, c, d := 1, 1.9999, false, "opa"

	fmt.Printf("\n%d %.1f %t %s", a, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
