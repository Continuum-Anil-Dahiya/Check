package readfromconsole

import "fmt"

//Check4  Use for behaviour of printf
func Check4() {

	var a = 4
	var b = 44.1
	var c = false
	var d = "Read"

	fmt.Printf("%d\n", a)   // Integer
	fmt.Printf("%f\n", b)   // float
	fmt.Printf("%t\n", c)   // bool
	fmt.Printf("%s\n", d)   // String
	fmt.Printf("%T\n", c)   // Print the Type of format
	fmt.Printf("%v\n", a)   // Used for all type
	fmt.Printf("%6d\n", a)  // Padding Space
	fmt.Printf("%10.3f", b) // Need total 10 chars if not add padding and 3 after decimal
}
