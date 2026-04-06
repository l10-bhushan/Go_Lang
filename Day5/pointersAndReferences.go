package main

import "fmt"

type Book struct {
	id    int
	title string
}

func (b *Book) setTitle(newTitle string) {
	b.title = newTitle
}

func change(a *int) {
	*a = 20
}

func pointersAndReferences() {
	fmt.Println("Pointer and References in go")
	x := 10
	y := &x // & denotes the memory address of x
	// * is the dereferencing operator returns value stored at y
	fmt.Println(x, *y)

	// To change the value of x we can change it through y
	*y = 100
	fmt.Println(x, *y)
	// Now we will get 100, 100
	z := 10
	change(&z)
	fmt.Println(z)

	b := Book{
		10,
		"Harry potter",
	}

	b.setTitle("Naruto")
	fmt.Println(b)

	i := 10
	j := &i
	k := &j

	fmt.Println(i, *j, **k)

}
