package sort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Methods14  calls from ch.go
func Methods14() {

	reader := bufio.NewReader(os.Stdin)
	newLine, _ := reader.ReadString('\n')
	str := strings.Split(strings.TrimSpace(newLine), " ")

	fmt.Println(str)
	for i := 0; i < len(str)-1; i++ {

		for j := 0; j < len(str)-i-1; j++ {
			a, _ := strconv.Atoi(str[j])
			b, _ := strconv.Atoi(str[j+1])
			if a > b {
				a, b = swap(a, b)

				str[j] = strconv.Itoa(a)
				str[j+1] = strconv.Itoa(b)
			}
		}

	}

	fmt.Println(str)

}

func swap(a int, b int) (int, int) {

	return b, a

}
