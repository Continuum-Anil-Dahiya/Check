package sort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Methods16 calls from main2/ch.go
func Methods16() {

	reader := bufio.NewReader(os.Stdin)
	newLine, _ := reader.ReadString('\n')
	str := strings.Split(strings.TrimSpace(newLine), " ")

	for i := 0; i < len(str); i++ {

		a, _ := strconv.Atoi(str[i])
		min := a
		index := i
		for j := i + 1; j < len(str); j++ {
			b, _ := strconv.Atoi(str[j])
			if b < min {
				min = b
				index = j
			}
		}

		str[index] = strconv.Itoa(a)
		str[i] = strconv.Itoa(min)

	}

	fmt.Println(str)

}
