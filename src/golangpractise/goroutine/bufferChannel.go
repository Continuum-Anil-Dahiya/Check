package goroutine

import "fmt"

// Check3 Buffered Channel having size two
// Send will block if it is full and recieve will block if it is empty
func Check3() {

	a := make(chan int, 2)
	a <- 2
	a <- 3
	fmt.Println(<-a)
	fmt.Println(<-a)
}
