package dp

import (
	"fmt"
	"math"
)

//Codechef problem https://www.codechef.com/problems/DELISH

//Main2 ...
func Main2() {

	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}

		minEndingHere := math.MaxInt64
		minSoFar := math.MaxInt64
		maxEndingHere := math.MinInt64
		maxSoFar := math.MinInt64
		var i, left, right int
		var l, leftmax, rightmax int
		for k := 0; k < n; k++ {
			if minEndingHere > 0 {
				minEndingHere = arr[k]
				i = k
			} else {
				minEndingHere = arr[k] + minEndingHere
			}
			min := math.Min(float64(minSoFar), float64(minEndingHere))
			if min == float64(minEndingHere) {
				right = k
				left = i
			}
			minSoFar = int(min)

			if maxEndingHere < 0 {
				maxEndingHere = arr[k]
				l = k
			} else {
				maxEndingHere = arr[k] + maxEndingHere
			}
			max := math.Max(float64(maxSoFar), float64(maxEndingHere))
			if max == float64(maxEndingHere) {
				rightmax = k
				leftmax = l
			}
			maxSoFar = int(max)

		}

		// fmt.Println(left)
		// fmt.Println(right)
		// fmt.Println(minSoFar)

		// fmt.Println(leftmax)
		// fmt.Println(rightmax)
		// fmt.Println(maxSoFar)

		var value int
		if leftmax == right {

			value = maxSoFar - (arr[leftmax] + minSoFar)
		} else if left == rightmax {

			value = maxSoFar - (arr[left] + minSoFar)
		} else {
			value = maxSoFar - minSoFar
		}

		fmt.Println(value)

	}

}
