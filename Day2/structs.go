package main

import "fmt"

type Sport struct {
	sportName string
	position  string
}

type Person struct {
	id     int
	name   string
	sports []Sport
}

func (p Person) getName() string {
	return p.name
}

func structs() {
	fmt.Println("Structs in go")
	// Empty version of a struct
	p1 := Person{name: "Bhushan", id: 1}
	name := p1.getName()
	fmt.Println("Name of the person is : ", name)
	fmt.Println(p1)
}
