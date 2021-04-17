package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Cristiano", 20})

	fmt.Println(person{name: "Vania", age: 30})

	fmt.Println(person{name: "isabela"})

	fmt.Println(&person{name: "Margarete", age: 40})

	fmt.Println(newPerson("Joao"))

	s := person{name: "Filipe", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
