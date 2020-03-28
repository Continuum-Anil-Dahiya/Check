package sort

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Methods17 calls from main2/ch.go
func Methods17() {

	reader := bufio.NewReader(os.Stdin)
	newLine, _ := reader.ReadString('\n')
	str := strings.Split(strings.TrimSpace(newLine), " ")
	fmt.Println(str)
	mergeSort(str, 0, len(str)-1)
	fmt.Println(str)

}

func mergeSort(str []string, low int, high int) {

	if low < high {
		mid := (high + low) / 2
		mergeSort(str, low, mid)
		mergeSort(str, mid+1, high)
		merge(str, low, mid, high)

	}

}

func merge(final []string, low int, mid int, high int) {

	i, j, l := 0, 0, low

	var left []string
	var right []string

	for k := low; k <= mid; k++ {
		left = append(left, final[k])
	}

	for k := mid + 1; k <= high; k++ {
		right = append(right, final[k])
	}

	for i < len(left) && j < len(right) {

		if left[i] < right[j] {
			final[l] = left[i]
			i++
			l++
		} else {
			final[l] = right[j]
			j++
			l++
		}
	}

	for i < len(left) {
		final[l] = left[i]
		l++
		i++
	}

	for j < len(right) {
		final[l] = right[j]
		j++
		l++
	}

}
