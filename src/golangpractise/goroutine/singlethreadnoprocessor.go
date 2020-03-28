package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

// func init() {

// }

//Start1 ...
func Start1() {

	runtime.GOMAXPROCS(2)
	fmt.Println("runtime")
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		check()
		wg.Done()
	}()

	go func() {
		check1()
		wg.Done()
	}()

	wg.Wait()
}

func check() {

	for i := 0; i < 600; i++ {
		fmt.Println("In Thread 1")
		//fmt.Println(i)
	}

}

func check1() {

	for i := 0; i < 600; i++ {
		fmt.Println("In Thread 2")
		//fmt.Println(i)
	}

}
