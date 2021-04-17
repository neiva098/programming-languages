package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("x" == "x")
	fmt.Println("x" != "x")
	fmt.Println(3 < 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"joao"}
	p2 := Pessoa{"pedro"}

	fmt.Println(p1 == p2)
}
