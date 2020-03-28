package goroutine

import "fmt"

//Check1 for Checking GORoutine calls from main2/ch.go
func Check1() {

	go say("Hello")
	say("World")
}

func say(a string) {

	for i := 0; i < 100; i++ {
		fmt.Println(a)
	}
}
