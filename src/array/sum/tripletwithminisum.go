package sum

import (
	"fmt"
	"math"
	"sort"
)

//Find triplet with minimum sum not need to be contigous
//https://www.geeksforgeeks.org/find-triplet-with-minimum-sum/

//MinSumTriplet ...
func MinSumTriplet() {

	var t int
	fmt.Scan(&t)

	for t > 0 {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		minSumTripletUtil2(arr, n)

		t--
	}

}

//O(nlogn)
func minSumTripletUtil(arr []int, n int) {
	sort.Ints(arr)

	fmt.Println(arr[0] + arr[1] + arr[2])
}

//O(n)
func minSumTripletUtil2(arr []int, n int) {

	min1 := math.MaxInt64
	min2 := math.MaxInt64
	min3 := math.MaxInt64

	for i := 0; i < n; i++ {

		if min1 > arr[i] {
			min3 = min2
			min2 = min1
			min1 = arr[i]
		} else if min2 > arr[i] {
			min3 = min2
			min2 = arr[i]
		} else if min3 > arr[i] {
			min3 = arr[i]

		}

	}

	fmt.Println(min1 + min2 + min3)
}
