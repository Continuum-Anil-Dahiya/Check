package typealias

import "fmt"

type mystring string

func (s mystring) greeting3() {

	fmt.Println(s)
}

//Method3 calls from main2/ch
func Method3() {

	Mystr := mystring("Hello")
	Mystr.greeting3()
}
