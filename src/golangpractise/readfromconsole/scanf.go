package readfromconsole

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Check2 is method for taking input from console
func Check2() {

	var a, c string
	abc := 4
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &c)

	// Converting string to interger
	b, err := strconv.ParseInt(a, 10, 64)
	//first attribute : variable to convert in INT
	//Second attribute : consider input as type of it 8 represents input is octal and convert answer into decimal
	//Third attribute : Convert it into 64 or 32 bit int

	// Replace functionality
	replacer := strings.NewReplacer("*", "#")
	c = replacer.Replace(c)

	if err != nil {
		log.Fatal("Error")
	}
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T %d", abc, abc)

}
