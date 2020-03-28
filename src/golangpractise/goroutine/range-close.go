package goroutine

import "fmt"

//Check4 ... range and close
func Check4() {

	a := make(chan int, 10)

	go closerange(a, cap(a))
	for i := range a {
		fmt.Println(i)
	}

}

func closerange(a chan int, b int) {

	for i := 0; i < b; i++ {
		a <- i
	}

	close(a)
}
