package dp

import (
	"fmt"
	"math"
)

//This contains kadane solution for Maxcontiguous subarray and Mincontiguous subarray

//Main3 ...
func Main3() {

	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		fmt.Println(findMinSubarray(arr, n))
		t--
	}

}

func findMaxSubarray(arr []int, n int) int {

	maxSoFar := math.MinInt64
	maxEndHere := 0

	for i := 0; i < n; i++ {

		maxEndHere = maxEndHere + arr[i]

		if maxEndHere > maxSoFar {
			maxSoFar = maxEndHere
		}

		if maxEndHere < 0 {
			maxEndHere = 0
		}

	}

	return maxSoFar

}

func findMinSubarray(arr []int, n int) int {

	minSoFar := math.MaxInt64
	minEndHere := 0

	for i := 0; i < n; i++ {

		minEndHere = minEndHere + arr[i]

		if minEndHere < minSoFar {
			minSoFar = minEndHere
		}

		if minEndHere > 0 {
			minEndHere = 0
		}

	}

	return minSoFar

}
