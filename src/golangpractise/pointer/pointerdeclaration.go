package pointer

import "fmt"

//Check1 demos the Pointer declartion and defination
func Check1() {

	var floatpointer *float64
	var intpointer *int
	var a = 32
	var b = 45.9
	intpointer = &a
	floatpointer = &b
	fmt.Println(intpointer)
	fmt.Println(*intpointer)
	fmt.Println(floatpointer)
	fmt.Println(*floatpointer)
}
