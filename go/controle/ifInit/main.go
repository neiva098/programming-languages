package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	a := rand.NewSource(time.Now().UnixNano())
	r := rand.New(a)

	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i < 5 {
		fmt.Println("ganhou")
	} else {
		fmt.Println("Perdeu")
	}
}
