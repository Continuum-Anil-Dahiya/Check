package array

import (
	"fmt"
	"math"
)

//Element with left side smaller and right side greater
//Start from left and side and make an array with small element till.

//Main2 ...
func Main2() {

	var t int
	fmt.Scan(&t)

	for t > 0 {
		var n int
		fmt.Scan(&n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		arrLeftSmall := make([]int, n)

		min := math.MaxInt64
		for i := n - 1; i >= 0; i-- {

			if arr[i] < min {
				min = arr[i]
			}
			arrLeftSmall[i] = min
		}

		arrLeftSmall[n-1] = math.MinInt64
		max := math.MinInt64
		check := false
		for i := 0; i < n; i++ {
			if max < arr[i] {
				max = arr[i]
			}

			if arr[i] == max && arr[i] == arrLeftSmall[i] {
				fmt.Println(arr[i])
				check = true
				break
			}
		}

		if check == false {
			fmt.Println("-1")
		}
		t--
	}
}
