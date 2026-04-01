package main

import "fmt"

func maps() {
	fmt.Println("Maps")

	mp := map[string]int{"a": 2}
	fmt.Println(mp)
	//mp1 := make(map[string]int)
	mp2 := make(map[int][]string)

	// Adding key value pair
	mp2[1] = []string{"Hello"}
	mp2[2] = []string{"World,", "I'm Leo"}
	fmt.Println(mp2)

	// deleting key value pair
	delete(mp2, 1)

	// check if key, value is present in map

	value, ok := mp2[2]
	fmt.Println(value, ok)

	divisible := map[int]int{}
	numbers := 100

	for i := 1; i <= numbers; i++ {
		for j := 1; j <= 5; j++ {
			if i%j == 0 {
				divisible[j]++
			}
		}
	}

	fmt.Println(divisible)

}
