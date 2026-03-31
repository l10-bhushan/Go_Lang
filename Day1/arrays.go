package main

import "fmt"

func arrays() {
	var arr [2]int
	fmt.Println(arr)
	// Array literal
	arr2 := [2]int{1, 2}
	fmt.Println(arr2)
	// Multi-dimensional array
	multi := [2][2]int{{1, 2}, {3, 4}}
	// Updating array
	multi[0] = [2]int{10, 20}
	fmt.Println(multi)
	fmt.Println(len(multi))

	for _, nested := range multi {
		for _, value := range nested {
			fmt.Println("Value is : ", value)
		}
	}
}
