package array

import (
	"fmt"
)

//DuplicateElements in O(n) time complexity and O(1) space complexity ...
//https://www.youtube.com/watch?v=GeHOlt_QYz8
func DuplicateElements() {

	var t int
	fmt.Scan(&t)

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	dupicateElementUtil(arr, n)

}

func dupicateElementUtil(arr []int, n int) {

	for i := 0; i < n; i++ {

		element := arr[abs(arr[i])-1]
		if element > 0 {
			arr[abs(arr[i])-1] = -element
		} else {
			fmt.Println("Duplicate Found", abs(arr[i]))
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
