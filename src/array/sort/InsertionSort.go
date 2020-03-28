package sort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Methods15 Calls from main2/ch.go
func Methods15() {

	reader := bufio.NewReader(os.Stdin)
	newLine, _ := reader.ReadString('\n')
	str := strings.Split(strings.TrimSpace(newLine), " ")

	fmt.Println(str)
	for i := 1; i < len(str); i++ {

		a, _ := strconv.Atoi(str[i])

		j := i - 1
		min := -1
		for ; j >= 0; j-- {
			b, _ := strconv.Atoi(str[j])
			if a >= b {
				min = j
				break
			}
		}

		if j == -1 {
			min = 0
		}

		for k := i; k > min; k-- {
			str[k] = str[k-1]
		}

		str[j+1] = strconv.Itoa(a)

	}

	fmt.Println(str)

}
