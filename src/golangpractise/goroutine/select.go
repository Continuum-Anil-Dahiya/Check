package goroutine

import (
	"fmt"
	"time"
)

//Check5 shows select magic
func Check5() {

	a := make(chan int)
	quit := make(chan int)
	go selectImpl(a, quit)

	for i := 0; i < 10; i++ {
		fmt.Println(<-a)
	}

	tick := time.Tick(100 * time.Millisecond)
	fmt.Println(tick)
	quit <- 0

}

func selectImpl(a chan int, quit chan int) {

	x := 0
	for {
		select {
		case a <- x:
			fmt.Println("a")

		case <-quit:
			fmt.Println("quiting")
			return

		}
	}

}
