package array

import "fmt"

//Minimum number of swaps required to sort an array of first N number

//MinimumSwaptoSortArray ...
func MinimumSwaptoSortArray() {

	var t int
	fmt.Scan(&t)

	for t > 0 {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		minimumSwaptoSortArray(arr, n)

		t--
	}

}

func minimumSwaptoSortArray(arr []int, n int) {

	count := 0
	for i := 0; i < n; i++ {

		if arr[i] != i+1 {
			for arr[i] != i+1 {
				temp := arr[i]
				arr[i] = arr[temp-1]
				arr[temp-1] = temp
				count++
			}
		}
	}

	fmt.Println(count)

}
