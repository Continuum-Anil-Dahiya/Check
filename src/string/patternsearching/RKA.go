package parternserching

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Methods13 calls from  main2/ch.go
func Methods13() {

	reader := bufio.NewReader(os.Stdin)

	//String given
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	run := []rune(str)

	//Pattern need to find
	pattern, _ := reader.ReadString('\n')
	pattern = strings.TrimSpace(pattern)

	hashPattern := findPatternHash(pattern)
	hash1Pattern := findPatternHash(string(run[0:len(pattern)]))

	var arr []int
	if hash1Pattern == hashPattern && string(run[0:len(pattern)]) == pattern {
		arr = append(arr, 0)
	}

	for i := len(pattern); i < len(run); i++ {
		hash1Pattern = (hash1Pattern-int(run[i-len(pattern)]))/26 + lastValue(int(run[i]), len(pattern))
		if hash1Pattern == hashPattern && pattern == string(run[(i-len(pattern)+1):i+1]) {
			arr = append(arr, (i - len(pattern) + 1))
		}
	}

	fmt.Println(arr)

}

func lastValue(value int, paternSize int) int {

	num := 1

	for i := 1; i < paternSize; i++ {
		num = num * 26
	}
	return num * value
}

func findPatternHash(pattern string) (sum int) {

	run := []rune(pattern)
	num := 1
	sum = 0
	for i := 0; i < len(run); i++ {

		k := int(run[i])
		sum += k * num
		num = num * 26

	}

	return

}
