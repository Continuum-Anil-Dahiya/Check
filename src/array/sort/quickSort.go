package sort

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Methods18 calls from main2/ch.go
func Methods18() {

	reader := bufio.NewReader(os.Stdin)
	newLine, _ := reader.ReadString('\n')
	str := strings.Split(strings.TrimSpace(newLine), " ")
	fmt.Println(str)
	quickSort(str, 0, len(str)-1)
	fmt.Println(str)

}

func quickSort(str []string, low int, high int) {

	if low < high {
		index := partition(str, low, high)
		quickSort(str, low, index-1)
		quickSort(str, index+1, high)

	}
}

func partition(str []string, low int, high int) int {

	pivoit := str[high]
	pivoitIndex := low

	// One less than last index as we set last as a pivoit
	for i := low; i < high; i++ {

		if str[i] <= pivoit {
			str[i], str[pivoitIndex] = swapElement(str[i], str[pivoitIndex])
			pivoitIndex++
		}
	}

	str[high], str[pivoitIndex] = swapElement(str[high], str[pivoitIndex])

	return pivoitIndex

}

func swapElement(str1 string, str2 string) (string, string) {

	return str2, str1
}
