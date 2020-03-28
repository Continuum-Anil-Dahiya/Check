package goroutine

import "fmt"

//Check2 for Checking Channel used for inter communication between two thread
func Check2() {

	a := make(chan int)

	go comm(a)
	go comm(a)

	x, y := <-a, <-a

	fmt.Println(x, y)
}

func comm(a chan int) {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}

	a <- sum
}
