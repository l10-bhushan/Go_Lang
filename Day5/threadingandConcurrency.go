package main

import (
	"fmt"
	"time"
)

// func run() {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("run")
// }

// func run2() {
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("run2")
// }

// func run3() {
// 	time.Sleep(6 * time.Second)
// 	fmt.Println("run3")
// }

// We use channels when we need to wait for result of a go routine
func Add(x int, y int, ch chan int, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println(x + y)
	ch <- x + y
}

func ThreadAndConcurrency() {
	// go run()
	// go run2()
	// go run3()
	// time.Sleep(7 * time.Second)
	// Multiple channels
	// The order of execution cannot be determined
	ch := make(chan int)
	ch2 := make(chan int)
	go Add(100, 200, ch, 4)
	go Add(23, 43, ch2, 2)

	select {
	case x := <-ch:
		fmt.Println(x)
	case y := <-ch2:
		fmt.Println(y)
	}
	// In the above statement whatever we get first is handled and we move on with our code
	// x := <-ch
	// y := <-ch2

	// go Add(5, 10, ch)
	// go Add(10, 10, ch)
	// go Add(89, 90, ch)
	// go Add(45, 3, ch)
	// // A operation where we are waiting for result is a blocking operation
	// // So if you add any other go routine after that it won't run parallely
	// x := <-ch
	// x = <-ch
	// x = <-ch
	// x = <-ch
	// fmt.Println(x, y)
}

func r(ch chan bool) {
	<-ch
}

func BufferedChannel() {
	ch := make(chan bool)
	go r(ch)
	ch <- true
	fmt.Println("Done")

	// Buffered channel
	// This no denotes no of values that can be added to the channel before we start having a blocking channel

	// ch := make(chan bool, 2)
}
