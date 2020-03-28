package readfromconsole

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var abc = "abc"

//Check1 is used for Read Io from console
func Check1() {
	reader := bufio.NewReader(os.Stdin)
	next, err := reader.ReadString('\n')

	//Convert to Integer
	var a int
	a, err = strconv.Atoi(strings.TrimSpace(next))

	//Convert to Float
	var b float64
	b, err = strconv.ParseFloat(strings.TrimSpace(next), 64)

	if err != nil {
		log.Fatal("Error")
	}

	fmt.Println(next)
	fmt.Println("Convert from String to Int ", a)
	fmt.Println("Convert from String to Float ", b)
}
