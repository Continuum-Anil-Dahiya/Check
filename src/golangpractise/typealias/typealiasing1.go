package typealias

import (
	"fmt"
	"strings"
)

type mystr string

func (st mystr) upperCase() {

	fmt.Println(strings.ToUpper(string(st)))

}

func (st *mystr) lowerCase() {
	*st = "Check "

	//fmt.Println(strings.ToLower(string(st)))
}
func (st mystr) lowerCase1() {
	st = mystr(strings.ToLower(string(st)))
	fmt.Println(st)
	//fmt.Println(strings.ToLower(string(st)))
}

func (st mystr) trim() {

	fmt.Println(strings.Trim(string(st), " "))
}

//Method4 calls from main2/ch
func Method4() {
	str1 := mystr(" Hello ")

	fmt.Println(string(str1))

	str := &str1
	str.lowerCase()
	str.lowerCase1()

	//str.upperCase()
	//str.trim()
}
