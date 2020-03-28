package array

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Check3 is made for generating random number
func Check3() {

	var i int

	rand.Seed(time.Now().UnixNano())
	randonNumber := rand.Intn(100)

	flag := false
	for i = 0; i < 10; i++ {

		fmt.Println("Your ", 10-i, " chance is left  .. please ")
		fmt.Println("Guess a Number")

		reader := bufio.NewReader(os.Stdin)
		next, _ := reader.ReadString('\n')
		val, _ := strconv.Atoi(strings.TrimSpace(next))

		if val == randonNumber {
			fmt.Println("Good Job! You Guessed it!")
			flag = true
			break
		} else if val > randonNumber {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Oops. Your guess was LOW.")
		}

	}

	if flag == false {
		fmt.Println("Sorry. You Didn'tg guess my number. It was:[", randonNumber, "]")
	}

}
