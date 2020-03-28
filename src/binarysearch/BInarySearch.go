package binarysearch

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Methods10 BinarySearch
func Methods10() {

	reader := bufio.NewReader(os.Stdin)
	array, _ := reader.ReadString('\n')
	array = strings.TrimSpace(array)
	splitedArray := strings.Split(array, " ")
	element, _ := reader.ReadString('\n')

	index := binarySearch(splitedArray, strings.TrimSpace(element), 0, len(splitedArray)-1)

	if index == -1 {
		fmt.Println("Element Not Found")
		return
	}
	fmt.Printf("Element Found at %d", index)

}

func binarySearch(array []string, element string, low int, high int) int {

	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if array[mid] == element {
		return mid
	} else if array[mid] < element {
		low = mid + 1
	} else {
		high = mid - 1
	}

	return binarySearch(array, element, low, high)
}
