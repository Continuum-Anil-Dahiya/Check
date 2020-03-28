package pointer

import "fmt"

import 

// PointerReturn demos the pass by reference
func PointerReturn() {
	var a = 12
	var b = 14
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	example(&a, &b)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
}

func example(a *int, b *int) {
	c := *b
	*b = *a
	*a = c
}
