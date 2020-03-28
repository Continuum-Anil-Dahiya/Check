package array

import (
	"fmt"
	"reflect"
	"time"
)

//Check ...
func Check() {

	var ch [5]int
	var t [3]time.Time
	var str [5]string
	//	var str1 [5]string = [5]string{"a", "", "b"}
	str2 := [5]string{
		"Hi I am Anil ",
		"I was using this",
		"Can You help me out ",
	}

	// for int array
	fmt.Println(reflect.TypeOf(ch))             // Type is [5]int
	fmt.Println("Default Data of int[] Ch", ch) // Default Data is [0 0 0 0 0]

	// for time array
	fmt.Println(reflect.TypeOf(t))                  // Type is [3]time.Time
	fmt.Println("Default Data of time.Time[] t", t) // Defalut value [0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC]
	//fmt.Println("value of t[3]", t[3])            // Array Index out of bound

	// for str string array
	fmt.Println("Default values for string", str)                  // Default values for string is blank
	fmt.Println("Type of single Element ", reflect.TypeOf(str[0])) // string

	// for str1 string
	//fmt.Println("str[0]" + str1[0])

	// for str2 string
	fmt.Println("str2[0]", str2[0], str2[1], str2[2])
	//fmt.Printf("%#v", str2)
	fmt.Printf("%T", str2) // To find Type

}
