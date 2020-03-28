package sum

import (
	"fmt"
	"strings"
)

// https://www.geeksforgeeks.org/given-an-array-of-numbers-arrange-the-numbers-to-form-the-biggest-number/

// Main1 ...
func Main1() {

	//t is testcases
	var t int
	fmt.Scan(&t)

	//size of array
	var n int
	fmt.Scan(&n)

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	mergeSort(arr, 0, n-1)

	for _, v := range arr {
		fmt.Print(v)
	}

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

		if strings.Compare(left[i]+right[j], right[j]+left[i]) == 1 {
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
