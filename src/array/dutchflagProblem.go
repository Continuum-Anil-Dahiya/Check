package array

import "fmt"

//DutchFlagProblem ... for sorting of 0's,1's,2's
func DutchFlagProblem() {

	var t int
	fmt.Scan(&t)

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sortArray(arr)
	fmt.Println(arr)

}

func sortArray(arr []int) {

	low := 0
	mid := 0
	high := len(arr) - 1

	for mid < high {
		switch arr[mid] {
		case 0:
			arr[mid], arr[low] = swapInteger(arr[mid], arr[low])
			mid++
			low++
		case 1:
			mid++
		case 2:
			arr[mid], arr[high] = swapInteger(arr[mid], arr[high])
			high--
		}
	}
}

func swapInteger(a int, b int) (int, int) {
	return b, a
}
