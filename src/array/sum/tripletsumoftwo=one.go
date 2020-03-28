package sum

import (
	"fmt"
	"sort"
)

//Find a triplet such that sum of two equals to third element if no such element found print -1
//https://www.geeksforgeeks.org/find-triplet-sum-two-equals-third-element/
// It will accept negaive number as well like -3 1 2 3 5 7 10  -> 7,10,-3

//FindTriplets ...
func FindTriplets() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		findTripletUtil(arr, n)

		t--
	}

}

func findTripletUtil(arr []int, n int) {

	sort.Ints(arr)

	flag := false
	for i := n - 1; i >= 0; i-- {
		j := 0
		k := n - 1
		for j < k {

			if arr[i] == arr[j]+arr[k] && i != j && i != k {
				fmt.Println("Triplet found")
				fmt.Println(arr[i], arr[j], arr[k])
				j++
				flag = true
			} else if arr[i] > arr[j]+arr[k] {
				j++
			} else {
				k--
			}

		}
	}

	if !flag {
		fmt.Println("-1")
	}

}
