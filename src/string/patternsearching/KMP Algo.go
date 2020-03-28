package parternserching

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Methods9 will call from ch.go
func Methods9() {

	reader := bufio.NewReader(os.Stdin)
	findValue, _ := reader.ReadString('\n')
	findValue = strings.TrimSpace(findValue)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	run := []rune(findValue)
	arr := preProcessing(run)

	fmt.Println(arr)
	runString := []rune(str)
	var arr1 []int
	for i, j := 0, 0; i < len(runString); {

		fmt.Println(string(runString[i]))
		fmt.Println(i, j)
		if runString[i] == run[j] {
			j++
			i++
		} else if j != 0 {
			j = arr[j-1]
		} else {
			i++
		}

		if j == (len(run)) {
			arr1 = append(arr1, i-len(run))
			j = arr[j-1]
		}
	}

	fmt.Println("String found at location", arr1)

}

func preProcessing(run []rune) (arr []int) {

	var i, j int = 1, 0
	arr = []int{0}

	for i < len(run) {

		val := string(run[i])
		if string(run[j]) == val {
			j++
			arr = append(arr, j)
			i++
		} else if j == 0 {
			arr = append(arr, 0)
			i++
		} else {
			j = arr[j-1]
		}

	}

	return

}
